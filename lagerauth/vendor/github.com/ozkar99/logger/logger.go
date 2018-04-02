package logger

import (
	"database/sql"
	"io"
	"sync"
)

type Logger interface {
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Warn(string, ...interface{})
	Error(string, ...interface{})
	Fatal(string, ...interface{})
}

type logger struct {
	options
	mu sync.Mutex
}

func New(database *sql.DB, writers ...io.Writer) Logger {
	createLogTableIfNotExists(database)
	return &logger{options: options{database, writers}, mu: sync.Mutex{}}
}

func (l *logger) Debug(format string, values ...interface{}) {
	log(l, "DEBUG", format, values...)
}

func (l *logger) Info(format string, values ...interface{}) {
	log(l, "INFO", format, values...)
}

func (l *logger) Warn(format string, values ...interface{}) {
	log(l, "WARN", format, values...)
}

func (l *logger) Error(format string, values ...interface{}) {
	log(l, "ERROR", format, values...)
}

func (l *logger) Fatal(format string, values ...interface{}) {
	log(l, "FATAL", format, values...)
}
