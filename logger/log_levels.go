package logger

import (
	"log"
)

// Log levels
const (
	LEVEL_TRACE = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR // max level
)

const (
	defaultLogLevel = LEVEL_INFO
)

var (
	currentLevel = LEVEL_INFO

	strs = []string{
		"[TRACE] ",
		"[DEBUG] ",
		"[INFO] ",
		"[WARN] ",
		"[ERROR] ",
	}

	Tracef, Debugf, Infof, Warnf, Errorf func(format string, v ...any)
	Trace, Debug, Info, Warn, Error      func(v ...any)

	noOpf = func(format string, v ...any) {}
	noOp  = func(v ...any) {}

	Fatalf = log.Fatalf
	Fatal  = log.Fatal
	Panicf = log.Panicf
	Panic  = log.Panic
)

func init() {
	SetLogLevel(defaultLogLevel)
	EnableLogs()
}

func EnableLogs() {
	Tracef = printfByLevel(LEVEL_TRACE)
	Debugf = printfByLevel(LEVEL_DEBUG)
	Infof = printfByLevel(LEVEL_INFO)
	Warnf = printfByLevel(LEVEL_WARN)
	Errorf = printfByLevel(LEVEL_ERROR)

	Trace = printByLevel(LEVEL_TRACE)
	Debug = printByLevel(LEVEL_DEBUG)
	Info = printByLevel(LEVEL_INFO)
	Warn = printByLevel(LEVEL_WARN)
	Error = printByLevel(LEVEL_ERROR)
}

func DisableLogs() {
	Tracef = noOpf
	Debugf = noOpf
	Infof = noOpf
	Warnf = noOpf
	Errorf = noOpf

	Trace = noOp
	Debug = noOp
	Info = noOp
	Warn = noOp
	Error = noOp
}

func SetLogLevel(level int) {
	if level < 0 || level > LEVEL_ERROR {
		log.Printf("[WARN]: Invalid log level '%d' provided. Defaulting to INFO level.\n", level)
		level = defaultLogLevel
	}
	currentLevel = level
	Infof("Log level set to %s\n", strs[level])
}

func printfByLevel(level int) func(format string, v ...any) {
	return func(format string, v ...any) {
		if currentLevel <= level {
			log.Printf(strs[level]+format, v...)
		}
	}
}

func printByLevel(level int) func(v ...any) {
	return func(v ...any) {
		if currentLevel <= level {
			log.Printf(strs[level]+"%v\n", v...)
		}
	}
}
