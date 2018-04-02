package sqlite

import "github.com/ozkar99/logger"

func init() {
	logger.SetCreateTableSQL(
		`create table if not exists logs (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			Level VARCHAR(10),
			Message VARCHAR(250),
			CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`)
}
