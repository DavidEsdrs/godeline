package logger

import "fmt"

type Logger struct {
	verbose bool
}

func NewLogger(v bool) Logger {
	return Logger{v}
}

func (l Logger) Debug(msg string, args ...any) {
	if l.verbose {
		fmt.Printf(msg, args...)
	}
}
