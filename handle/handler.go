package handle

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/base"
	"net/http"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
)

var rwMutex sync.RWMutex

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
		path, params, values := ctx.Path(), ctx.Params(), ctx.Values()
		if !(strings.EqualFold(path, "/api/sys") && strings.EqualFold(params["method"], "Sync")) {
			logrus.Debugf("[%s %q]: {Params: %s, Values: %s}", ctx.Method, path, base.MapString(params), base.MapString(values))
		}

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

		// 由于使用的是数据库事务，需要添加全局读写锁，保证数据更新的正确性，典型的应用场景：批量的自动编码
		if strings.EqualFold(r.Method, "GET") {
			rwMutex.RLock()
			defer rwMutex.RUnlock()
		} else {
			rwMutex.Lock()
			defer rwMutex.Unlock()
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
				if strings.EqualFold(path, "/api/sys") && strings.EqualFold(params["method"], "Sync") {
				} else {
					op.success(bs)
				}
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

	// Some Other?
	some := reflect.ValueOf(md).MethodByName("Any")
	if some.IsValid() {
		return some, nil
	}

	return reflect.Value{}, fmt.Errorf("没找到与请求方法对应的执行方法 %s%s ", mth, xth)
}

func handlerError(op *Operate, w http.ResponseWriter, msg error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	// 记录错误日志
	_, file, line, _ := runtime.Caller(1)
	logrus.Error(fmt.Errorf("%s:%d %s", strings.Split(file, "/go-phoenix/")[1], line, msg.Error()))
	if op != nil {
		op.error(msg)
	}

	bs, err := json.Marshal(map[string]interface{}{"status": "error", "error": msg.Error()})
	if err != nil {
		logrus.Errorf("Handler Error Failure %s", err.Error())
		return
	}

	if _, err := w.Write(bs); err != nil {
		logrus.Errorf("HTTP Write Failure %s", err.Error())
	}
}
