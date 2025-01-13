package main

import (
	"database/sql"
	"fmt"
	"github.com/antonfisher/nested-logrus-formatter"
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
	var scripts = map[string]string{
		"dvdjnh2c9k35y7v49p1s7qftj7oxxdbt": `
var params = ctx.Params();

var select = "SELECT M2.ly, T.ldbh AS ldbh, T.rkrq AS rkrq,\n\t" + 
	"       M1.htbh, M1.khbh AS khbh, M1.khmc AS khmc,\n\t" + 
	"       T.gcbh AS gcbh, X1.gcmc AS gcmc, T.ckbh AS ckbh, X2.ckmc AS ckmc, T.kwbh AS kwbh, X3.kwmc AS kwmc,\n\t" + 
	"       T.wzbh AS wzbh, M2.wzmc AS wzmc, M2.ggxh AS ggxh, M2.jldw AS jldw, M2.sccjmc AS sccjmc,\n\t" + 
	"       M2.sccjmc AS sccjmc, M2.clph AS clph, M2.scrq AS scrq,\n\t" + 
	"       T.dj AS dj, T.sssl AS sssl, T.rkje AS rkje,\n\t" + 
	"       M1.cgy AS cgy, M1.kdrq AS kdrq, M2.jyry AS jyry, M2.jyrq AS jyrq, M2.bgy AS bgy, M2.rkrq AS mxrkrq";
	
var from = "FROM JZWZ_WZYE T\n\t" + 
	"       INNER JOIN JZWZ_WZRKDWJ M1 ON M1.id = T.wzrkd_id\n\t" + 
	"       INNER JOIN JZWZ_WZRKDWJMX M2 ON M2.id = T.wzrkdmx_id\n\t" + 
	"       LEFT JOIN JZMD_GCDM X1 ON X1.gcbh = T.gcbh\n\t" + 
	"       LEFT JOIN JZMD_CKDM X2 ON X2.ckbh = T.ckbh\n\t" + 
	"       LEFT JOIN JZMD_KWDM X3 ON X3.ckdm_id = X2.id AND X3.kwbh = T.kwbh";

var where = "WHERE 1 = 1";

if (_.has(params,"start_date") && !_.isEmpty(params["start_date"])) {
   where = where + " AND T.rkrq >= '"+params["start_date"]+"'";
}

if (_.has(params,"end_date") && !_.isEmpty(params["end_date"])) {
   where = where + " AND T.rkrq <= '"+params["end_date"]+"'";
}
	
var orderBy = "ORDER BY T.ldbh ASC, T.wzbh ASC, T.order_ ASC";

delete params.start_date;
delete params.end_date;

sql.Select(select, from, where, orderBy);
`,
		"dvdmcy17pbm7a5nx9ytsruvkwkcnfzhj": `
var params = ctx.Params();

var select = "SELECT T.ly AS ly, T.ldbh AS ldbh, T.llrq AS llrq, M3.ldbh AS rkldbh, \n\t" +  
	"       XM1.gcbh AS sq_gcbh, XM1.gcmc AS sq_gcmc,XM3.gcbh AS sf_gcbh, XM3.gcmc AS sf_gcmc,\n\t" +  
	"       T.ckbh AS ckbh, X2.ckmc AS ckmc, T.kwbh AS kwbh, X3.kwmc AS kwmc,\n\t" +  
	"       T.wzbh AS wzbh, M2.wzmc AS wzmc, M2.ggxh AS ggxh, M2.wzph AS wzph, M2.bzdh AS bzdh, M2.jldw AS jldw,\n\t" +  
	"       T.qls AS qls, T.dj AS dj, T.sfs AS sfs, T.ckje AS ckje,\n\t" +  
	"       M1.sqry AS sqry, M1.sqbm AS sqbm, M1.kdrq AS sqrq, M1.bmld AS bmld, M1.bmld_shrq AS bmld_shrq,\n\t" +  
	"       M2.qx AS qx, M2.llrq AS mxllrq, M2.lly AS lly, M2.bgy AS bgy";
	
var from = "FROM JZWZ_WZCK T\n\t" +  
	"       INNER JOIN JZWZ_WZLLSQWJ M1 ON M1.id = T.sq_id\n\t" +  
	"       INNER JOIN JZWZ_WZLLSQWJMX M2 ON M2.id = T.sqmx_id\n\t" + 
	"       INNER JOIN JZWZ_WZYE M3 ON M3.id = T.ye_id\n\t" +  
	"       LEFT JOIN JZMD_GCDM XM1 ON XM1.gcbh = M1.gcbh\n\t" +  
	"       LEFT JOIN JZMD_GCDM XM3 ON XM3.gcbh = M3.gcbh\n\t" +  
	"       LEFT JOIN JZMD_CKDM X2 ON X2.ckbh = T.ckbh\n\t" +  
	"       LEFT JOIN JZMD_KWDM X3 ON X3.ckdm_id = X2.id AND X3.kwbh = T.kwbh";

var where = "WHERE 1 = 1";

if (_.has(params,"start_date") && !_.isEmpty(params["start_date"])) {
   where = where + " AND T.llrq >= '"+params["start_date"]+"'";
}

if (_.has(params,"end_date") && !_.isEmpty(params["end_date"])) {
   where = where + " AND T.llrq <= '"+params["end_date"]+"'";
}

var orderBy = "ORDER BY T.ldbh ASC, T.wzbh ASC, T.order_ ASC";

delete params.start_date;
delete params.end_date;

sql.Select(select, from, where, orderBy);
`}

	for id, script := range scripts {
		_, err := db.Exec("UPDATE sys_data_service SET source_ = ? WHERE id = ?", script, id)
		if err != nil && err == sql.ErrNoRows {
			return fmt.Errorf("执行更新脚本ID %q 出现错误：%s\n", id, err.Error())
		}
	}

	return nil
}
