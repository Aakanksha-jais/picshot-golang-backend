package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"golang.org/x/crypto/ssh/terminal"
)

type Logger interface {
	// To register DEBUG Logs.
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	// To register INFO Logs.
	Info(args ...interface{})
	Infof(format string, args ...interface{})

	// To register WARN Logs.
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	// To register ERROR Logs.
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	// To register FATAL Logs.
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

type logger struct {
	level      level
	out        io.Writer
	isTerminal bool
}

// log does the actual logging. This function creates the log line and outputs it in color format
// in terminal context and gives out json in non terminal context.
func (l *logger) log(level level, format string, args ...interface{}) {
	if level < l.level {
		return
	}

	ll := logLine{
		Level: level,
		Time:  time.Now(),
	}

	switch {
	case len(args) == 1 && format == "":
		ll.Message = args[0]
	case len(args) != 1 && format == "":
		ll.Message = args
	case format != "":
		ll.Message = fmt.Sprintf(format+"", args...) // TODO - fix empty string
	}

	if l.isTerminal {
		fmt.Fprint(l.out, ll.TerminalOutput())
	} else {
		_ = json.NewEncoder(l.out).Encode(ll)
	}

}
func (l *logger) Debug(args ...interface{}) {
	l.log(DEBUG, "", args...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.log(INFO, "", args...)
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.log(WARN, "", args...)
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.log(ERROR, "", args...)
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.log(FATAL, "", args...)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.log(FATAL, format, args...)
}

func newLogger(level level) *logger {
	l := &logger{
		level: level,
		out:   os.Stdout,
	}

	l.isTerminal = checkIfTerminal(l.out)

	return l
}

func checkIfTerminal(out io.Writer) bool {
	switch v := out.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}

func NewLogger(level level) Logger {
	return newLogger(level)
}
