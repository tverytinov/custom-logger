package logger

import (
	"log"
	"os"

	"github.com/fatih/color"
)

const (
	DebugLevelInt int = iota
	InfoLevelInt
	WarnLevelInt
	ErrorLevelInt

	DebugLevelStr string = "DEBUG"
	InfoLevelStr  string = "INFO"
	WarnLevelStr  string = "WARN"
	ErrorLevelStr string = "ERROR"
	msgErrorType  string = "error"
	msgNoticeType string = "notice"
)

type Logger struct {
	log      *log.Logger
	logLevel int
}

func NewLogger(logLevel string) Logger {
	decor := log.Lshortfile | log.Ltime | log.Ldate | log.Lmsgprefix

	logger := log.New(os.Stdout, " ", decor)

	return Logger{
		log:      logger,
		logLevel: defineLogLevel(logLevel),
	}
}

// handlerName - name of handler where error has occurred.
// msgType - often it is error message type but also it can be something noticeable.
// msg - the main entity of this logging.
func (l *Logger) Debug(handlerName, msgType, msg string) {
	if l.logLevel <= DebugLevelInt {
		l.log.Println(
			debugColor(DebugLevelStr),
			color.MagentaString(handlerName),
			msgTypeColor(msgType),
			color.CyanString(msg),
		)
	}
}

// handlerName - name of handler where error has occurred.
// msgType - often it is error message type but also it can be something noticeable.
// msg - the main entity of this logging.
func (l *Logger) Info(handlerName, msgType, msg string) {
	if l.logLevel <= InfoLevelInt {
		l.log.Println(
			debugColor(InfoLevelStr),
			color.MagentaString(handlerName),
			msgTypeColor(msgType),
			color.CyanString(msg),
		)
	}
}

// handlerName - name of handler where error has occurred.
// msgType - often it is error message type but also it can be something noticeable.
// msg - the main entity of this logging.
func (l *Logger) Warn(handlerName, msgType, msg string) {
	if l.logLevel <= WarnLevelInt {
		l.log.Println(
			debugColor(WarnLevelStr),
			color.MagentaString(handlerName),
			msgTypeColor(msgType),
			color.CyanString(msg),
		)
	}
}

// handlerName - name of handler where error has occurred.
// msgType - often it is error message type but also it can be something noticeable.
// msg - the main entity of this logging.
func (l *Logger) Error(handlerName, msgType, msg string) {
	if l.logLevel <= ErrorLevelInt {
		l.log.Println(
			debugColor(ErrorLevelStr),
			color.MagentaString(handlerName),
			msgTypeColor(msgType),
			color.CyanString(msg),
		)
	}
}

func msgTypeColor(msgType string) string {
	switch msgType {
	case msgErrorType:
		return color.RedString(msgType)
	case msgNoticeType:
		return color.YellowString(msgType)
	default:
		return color.HiBlackString(msgType)
	}
}

func debugColor(logLevel string) string {
	switch logLevel {
	case DebugLevelStr:
		return color.GreenString(logLevel)
	case InfoLevelStr:
		return color.YellowString(logLevel)
	case WarnLevelStr:
		return color.BlueString(logLevel)
	case ErrorLevelStr:
		return color.RedString(logLevel)
	default:
		return color.HiBlackString(logLevel)
	}
}

func defineLogLevel(logLevel string) int {
	switch logLevel {
	case DebugLevelStr:
		return 0
	case InfoLevelStr:
		return 1
	case WarnLevelStr:
		return 2
	case ErrorLevelStr:
		return 3
	default:
		return 1
	}
}
