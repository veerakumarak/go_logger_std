package std

import (
	logger "github.com/veerakumarak/go_logger_core"
	"log"
)

type std struct {
	l     *log.Logger
	level int32
}

func (s std) Debug(msg string) {
	s.l.Println(msg)
}

func (s std) Debugf(format string, args ...interface{}) {
	s.l.Printf(format, args)
}

func (s std) Info(msg string) {
	s.l.Println(msg)
}

func (s std) Infof(format string, args ...interface{}) {
	s.l.Printf(format, args)
}

func (s std) Warn(msg string) {
	s.l.Println(msg)
}

func (s std) Warnf(format string, args ...interface{}) {
	s.l.Printf(format, args)
}

func (s std) Error(msg string) {
	s.l.Println(msg)
}

func (s std) Errorf(format string, args ...interface{}) {
	s.l.Printf(format, args)
}

func (s std) With(args ...interface{}) logger.Logger {
	//TODO implement me
	panic("implement me")
}

func NewLogger() logger.Logger {
	return std{l: log.Default(), level: 0}
}
