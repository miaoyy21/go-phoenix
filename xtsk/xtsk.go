package xtsk

import (
	"database/sql"
	"fmt"
	"go-phoenix/asql"
)

func Init(tx *sql.Tx) error {
	tasks, err := asql.Select(tx, "SELECT * FROM sys_time_task")
	if err != nil {
		return err
	}

	fmt.Printf("Tasks is %#v\n", tasks)
	return nil
}
