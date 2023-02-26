package handle

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"reflect"
	"runtime/debug"
	"strings"
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

		ctx := NewContext(db, r, w)
		method, path, params, values := ctx.GetMethod(), ctx.GetPath(), ctx.GetParams(), ctx.GetValues()
		logrus.Debugf(">>>>>>>>>>> Handler => Method: %s , Path: %s , Params: %#v ,Values: %#v", method, path, params, values)

		// 操作日志
		op = newOperate(db, ctx)

		// 是否为有效的Token
		if err := ctx.Parse(); err != nil {
			if !strings.EqualFold(path, "/api/sys") {
				handlerError(op, w, errors.New("[PHOENIX_TOKEN_EXPIRE]"))
				return
			}
		}

		// 设置接收参数和提交表单
		if err := op.refresh(); err != nil {
			handlerError(op, w, err)
			return
		}

		caller, err := getCaller(r, md)
		if err != nil {
			handlerError(op, w, err)
			return
		}

		// 开启事务
		tx, err := db.Begin()
		if err != nil {
			handlerError(op, w, err)
			return
		}

		// 业务逻辑调用
		outs := caller.Call([]reflect.Value{reflect.ValueOf(tx), reflect.ValueOf(ctx)})

		// 返回两个值：响应数据和错误信息
		result, failure := outs[0].Interface(), outs[1].Interface()
		if failure != nil {
			if err := tx.Rollback(); err != nil {
				logrus.Errorf("Request Rollback Failure :: %s", err.Error())
				return
			}

			handlerError(op, w, failure.(error))
		} else {
			if err := tx.Commit(); err != nil {
				handlerError(op, w, err)
				return
			}

			bs, err := json.Marshal(result)
			if err != nil {
				handlerError(op, w, err)
				return
			}

			if !strings.EqualFold(path, "/api/sys/operate_logs") {
				op.success(string(bs))
			}

			if _, err := w.Write(bs); err != nil {
				logrus.Errorf("Response Write JSON Failure %s", err.Error())
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

	// Some?
	some := reflect.ValueOf(md).MethodByName("Any")
	if some.IsValid() {
		return some, nil
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
