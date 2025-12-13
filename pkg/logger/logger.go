package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func New() *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
	}
}

func (l *Logger) Info(msg string) {
	l.Println("[INFO]", msg)
}

func (l *Logger) Error(err error) {
	l.Println("[ERROR]", err.Error())
}

func (l *Logger) Fatal(err error) {
	l.Println("[FATAL]", err.Error())
	os.Exit(1)
}
