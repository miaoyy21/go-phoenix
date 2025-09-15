package test

import (
	"database/sql"
	"go-phoenix/base"
	"log"
	"os"
	"path/filepath"
)

func GetDB() *sql.DB {

	// 获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd() Failure :: %s", err.Error())
	}

	baseDir, _ := filepath.Split(dir)

	if err := base.InitConfig(baseDir); err != nil {
		log.Fatalf("base.InitConfig() Failure :: %s", err.Error())
	}

	// 数据库链接
	db, err := sql.Open(base.Config.DBDriver, base.Config.DBDataSource)
	if err != nil {
		log.Fatalf("sql.Open() Failure :: %s", err.Error())
	}

	// Ping ...
	if err := db.Ping(); err != nil {
		log.Fatalf("db.Ping() Failure :: %s", err.Error())
	}

	log.Println("连接数据库成功 ...")

	return db
}
