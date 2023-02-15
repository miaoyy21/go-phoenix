package main

import (
	"database/sql"
	"fmt"
	"go-phoenix/base"
	"go-phoenix/xsys"
	"os"
	"runtime"

	"go-phoenix/handle"
	"go-phoenix/xwf"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	// 默认的日志级别
	logrus.SetLevel(logrus.TraceLevel)

	// 设置日志输出样式
	logrus.SetFormatter(
		&logrus.TextFormatter{EnvironmentOverrideColors: true,
			FullTimestamp:   true,
			PadLevelText:    true,
			TimestampFormat: "2006-01-02 15:04:05.000",
			CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				loc := strings.TrimPrefix(frame.File, base.Config.Dir()+string(os.PathSeparator))
				return ":", fmt.Sprintf("%s:%d", loc, frame.Line)
			},
		},
	)
	logrus.SetReportCaller(true)

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

	logrus.Info("Connect Database Successful ...")

	// 静态文件
	http.Handle("/", http.FileServer(http.Dir("www")))

	http.Handle("/api/sys/login", handle.Handler(db, &xsys.SysLogin{}))                          // 登录
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

	//http.Handle("/api/sys/table_foreign_keys", handle.Handler(db, &xsys.SysTableForeignKeys{}))  // 数据库表外键
	//http.Handle("/api/sys/table_indexes", handle.Handler(db, &xsys.SysTableIndexes{}))           // 数据库表索引
	//http.Handle("/api/sys/ui_widget", handle.Handler(db, &xsys.SysUIWidget{}))                   // UI组件设计

	http.Handle("/api/wf/diagrams", handle.Handler(db, &xwf.Diagrams{})) // 流程图
	http.Handle("/api/wf/flows", handle.Handler(db, &xwf.Flows{}))       // 流程执行

	// 路由白名单
	handle.AddWhiteRoute("GET", "/api/sys/dict_items", map[string]string{"scope": "ALL"}, map[string]string{})
	handle.AddWhiteRoute("GET", "/api/sys/setting", map[string]string{"scope": "LOGIN"}, map[string]string{})
	handle.AddWhiteRoute("GET", "/api/sys/users", map[string]string{"scope": "LOGIN"}, map[string]string{})
	handle.AddWhiteRoute("POST", "/api/sys/login", map[string]string{"method": "ByPassword"}, map[string]string{})

	logrus.Infof("Start Listen %s ......", "127.0.0.1:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		logrus.Errorf("Listen Failure %s", err.Error())
	}
}
