package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	DEBUG = "DEBUG"
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

var logFile *os.File

func InitLogger(logPath string) {
	var err error
	// TODO: an absolute monstrosity fix this later
	path := filepath.Join(append([]string{"/"}, strings.Split(logPath, "/")...)...)
	fmt.Printf(filepath.Dir(path))
	if (filepath.Dir(path) != ".") {
		err = os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			fmt.Printf("Critical: Could not create directory: %v\n", err)
			os.Exit(1)
		}
	}
	logFile, err = os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Critical: Could not open log file: %v\n", err)
		os.Exit(1)
	}
}

func Log(level, message string) {
	if logFile == nil {
		InitLogger("app.log")
	}
	_, file, line, ok := runtime.Caller(1)
	location := "unknown:0"
	if ok {
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				file = file[i+1:]
				break
			}
		}
		location = fmt.Sprintf("%s:%d", file, line)
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(logFile, "[%s] [%s] [%s] %s\n", level, location, timestamp, message)
}
