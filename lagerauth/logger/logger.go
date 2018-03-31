package logger

import (
	"database/sql"
	"fmt"
	"io"
	"sync"
	"time"
)

type Logger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
}

type Options struct {
	Database *sql.DB
	Writer   io.Writer
}

type log struct {
	Options
	mu sync.Mutex
}

func New(opts Options) Logger {
	return &log{Options: opts, mu: sync.Mutex{}}
}

func (l *log) Debug(format string, values ...interface{}) {
	logger(l, "DEBUG", format, values...)
}

func (l *log) Info(format string, values ...interface{}) {
	logger(l, "INFO", format, values...)
}

func (l *log) Warn(format string, values ...interface{}) {
	logger(l, "WARN", format, values...)
}

func (l *log) Error(format string, values ...interface{}) {
	logger(l, "ERROR", format, values...)
}

func (l *log) Fatal(format string, values ...interface{}) {
	logger(l, "FATAL", format, values...)
}

func logger(l *log, level, format string, values ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	time := time.Now()

	if l.Writer != nil {
		loggerWriter(l.Writer, time, level, format, values...)
	}

	if l.Database != nil {
		loggerQuery(l.Database, time, level, format, values...)
	}
}

func loggerQuery(db *sql.DB, time time.Time, level, format string, values ...interface{}) error {
	message := getFormatedMessage(format, values...)
	_, err := db.Exec(`insert into logs (level, message) values (?, ?)`, level, message)
	return err
}

func loggerWriter(w io.Writer, time time.Time, level, format string, values ...interface{}) {
	message := getFormatedMessage(format, values...)
	message = fmt.Sprintf("%s %s: %s", time.Format("2006/04/02 15:04:05"), level, message)
	fmt.Fprintln(w, message)
}

func getFormatedMessage(format string, values ...interface{}) string {
	return fmt.Sprintf(format, values...)
}
