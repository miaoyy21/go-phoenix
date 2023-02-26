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
	db  *sql.DB
	ctx *Context

	id string

	menu      string
	userAgent string
	params    string
	values    string

	start    time.Time
	end      time.Time
	duration int64

	status  string
	message string
}

func newOperate(db *sql.DB, ctx *Context) *Operate {
	op := &Operate{
		db:  db,
		ctx: ctx,
		id:  asql.GenerateId(),

		start: time.Now(),
	}

	return op
}

func (op *Operate) refresh() error {
	params, values := op.ctx.GetParams(), op.ctx.GetValues()

	// Params
	bsParams, err := json.Marshal(params)
	if err != nil {
		return err
	}
	op.params = formatBytesWithSize(bsParams)

	// Values
	bsValues, err := json.Marshal(values)
	if err != nil {
		return err
	}
	op.params = formatBytesWithSize(bsValues)

	// Using Menu
	menu, err := op.ctx.GetMenu()
	if err != nil {
		return err
	}
	op.menu = menu

	// User Agent
	if strings.EqualFold(op.ctx.GetPath(), "/api/sys") {
		op.userAgent = op.ctx.UserAgent()
	}

	return nil
}

func (op *Operate) error(err error) {
	op.status = "error"
	op.message = err.Error()

	op.save()
}

func formatBytesWithSize(bytes []byte) string {
	size := len(bytes)
	if len(bytes) <= 1024 {
		return string(bytes)
	}
	if size < 1<<10 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1<<20 {
		return fmt.Sprintf("%.2f KB", float64(size)/(1<<10))
	} else if size < 1<<30 {
		return fmt.Sprintf("%.2f MB", float64(size)/(1<<20))
	}

	return fmt.Sprintf("%.2f GB", float64(size)/(1<<30))
}

func (op *Operate) success(bytes []byte) {
	op.status = "success"
	op.message = formatBytesWithSize(bytes)

	op.save()
}

func (op *Operate) save() {
	op.end = time.Now()
	op.duration = op.end.Sub(op.start).Milliseconds()

	query := "INSERT INTO sys_operate_log(" +
		"	id, ip_, size_, agent_, method_, " +
		"	menu_id_, path_, params_, values_, " +
		"	user_id_, user_code_, user_name_, " +
		"	depart_id_, depart_code_, depart_name_, " +
		"	start_, end_, duration_, status_, message_) " +
		"VALUES (?,?,?,?,?, ?,?,?,?, ?,?,?, ?,?,?, ?,?,?,?,?)"
	args := []interface{}{
		op.id, op.ctx.GetIP(), op.ctx.GetSize(), op.userAgent, op.ctx.GetMethod(),
		op.menu, op.ctx.GetPath(), op.params, op.values,
		op.ctx.GetUserId(), op.ctx.GetUserCode(), op.ctx.GetUserName(),
		op.ctx.GetDepartId(), op.ctx.GetDepartCode(), op.ctx.GetDepartName(),
		op.start, op.end, op.duration, op.status, op.message}
	if _, err := op.db.Exec(query, args...); err != nil {
		logrus.Errorf("Write Operate Log Failure :: %s", err.Error())
	}
}
