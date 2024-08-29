package vm

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/rujs/jsvm/sql/save"
)

func (s *Sql) Save(table string, values map[string]string) (result map[string]interface{}) {
	switch values["operation"] {
	case "insert":
		result = vm.Insert(s.tx, s.ctx, table, values)
	case "update":
		result = vm.Update(s.tx, s.ctx, table, values)
	case "delete":
		result = vm.Delete(s.tx, s.ctx, table, values)
	case "order":
		result = vm.Order(s.tx, s.ctx, table, values)
	default:
		logrus.Panic(fmt.Sprintf("无法识别的操作类型 %q", values["operation"]))
	}

	return
}
