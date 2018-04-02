package logger

import "database/sql"

var createTableSQL string

func SetCreateTableSQL(sql string) {
	createTableSQL = sql
}

func createLogTableIfNotExists(db *sql.DB) {
	if db == nil {
		return
	}

	db.Exec(createTableSQL)
}
