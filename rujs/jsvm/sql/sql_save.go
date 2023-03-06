package vm

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/rujs/jsvm/sql/save"
)

func (s *Sql) Save(table string) (result map[string]interface{}) {
	values := s.ctx.Values()

	switch values["operation"] {
	case "insert":
		result = vm.Insert(s.tx, s.ctx, table)
	case "update":
		result = vm.Update(s.tx, s.ctx, table)
	case "delete":
		result = vm.Delete(s.tx, s.ctx, table)
	case "order":
		result = vm.Order(s.tx, s.ctx, table)
	default:
		logrus.Panic(fmt.Sprintf("无法识别的操作类型 %q", values["operation"]))
	}

	return
}
