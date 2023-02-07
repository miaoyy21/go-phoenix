package handle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go-phoenix/asql"
	"strings"
	"time"
)

type Operate struct {
	db *sql.DB
	id string

	ip     string
	size   string
	agent  string
	method string
	menu   string
	path   string
	params string
	values string

	userId     string
	userCode   string
	userName   string
	departId   string
	departCode string
	departName string

	start    time.Time
	end      time.Time
	duration int64

	status  string
	message string
}

func newOperate(db *sql.DB, ctx *Context, menu string) *Operate {

	op := &Operate{
		db: db,
		id: asql.GenerateId(),

		ip:     ctx.GetIP(),
		size:   ctx.GetSize(),
		method: ctx.GetMethod(),
		menu:   menu,
		path:   ctx.GetPath(),

		start: time.Now(),
	}

	// 只有登录请求才记录客户端 User-Agent 信息
	logrus.Debugf("Path is %s", op.path)
	if strings.EqualFold(op.path, "/api/sys/login") {
		op.agent = ctx.UserAgent()
	}

	return op
}

func (op *Operate) SetParamsValues(params map[string]string, values map[string]string) error {
	// Params
	bs, err := json.Marshal(params)
	if err != nil {
		return err
	}

	if len(bs) <= 1024 {
		op.params = string(bs)
	} else {
		op.params = formatSize(len(bs))
	}

	// Values
	var sValues string

	if strings.EqualFold(op.path, "/api/sys/login") {
		if params["method"] == "ByPassword" || params["method"] == "ChangePassword" {
			sValues = "****** 因安全原因，系统隐藏用户敏感信息 ******"
		}
	}

	if len(sValues) < 1 {
		bs, err := json.Marshal(values)
		if err != nil {
			return err
		}

		sValues = string(bs)
		if len(sValues) > 1024 {
			sValues = formatSize(len(bs))
		}
	}

	op.values = sValues

	return nil
}

func (op *Operate) SetUserData(userId, userCode, userName, departId, departCode, departName string) {

	op.userId = userId
	op.userCode = userCode
	op.userName = userName
	op.departId = departId
	op.departCode = departCode
	op.departName = departName
}

func (op *Operate) error(err error) {
	op.status = "error"
	op.message = err.Error()

	op.save()
}

func formatSize(size int) string {
	if size < 1<<10 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1<<20 {
		return fmt.Sprintf("%.2f KB", float64(size)/(1<<10))
	} else if size < 1<<30 {
		return fmt.Sprintf("%.2f MB", float64(size)/(1<<20))
	}

	return fmt.Sprintf("%.2f GB", float64(size)/(1<<30))
}

func (op *Operate) success(msg string) {
	op.status = "success"
	if op.method == "GET" {
		op.message = formatSize(len(msg))
	} else {
		op.message = msg
	}

	op.save()
}

func (op *Operate) save() {
	op.end = time.Now()
	op.duration = op.end.Sub(op.start).Milliseconds()

	query := "INSERT INTO sys_operate_log(id, ip_, size_, agent_, method_, menu_id_, path_, params_, values_, " +
		"	user_id_, user_code_, user_name_, depart_id_, depart_code_, depart_name_, start_, end_, duration_, status_, message_) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	args := []interface{}{op.id, op.ip, op.size, op.agent, op.method, op.menu, op.path, op.params, op.values, op.userId, op.userCode, op.userName, op.departId, op.departCode, op.departName, op.start, op.end, op.duration, op.status, op.message}
	if _, err := op.db.Exec(query, args...); err != nil {
		logrus.Errorf("Write Operate Log Failure :: %s", err.Error())
	}
}
