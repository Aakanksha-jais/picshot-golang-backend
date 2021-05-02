package log

import (
	"fmt"
	"time"
)

type logLine struct {
	Level   level
	Time    time.Time
	Message interface{}
}

func (l logLine) TerminalOutput() string {
	return fmt.Sprintf("\u001B[%dm%s\u001B[0m [%s] %v\n", l.Level.color(), l.Level.String()[0:4], l.Time.Format("15:04:05"), l.Message)
}
