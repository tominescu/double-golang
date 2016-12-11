package simplelog

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

var log_level int = DEBUG
var level_str = []string{"DEBUG", "INFO", "WARN", "ERROR"}
var logger = log.New(os.Stdout, "", 0)

func output(level int, format string, args ...interface{}) {
	if level < log_level {
		return
	}
	buf := fmt.Sprintf(format, args...)
	now := time.Now().Format("2006-01-02 15:04:05")
	logger.Printf("%s\t%s\t%s\n", now, level_str[level], buf)
}

func SetLevel(level int) {
	log_level = level
}

func GetLevel() int {
	return log_level
}

func Debug(format string, args ...interface{}) {
	output(DEBUG, format, args...)
}

func Info(format string, args ...interface{}) {
	output(INFO, format, args...)
}

func Warn(format string, args ...interface{}) {
	output(WARN, format, args...)
}

func Error(format string, args ...interface{}) {
	output(ERROR, format, args...)
}
