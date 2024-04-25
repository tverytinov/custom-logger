package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	Log      *log.Logger
	LogLevel string
}

func NewLogger(logLevel string) Logger {
	decor := log.Lshortfile | log.Ltime | log.Ldate | log.Lmsgprefix

	logger := log.New(os.Stdout, " ", decor)

	return Logger{
		Log:      logger,
		LogLevel: logLevel,
	}
}

func (l *Logger) Debug(msg, msgtype, description string) {
	l.Log.Println(
		color.GreenString(l.LogLevel),
		color.MagentaString(msg),
		color.RedString(msgtype),
		color.BlueString(msgtype),
	)
}
