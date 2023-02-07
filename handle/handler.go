package handle

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

// Handler 处理所有的HTTP请求
func Handler(db *sql.DB, md interface{}) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var op *Operate

		defer func() {
			if msg := recover(); msg != nil {
				debug.PrintStack()
				handlerError(op, w, fmt.Errorf("PANIC :: %s", msg))
			}
		}()

		context := NewContext(db, r, w)
		method, path, params, values := context.GetMethod(), context.GetPath(), context.GetParams(), context.GetValues()
		logrus.Debugf(">>>>>>>>>>> Handler => Method: %s , Path: %s , Params: %#v ,Values: %#v", method, path, params, values)

		// 当前打开的界面
		menu := context.GetUsingMenu()

		// 不记录：未指定菜单和查询操作日志
		if !strings.EqualFold(path, "/api/sys/operate_logs") {
			isLog := false
			if len(menu) > 4 {
				isLog = true
			} else if strings.EqualFold(path, "/api/sys/login") {
				if strings.EqualFold(params["method"], "ByPassword") {
					menu = "密码登录"
					isLog = true
				} else if strings.EqualFold(params["method"], "ChangePassword") {
					menu = "修改密码"
					isLog = true
				} else if strings.EqualFold(params["method"], "ByToken") {
					if len(params["byPassword"]) < 1 {
						menu = "Token登录"
						isLog = true
					}
				}
			}

			// 是否需要写操作日志
			if isLog {
				op = newOperate(db, context, menu)
				if err := op.SetParamsValues(params, values); err != nil {
					handlerError(op, w, err)
					return
				}
			}
		}

		// 是否为有效的Token
		if !context.HasInWhiteRoute(method, path, params, values) {
			err := context.Parse()
			if op != nil {
				op.SetUserData(context.GetUserId(), context.GetUserCode(), context.GetUserName(), context.GetDepartId(), context.GetDepartCode(), context.GetDepartName())
			}

			if err != nil {
				handlerError(op, w, errors.New("[PHOENIX_TOKEN_EXPIRE]"))
				return
			}
		}

		caller, ee := getCaller(r, md)
		if ee != nil {
			handlerError(op, w, ee)
			return
		}

		// 开启事务
		tx, err := db.Begin()
		if err != nil {
			handlerError(op, w, err)
			return
		}

		outs := caller.Call([]reflect.Value{reflect.ValueOf(tx), reflect.ValueOf(context)})

		var data, eer interface{}
		if len(outs) == 1 {
			eer = outs[0].Interface()
		} else {
			data, eer = outs[0].Interface(), outs[1].Interface()
		}

		if eer != nil {
			if err := tx.Rollback(); err != nil {
				logrus.Errorf("Request Rollback Failure :: %s", err.Error())
				return
			}

			handlerError(op, w, eer.(error))
		} else {
			if err := tx.Commit(); err != nil {
				handlerError(op, w, err)
				return
			}

			// 是否登录
			if strings.EqualFold(path, "/api/sys/login") {
				if strings.EqualFold(params["method"], "ByPassword") || strings.EqualFold(params["method"], "ChangePassword") {
					cookies := data.(map[string]string)

					expire, err := strconv.Atoi(cookies["expire"])
					if err != nil {
						logrus.Errorf("Get Cookie Expire Time Failure %s", err.Error())
					} else {
						setCookie(w, "PHOENIX_LOGIN_TOKEN", cookies["token"], time.Unix(int64(expire), 0))

						setCookie(w, "PHOENIX_USER_ID", cookies["user_id"], time.Unix(int64(expire), 0))
						setCookie(w, "PHOENIX_USER_CODE", cookies["user_code"], time.Unix(int64(expire), 0))
						setCookie(w, "PHOENIX_USER_NAME", base64.StdEncoding.EncodeToString([]byte(cookies["user_name"])), time.Unix(int64(expire), 0))
						setCookie(w, "PHOENIX_DEPART_ID", cookies["depart_id"], time.Unix(int64(expire), 0))
						setCookie(w, "PHOENIX_DEPART_CODE", cookies["depart_code"], time.Unix(int64(expire), 0))
						setCookie(w, "PHOENIX_DEPART_NAME", base64.StdEncoding.EncodeToString([]byte(cookies["depart_name"])), time.Unix(int64(expire), 0))
					}

					if op != nil {
						op.SetUserData(cookies["user_id"], cookies["user_code"], cookies["user_name"], cookies["depart_id"], cookies["depart_code"], cookies["depart_name"])
					}

					data = map[string]string{"status": "success"}
				}
			}

			if data == nil {
				if op != nil {
					op.success("不可显示内容")
				}
			} else {
				bs, err := json.Marshal(data)
				if err != nil {
					handlerError(op, w, err)
					return
				}

				if op != nil {
					op.success(string(bs))
				}

				if _, err := w.Write(bs); err != nil {
					logrus.Errorf("Response Write JSON Failure %s", err.Error())
				}
			}
		}
	})
}

// 根据请求规则，判断该调用哪个函数
func getCaller(r *http.Request, md interface{}) (reflect.Value, error) {
	mth := strings.Title(strings.ToLower(r.Method))
	xth := r.URL.Query().Get("method")

	// 根据调用指定的方法
	get := reflect.ValueOf(md).MethodByName(fmt.Sprintf("%s%s", mth, xth))
	if get.IsValid() {
		return get, nil
	}

	// Any?
	any := reflect.ValueOf(md).MethodByName("Any")
	if any.IsValid() {
		return any, nil
	}

	return reflect.Value{}, fmt.Errorf("没找到与请求方法对应的执行方法 %s%s ", mth, xth)
}

func handlerError(op *Operate, w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	// 记录错误日志
	logrus.Error(err)
	if op != nil {
		op.error(err)
	}

	bs, erx := json.Marshal(map[string]interface{}{"status": "error", "error": err.Error()})
	if erx != nil {
		logrus.Errorf("Handler Error Failure %s", erx.Error())
		return
	}

	if _, erv := w.Write(bs); erv != nil {
		logrus.Errorf("HTTP Write Failure %s", erv.Error())
	}
}
