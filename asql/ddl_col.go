package asql

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/base"
	"strings"
)

func (d *DDL) SyncColumns(present map[string]string, sorts []string) error {
	if len(d.columns) < 1 {
		buf := &bytes.Buffer{}

		// 新创建表
		buf.WriteString(fmt.Sprintf("CREATE TABLE %s (\n", d.table))
		buf.WriteString("\tid VARCHAR(256) NOT NULL,\n")

		// 字段
		for _, key := range sorts {
			if strings.EqualFold(key, "id") {
				continue
			}

			buf.WriteString(fmt.Sprintf("\t%s %s DEFAULT NULL,\n", key, present[key]))
		}

		// 主键
		buf.WriteString("\tPRIMARY KEY (id)\n")
		buf.WriteString(")")

		/****** 创建数据库表 ******/
		if _, err := Exec(d.tx, buf.String()); err != nil {
			return err
		}
	} else {
		added, changed, removed := base.CompareMap(d.columns, present)

		// 添加（按照给定的顺序）
		for _, key := range sorts {
			value, ok := added[key]
			if !ok {
				continue
			}

			if _, err := Exec(d.tx, fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s NULL;", d.table, key, value)); err != nil {
				logrus.Errorf("SyncColumns ADD Failure :: %s", err.Error())
				return err
			}
		}

		// 更新
		for key, value := range changed {
			if _, err := Exec(d.tx, fmt.Sprintf("ALTER TABLE %s CHANGE COLUMN %s %s %s NULL;", d.table, key, key, value)); err != nil {
				logrus.Errorf("SyncColumns CHANGE Failure :: %s", err.Error())
				return err
			}
		}

		// 移除
		for key, value := range removed {
			if _, err := Exec(d.tx, fmt.Sprintf("ALTER TABLE %s CHANGE COLUMN %s _%s %s NULL;", d.table, key, key, value)); err != nil {
				logrus.Errorf("SyncColumns DROP Failure :: %s", err.Error())
				return err
			}
		}
	}

	return nil
}
