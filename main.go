package main

import (
	"database/sql"
	"fmt"
	"github.com/antonfisher/nested-logrus-formatter"
	"log"

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

	// 执行更新SQL脚本
	if err := runScripts(db); err != nil {
		logrus.Fatalf("runScripts() Failure :: %s", err.Error())
	}
	logrus.Info("执行更新脚本成功 ...")
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

func runScripts(db *sql.DB) error {
	if _, err := db.Exec("update JZMD_WZDM set jldw = '个' where wzbh in ('B01987','B01981','B01983','B01984','B01986','B01982','B01985') and jldw in ('2','4','8','10','20')"); err != nil {
		return fmt.Errorf("执行更新脚本1 出现错误：%s\n", err.Error())
	}

	if _, err := db.Exec("update JZWZ_WZRKDWJMX set jldw = '个' where wzbh in ('B01987','B01981','B01983','B01984','B01986','B01982','B01985') and jldw in ('2','4','8','10','20')"); err != nil {
		return fmt.Errorf("执行更新脚本2 出现错误：%s\n", err.Error())
	}

	if _, err := db.Exec("update JZWZ_WZHCDWJMX set jldw = '个' where wzbh in ('B01987','B01981','B01983','B01984','B01986','B01982','B01985') and jldw in ('2','4','8','10','20')"); err != nil {
		return fmt.Errorf("执行更新脚本3 出现错误：%s\n", err.Error())
	}

	if _, err := db.Exec("update JZWZ_WZLLSQWJMX set jldw = '个' where wzbh in ('B01987','B01981','B01983','B01984','B01986','B01982','B01985') and jldw in ('2','4','8','10','20')"); err != nil {
		return fmt.Errorf("执行更新脚本4 出现错误：%s\n", err.Error())
	}

	if _, err := db.Exec("ALTER TABLE JZWZ_WZRKDWJ ALTER COLUMN htbh varchar(256)"); err != nil {
		return fmt.Errorf("执行更新脚本5 出现错误：%s\n", err.Error())
	}

	if _, err := db.Exec("UPDATE sys_table_column SET type_ = 'VARCHAR(256)' WHERE id = 'dv9xszd2wmk2g9bm9xut2vw7sdgnhvkv'"); err != nil {
		return fmt.Errorf("执行更新脚本6 出现错误：%s\n", err.Error())
	}

	return nil
}
