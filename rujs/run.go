package rujs

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
	"github.com/sirupsen/logrus"
	"go-phoenix/handle"
	"go-phoenix/rujs/jsvm"
	"time"
)

var timeout = errors.New("timeout runtime")

func Run(tx *sql.Tx, ctx *handle.Context, source string, second int, reg func(*otto.Otto) error) (value otto.Value, err error) {
	start := time.Now()

	defer func() {
		duration := time.Since(start)
		if caught := recover(); caught != nil {
			if caught == timeout {
				err = fmt.Errorf("It is timeout %s, expect less than %s\n", duration, time.Duration(second*1e9))
			} else {
				if erx, ok := caught.(*logrus.Entry); ok {
					err = errors.New(erx.Message)
				} else {
					err = fmt.Errorf("%s", caught)
				}
			}

			return
		}
	}()

	js, err := jsvm.NewVM(tx, ctx, reg)
	if err != nil {
		return otto.NullValue(), fmt.Errorf("jsvm.NewVM() Failure :: %s", err.Error())
	}

	js.Interrupt = make(chan func(), 1)
	go func() {
		if second > 0 {
			time.Sleep(time.Duration(second) * time.Second)
			js.Interrupt <- func() {
				logrus.Panic(timeout)
			}
		}
	}()

	value, err = js.Run(source)
	if err != nil {
		return otto.NullValue(), fmt.Errorf("js.Run() Failure :: %s", err.Error())
	}

	return value, nil
}
