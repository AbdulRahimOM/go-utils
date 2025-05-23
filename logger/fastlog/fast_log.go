package fastlog

import (
	"log"
)

// Log levels
const (
	LEVEL_TRACE = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
	LEVEL_PANIC // max level
)

type logLevel int8

const (
	defaultLogLevel = LEVEL_INFO
)

var (
	currentLevel logLevel = LEVEL_INFO

	strs = []string{
		"[TRACE] ",
		"[DEBUG] ",
		"[INFO] ",
		"[WARN] ",
		"[ERROR] ",
		"[FATAL] ",
		"[PANIC] ",
	}
)

func init() {
	SetLogLevel(defaultLogLevel)
}

// SetLogLevel sets the log level for the logger.
// If the level is beyond the defined range, it will be set to the closest valid level.
// LEVEL_TRACE is the lowest level, and LEVEL_ERROR is the highest.
func SetLogLevel(level logLevel) {
	switch {
	case level < 0:
		level = LEVEL_TRACE
	case level > LEVEL_ERROR:
		level = LEVEL_ERROR
	}
	currentLevel = level
	Infof("Log level set to %s\n", strs[level])
}

func Tracef(format string, v ...any) {
	if currentLevel <= LEVEL_TRACE {
		log.Printf(strs[LEVEL_TRACE]+format, v...)
	}
}

func Trace(v ...any) {
	if currentLevel <= LEVEL_TRACE {
		log.Printf(strs[LEVEL_TRACE]+"%v\n", v...)
	}
}

func Debugf(format string, v ...any) {
	if currentLevel <= LEVEL_DEBUG {
		log.Printf(strs[LEVEL_DEBUG]+format, v...)
	}
}

func Debug(v ...any) {
	if currentLevel <= LEVEL_DEBUG {
		log.Printf(strs[LEVEL_DEBUG]+"%v\n", v...)
	}
}

func Infof(format string, v ...any) {
	if currentLevel <= LEVEL_INFO {
		log.Printf(strs[LEVEL_INFO]+format, v...)
	}
}

func Info(v ...any) {
	if currentLevel <= LEVEL_INFO {
		log.Printf(strs[LEVEL_INFO]+"%v\n", v...)
	}
}

func Warnf(format string, v ...any) {
	if currentLevel <= LEVEL_WARN {
		log.Printf(strs[LEVEL_WARN]+format, v...)
	}
}

func Warn(v ...any) {
	if currentLevel <= LEVEL_WARN {
		log.Printf(strs[LEVEL_WARN]+"%v\n", v...)
	}
}

func Errorf(format string, v ...any) {
	if currentLevel <= LEVEL_ERROR {
		log.Printf(strs[LEVEL_ERROR]+format, v...)
	}
}

func Error(v ...any) {
	if currentLevel <= LEVEL_ERROR {
		log.Printf(strs[LEVEL_ERROR]+"%v\n", v...)
	}
}

func Fatalf(format string, v ...any) {
	if currentLevel <= LEVEL_FATAL {
		log.Fatalf(strs[LEVEL_FATAL]+format, v...)
	}
}

func Fatal(v ...any) {
	if currentLevel <= LEVEL_FATAL {
		log.Fatalf(strs[LEVEL_FATAL]+"%v\n", v...)
	}
}

func Panicf(format string, v ...any) {
	if currentLevel <= LEVEL_PANIC {
		log.Panicf(strs[LEVEL_PANIC]+format, v...)
	}
}

func Panic(v ...any) {
	if currentLevel <= LEVEL_PANIC {
		log.Panicf(strs[LEVEL_PANIC]+"%v\n", v...)
	}
}
