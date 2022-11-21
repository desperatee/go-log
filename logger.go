package logger

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fatih/color"
)

type LogType string
const (
	LoggerError LogType = "err"
	LoggerInfo  LogType = "info"
	LoggerWarn  LogType = "warn"
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
	case LoggerError:
		if runtime.GOOS == "windows" {
			fmt.Fprintf(color.Output, GenerateDateString()+color.RedString(" ERROR ")+message+"\n")
		} else {
			fmt.Println(GenerateDateString() + color.RedString(" ERROR ") + message + "\n")
		}
	case LoggerInfo:
		if runtime.GOOS == "windows" {
			fmt.Fprintf(color.Output, GenerateDateString()+color.CyanString(" INFO ")+message+"\n")
		} else {
			fmt.Println(GenerateDateString() + color.CyanString(" INFO ") + message + "\n")
		}
	case LoggerWarn:
		if runtime.GOOS == "windows" {
			fmt.Fprintf(color.Output, GenerateDateString()+color.YellowString(" WARN ")+message+"\n")
		} else {
			fmt.Println(GenerateDateString() + color.YellowString(" WARN ") + message + "\n")
		}
	}
}
