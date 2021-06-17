package log

import (
	"fmt"
	"runtime"
	"time"
)

type logLine struct {
	Level   level
	Time    time.Time
	Message interface{}
}

func (l logLine) TerminalOutput() string {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	mem := float32(m.Alloc) / float32(1024*1024)

	return fmt.Sprintf("\u001B[%dm%s\u001B[0m [%s] %v", l.Level.color(), l.Level.String()[0:4], l.Time.Format("15:04:05"), l.Message) +
		fmt.Sprintf("\n\t\u001B[%dm (Memory: %v MB GoRoutines: %v) \u001B[0m\n\n", 37, mem, runtime.NumGoroutine())
}
