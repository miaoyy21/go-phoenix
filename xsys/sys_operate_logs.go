package xsys

import (
	"database/sql"
	"go-phoenix/handle"
)

type SysOperateLogs struct {
}

func (o *SysOperateLogs) Get(tx *sql.Tx, ctx *handle.Context) (interface{}, error) {
	pagination := handle.NewPagination(ctx, []string{"start_ DESC"})

	query := `
		SELECT sys_operate_log.id, sys_operate_log.user_name_, sys_operate_log.method_, 
			CASE WHEN sys_menu.id IS NULL THEN sys_operate_log.menu_id_ ELSE sys_menu.name_ END AS menu_name_, 
			sys_operate_log.path_, sys_operate_log.params_, sys_operate_log.values_, 
			sys_operate_log.start_, sys_operate_log.end_, sys_operate_log.duration_, 
			sys_operate_log.status_, sys_operate_log.message_, sys_operate_log.ip_, 
			sys_operate_log.agent_, sys_operate_log.depart_name_, sys_operate_log.user_code_
		FROM sys_operate_log LEFT JOIN sys_menu ON sys_operate_log.menu_id_ = sys_menu.id
	`
	if err := pagination.SetData(tx, query); err != nil {
		return nil, err
	}

	return pagination, nil
}
