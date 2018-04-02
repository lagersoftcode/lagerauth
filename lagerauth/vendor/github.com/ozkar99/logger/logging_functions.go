package logger

import (
	"database/sql"
	"fmt"
	"io"
	"time"
)

type options struct {
	database *sql.DB
	writers  []io.Writer
}

func log(l *logger, level, format string, values ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	time := time.Now().UTC()

	for _, writer := range l.writers {
		logToWriter(writer, time, level, format, values...)
	}

	if l.database != nil {
		logToDatabase(l.database, time, level, format, values...)
	}
}

func logToDatabase(db *sql.DB, time time.Time, level, format string, values ...interface{}) error {
	message := getFormatedMessage(format, values...)
	_, err := db.Exec(`insert into logs (level, message) values (?, ?)`, level, message)
	return err
}

func logToWriter(w io.Writer, time time.Time, level, format string, values ...interface{}) {
	message := getFormatedMessage(format, values...)
	message = fmt.Sprintf("%s %s: %s", time.Format("2006/04/02 15:04:05"), level, message)
	fmt.Fprintln(w, message)
}

func getFormatedMessage(format string, values ...interface{}) string {
	return fmt.Sprintf(format, values...)
}
