package asql
//
//import (
//	"fmt"
//	"github.com/sirupsen/logrus"
//	"phoenix/base"
//	"phoenix/cm"
//)
//
//type ForeignKey struct {
//	code             string `name:"code_"`
//	Column           string `name:"column_"`
//	ReferencedTable  string `name:"referenced_table_"`
//	ReferencedColumn string `name:"referenced_column_"`
//}
//
//func NewForeignKeys(ss []map[string]string) (map[string]ForeignKey, map[string]string, error) {
//	fks := make(map[string]ForeignKey)
//	cms := make(map[string]string)
//
//	for _, s := range ss {
//		var fk ForeignKey
//
//		if err := cm.ReflectToStruct(s, &fk); err != nil {
//			return nil, nil, err
//		}
//
//		fks[fk.code] = fk
//		cms[fk.code] = fmt.Sprintf("%s|%s|%s", fk.Column, fk.ReferencedTable, fk.ReferencedColumn)
//	}
//
//	return fks, cms, nil
//}
//
//func (d *DDL) SyncForeignKeys(present map[string]ForeignKey, sPresent map[string]string) error {
//	query := `
//			SELECT C.constraint_name AS code_,C.column_name AS column_,
//				   C.referenced_table_name AS referenced_table_,C.referenced_column_name AS referenced_column_
//			FROM information_schema.key_column_usage C
//					 INNER JOIN information_schema.tables T ON T.TABLE_NAME = C.TABLE_NAME
//					 INNER JOIN information_schema.referential_constraints R
//								ON R.table_name = C.table_name AND R.constraint_name = C.constraint_name AND R.referenced_table_name = C.referenced_table_name
//			WHERE C.referenced_table_name IS NOT NULL
//			  AND C.table_schema = ? AND C.table_name = ?
//		`
//	res, err := Select(d.tx, query, base.Config.Schema(), d.table)
//	if err != nil {
//		return err
//	}
//
//	foreignKeys, sForeignKeys, err := NewForeignKeys(res)
//	if err != nil {
//		return err
//	}
//
//	added, changed, removed := cm.CompareMap(sForeignKeys, sPresent)
//
//	// 将更新的外键分别放入添加和删除队列，先删除再添加
//	for key, value := range changed {
//		removed[key] = value
//		added[key] = value
//	}
//
//	// 移除
//	for key := range removed {
//		foreignKey := foreignKeys[key]
//
//		// 移除外键
//		query := fmt.Sprintf("ALTER TABLE %s DROP FOREIGN KEY %s", d.table, foreignKey.code)
//		if err := Exec(d.tx, query); err != nil {
//			logrus.Errorf("SyncForeignKeys DROP FOREIGN KEY Failure :: %s", err.Error())
//			return err
//		}
//
//		// 移除索引
//		query = fmt.Sprintf("ALTER TABLE %s DROP INDEX %s", d.table, foreignKey.code)
//		if err := Exec(d.tx, query); err != nil {
//			logrus.Errorf("SyncForeignKeys DROP INDEX Failure :: %s", err.Error())
//			return err
//		}
//	}
//
//	// 添加
//	for key := range added {
//		foreignKey := present[key]
//
//		query := fmt.Sprintf(`ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s)`,
//			d.table, foreignKey.code, foreignKey.Column, foreignKey.ReferencedTable, foreignKey.ReferencedColumn)
//		if err := Exec(d.tx, query); err != nil {
//			logrus.Errorf("SyncForeignKeys ADD Failure :: %s", err.Error())
//			return err
//		}
//	}
//
//	return nil
//}
