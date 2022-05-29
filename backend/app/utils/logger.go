package utils

import (
	"log"
	"os"
)

/*
Handling logs
This functions are suitable for logging into the console using some
prefix to identify better the nature of the errors and logs.
*/
type logging struct {
	p *log.Logger
}

var Log *logging
var l logging

func InitLogs(prefix string) *log.Logger {
	l.p = log.New(os.Stdout, prefix, log.LstdFlags)
	l.Info("Initializing logs...")
	return l.p
}

func (lg *logging) Info(x ...interface{}) {
	s := "[INFO] "
	args := append([]interface{}{s}, x...)
	l.p.Println(args...)
}

func (lg *logging) Error(x ...interface{}) {
	s := "[ERROR]"
	args := append([]interface{}{s}, x...)
	l.p.Println(args...)
}

func (lg *logging) Debug(x ...interface{}) {
	s := "[DEBUG]"
	args := append([]interface{}{s}, x...)
	l.p.Println(args...)
}
