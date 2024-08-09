package asql

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func Exec(tx *sql.Tx, query string, args ...interface{}) (int64, error) {
	logrus.Debugf("%s %s", FnArgs(args...), query)

	res, err := tx.Exec(query, args...)
	if err != nil {
		return 0, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return n, nil
}
