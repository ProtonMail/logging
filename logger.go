package logging

import (
	"fmt"
	"os"

	logrus "github.com/Sirupsen/logrus"
)

// Logger is logrus.logger extension
type Logger struct {
	*logrus.Logger
	name string
	path string
	file string
}

var loggers = make(map[string]*Logger)

func init() {
	if len(loggers) <= 0 {
		GetLogger("default")
	}
}

// PrintStatus info printing
func PrintStatus() {
	fmt.Println("Logger Count: ", len(loggers))
	for _, value := range loggers {
		fmt.Println("  Logger Name: ", value.name)
		if value.path == "" {
			fmt.Println("    ", value.name, " Path: ", "(not set)")
		} else {
			fmt.Println("    ", value.name, " Path: ", value.path)
		}

		if value.file == "" {
			fmt.Println("    ", value.name, " File: ", "(not set)")
		} else {
			fmt.Println("    ", value.name, " File: ", value.file)
		}
	}
}

// GetDefaultLogger get a default logger will the std::out
func GetDefaultLogger() *Logger {
	return GetLogger("default")
}

// GetLogger get a logger based on a module, module can't be null
func GetLogger(module string) *Logger {
	var logger = loggers[module]
	if logger != nil {
		return logger
	}
	var newLogger = logrus.New()
	var l = &Logger{newLogger, module, "", ""}
	loggers[module] = l
	l.ConfigUseDefault()
	return l
}

// Config setup the logger with file or default
func (l *Logger) Config(path string, file string, level logrus.Level) {

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666) //file will based on time open
	if err != nil {
		fmt.Println(err)
	}

	l.path = path
	l.file = file
	l.Formatter = new(logrus.JSONFormatter)
	l.Out = f
	l.Level = level
}

// ConfigUseDefault use the default setting to config a logger output will be std::err
func (l *Logger) ConfigUseDefault() {
	l.Formatter = new(logrus.TextFormatter)
	l.Out = os.Stderr
	l.Level = logrus.DebugLevel
}

// GetReport get log report
func (l *Logger) GetReport() string {
	return "" //here need add return current log file example return 1000 lines of log file //TODO:: later
}

// GetAllReport get all exsiting logger's log
func GetAllReport() string {
	return "" //here need add return current log file example return 1000 lines of log file //TODO:: later
}
