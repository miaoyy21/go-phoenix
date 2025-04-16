package asql

import (
	"database/sql"
	"log"
)

func Insert(tx *sql.Tx, query string, args ...interface{}) error {
	//logrus.Debugf("%s %s", FnArgs(args...), query)
	log.Printf("%s \n%s \n", FnArgs(args...), query)

	if _, err := tx.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
