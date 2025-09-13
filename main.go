package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/antonfisher/nested-logrus-formatter"
	"io"
	"io/fs"
	"log"
	"path/filepath"
	"sort"
	"strings"

	//_ "dm"                               // 达梦 驱动
	_ "github.com/denisenkom/go-mssqldb" // SQL Server 驱动
	//_ "github.com/go-sql-driver/mysql"   // MySQL 驱动
	"github.com/sirupsen/logrus"
	"go-phoenix/base"
	"go-phoenix/handle"
	"go-phoenix/xsys"
	"go-phoenix/xwf"
	"net"
	"net/http"
	"os"
)

func main() {
	// 默认的日志级别
	logrus.SetLevel(logrus.TraceLevel)

	// 设置日志输出样式
	logrus.SetFormatter(&formatter.Formatter{TimestampFormat: "2006-01-02 15:04:05", HideKeys: true})

	// 获取当前目录
	dir, err := os.Getwd()
	if err != nil {
		logrus.Fatalf("os.Getwd() Failure :: %s", err.Error())
	}

	if err := base.InitConfig(dir); err != nil {
		logrus.Fatalf("base.InitConfig() Failure :: %s", err.Error())
	}

	// 数据库链接
	db, err := sql.Open(base.Config.DBDriver, base.Config.DBDataSource)
	if err != nil {
		logrus.Fatalf("sql.Open() Failure :: %s", err.Error())
	}

	// Ping ...
	if err := db.Ping(); err != nil {
		logrus.Fatalf("db.Ping() Failure :: %s", err.Error())
	}
	logrus.Info("连接数据库成功 ...")

	// 加载SQL脚本
	if err := loadScripts(db, dir); err != nil {
		logrus.Fatalf("loadScripts() Failure :: %s", err.Error())
	}
	log.Printf("当前软件版本为 %s >>>>>>\n", "2025.04.16")

	// 静态文件
	http.Handle("/", http.FileServer(http.Dir("www")))

	http.Handle("/api/sys", handle.Handler(db, &xsys.Sys{}))                                     // 加载系统（登录前）
	http.Handle("/api/sys/system", handle.Handler(db, &xsys.SysSystem{}))                        // 系统信息
	http.Handle("/api/sys/setting", handle.Handler(db, &xsys.SysSetting{}))                      // 系统设置
	http.Handle("/api/sys/roles", handle.Handler(db, &xsys.SysRoles{}))                          // 角色
	http.Handle("/api/sys/departs", handle.Handler(db, &xsys.SysDeparts{}))                      // 部门
	http.Handle("/api/sys/users", handle.Handler(db, &xsys.SysUsers{}))                          // 用户
	http.Handle("/api/sys/tables", handle.Handler(db, &xsys.SysTables{}))                        // 数据库表
	http.Handle("/api/sys/table_columns", handle.Handler(db, &xsys.SysTableColumns{}))           // 数据库表字段
	http.Handle("/api/sys/dict_kinds", handle.Handler(db, &xsys.SysDictKinds{}))                 // 数据字典
	http.Handle("/api/sys/dict_items", handle.Handler(db, &xsys.SysDictItems{}))                 // 数据字典项
	http.Handle("/api/sys/auto_nos", handle.Handler(db, &xsys.SysAutoNos{}))                     // 数据字典项
	http.Handle("/api/sys/auto_no_kinds", handle.Handler(db, &xsys.SysAutoNoKinds{}))            // 自动编码
	http.Handle("/api/sys/auto_no_items", handle.Handler(db, &xsys.SysAutoNoItems{}))            // 自动编码项
	http.Handle("/api/sys/menus", handle.Handler(db, &xsys.SysMenus{}))                          // 菜单
	http.Handle("/api/sys/role_menus", handle.Handler(db, &xsys.SysRoleMenus{}))                 // 角色关联菜单
	http.Handle("/api/sys/organization_roles", handle.Handler(db, &xsys.SysOrganizationRoles{})) // 组织关联角色
	http.Handle("/api/sys/operate_logs", handle.Handler(db, &xsys.SysOperateLogs{}))             // 操作日志
	http.Handle("/api/sys/data_service", handle.Handler(db, &xsys.SysDataService{}))             // 用户数据服务
	http.Handle("/api/sys/docs", handle.Handler(db, &xsys.SysDocs{}))                            // 文档
	http.Handle("/api/sys/time_tasks", handle.Handler(db, &xsys.SysTimeTasks{}))                 // 定时任务

	http.Handle("/api/sys/home_menus", handle.Handler(db, &xsys.SysHomeMenus{})) // 首页常用功能
	http.Handle("/api/wf/diagrams", handle.Handler(db, &xwf.Diagrams{}))         // 流程图
	http.Handle("/api/wf/flows", handle.Handler(db, &xwf.Flows{}))               // 流程执行

	//http.Handle("/api/sys/table_foreign_keys", handle.Handler(db, &xsys.SysTableForeignKeys{}))  // 数据库表外键
	//http.Handle("/api/sys/table_indexes", handle.Handler(db, &xsys.SysTableIndexes{}))           // 数据库表索引
	//http.Handle("/api/sys/ui_widget", handle.Handler(db, &xsys.SysUIWidget{}))                   // UI组件设计

	addr := net.JoinHostPort(base.Config.Host, base.Config.Port)
	logrus.Infof("HTTP服务器监听地址: %s ......", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logrus.Errorf("Listen Failure %s", err.Error())
	}
}

func loadScripts(db *sql.DB, dir string) error {
	root, scripts := filepath.Join(dir, "scripts"), make([]string, 0)
	if err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		scripts = append(scripts, path)

		return err
	}); err != nil {
		if os.IsNotExist(err) {
			logrus.Infof("脚本目录%q不存在 >>>>>>", root)
			return nil
		}

		return fmt.Errorf("遍历脚本目录%q出现异常: %s", root, err.Error())
	}

	sort.Strings(scripts)
	for i, script := range scripts {
		buf := new(bytes.Buffer)

		f, err := os.Open(script)
		if err != nil {
			return fmt.Errorf("读取脚本文件%q出现异常: %s", strings.ReplaceAll(script, root, ""), err.Error())
		}

		if _, err := io.Copy(buf, f); err != nil {
			return fmt.Errorf("读取脚本文件%q出现异常: %s", strings.ReplaceAll(script, root, ""), err.Error())
		}

		if _, err := db.Exec(buf.String()); err != nil {
			return fmt.Errorf("执行脚本文件%q出现异常: %s", strings.ReplaceAll(script, root, ""), err.Error())
		}

		if err := f.Close(); err != nil {
			return err
		}

		logrus.Infof("【%d】加载脚本文件%q ...", i+1, strings.ReplaceAll(script, root, ""))
	}
	return nil
}
