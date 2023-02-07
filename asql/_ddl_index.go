package asql

//
//import (
//	"fmt"
//	"github.com/sirupsen/logrus"
//	"phoenix/cm"
//	"strings"
//)
//
//func (d *DDL) SyncIndexes(present map[string]string) error {
//	res, err := Select(d.tx, fmt.Sprintf("SHOW INDEX FROM %s", d.table))
//	if err != nil {
//		return err
//	}
//
//	logrus.Debugf("Get Table's Indexes %#v", res)
//
//	indexes := make(map[string]string)
//	for _, s := range res {
//		code, noUnique, column := s["Key_name"], s["Non_unique"], s["Column_name"]
//
//		// 主键忽略
//		if strings.EqualFold(code, "PRIMARY") {
//			continue
//		}
//
//		index, ok := indexes[code]
//		if !ok {
//			index = fmt.Sprintf("%s|%s", noUnique, column)
//		} else {
//			index = strings.Join([]string{index, column}, ",")
//		}
//
//		indexes[code] = index
//	}
//
//	logrus.Debugf("Your's Indexes %#v", present)
//	logrus.Debugf("Table's Indexes %#v", indexes)
//	added, changed, removed := cm.CompareMap(indexes, present)
//
//	// 将更新的外键分别放入添加和删除队列，先删除再添加
//	for key, value := range changed {
//		removed[key] = value
//		added[key] = value
//	}
//
//	// 移除
//	for key := range removed {
//		query := fmt.Sprintf("ALTER TABLE %s DROP INDEX %s", d.table, key)
//		if err := Exec(d.tx, query); err != nil {
//			logrus.Errorf("SyncIndexes DROP INDEX Failure :: %s", err.Error())
//			return err
//		}
//	}
//
//	// 添加
//	for key := range added {
//		ss := strings.Split(present[key], "|")
//		if len(ss) != 2 {
//			panic("unreachable")
//		}
//
//		index := "INDEX"
//		if strings.EqualFold(ss[0], "0") {
//			index = "UNIQUE INDEX"
//		}
//
//		query := fmt.Sprintf("ALTER TABLE %s ADD %s %s (%s)", d.table, index, key, ss[1])
//		if err := Exec(d.tx, query); err != nil {
//			logrus.Errorf("SyncIndexes ADD INDEX Failure :: %s", err.Error())
//			return err
//		}
//	}
//
//	return nil
//}
