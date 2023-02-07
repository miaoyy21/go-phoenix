package asql

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func Update(tx *sql.Tx, query string, args ...interface{}) error {
	logrus.Debugf("Update SQL with Arguments [%s] %s", Arguments(args...), query)

	if _, err := tx.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
