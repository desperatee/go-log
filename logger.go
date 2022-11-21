package log

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fatih/color"
)

type LogType string
const (
	Error LogType = "err"
	Info  LogType = "info"
	Warn  LogType = "warn"
)

type Logger struct {
}

func GenerateDateString() string {
	current_time := time.Now()
	stamp := current_time.Local().Format("15:04:05.000")
	return stamp
}

func NewLogger() Logger {
	logger := Logger{}
	return logger
}

func (logger *Logger) Log(logType LogType, message string) {
	switch logType {
	case Error:
		if runtime.GOOS == "windows" {
			fmt.Fprintf(color.Output, GenerateDateString()+color.RedString(" ERROR ")+message+"\n")
		} else {
			fmt.Println(GenerateDateString() + color.RedString(" ERROR ") + message)
		}
	case Info:
		if runtime.GOOS == "windows" {
			fmt.Fprintf(color.Output, GenerateDateString()+color.CyanString(" INFO ")+message+"\n")
		} else {
			fmt.Println(GenerateDateString() + color.CyanString(" INFO ") + message)
		}
	case Warn:
		if runtime.GOOS == "windows" {
			fmt.Fprintf(color.Output, GenerateDateString()+color.YellowString(" WARN ")+message+"\n")
		} else {
			fmt.Println(GenerateDateString() + color.YellowString(" WARN ") + message)
		}
	}
}
