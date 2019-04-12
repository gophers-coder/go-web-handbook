package log

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	log.Logger
	Prefix string
	Flag   int
}

func NewLogger(io io.Writer, prefix string, flag int) *Logger {
	logger := log.New(io, prefix, flag)
	return &Logger{
		*logger, prefix, flag,
	}
}

var logger = NewLogger(os.Stdout, "Logger ", log.LstdFlags)

func (l *Logger) NewLogger(io io.Writer) {
	logger := log.New(io, l.Prefix, l.Flag)
	l.Logger = *logger
}

func (l *Logger) Panicf(format string, message interface{}) {
	l.Logger.Panicf(format, message)
}

func (l *Logger) Printf(format string, message interface{}) {
	l.Logger.Printf(format, message)
}
func (l *Logger) Fatalf(format string, message interface{}) {
	l.Logger.Fatalf(format, message)
}

func Panicf(format string, message interface{}) {
	logger.Panicf(format, message)
}

func Printf(format string, message interface{}) {
	logger.Printf(format, message)
}

func Fatalf(format string, message interface{}) {
	logger.Fatalf(format, message)
}
