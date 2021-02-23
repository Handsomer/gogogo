package lib

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Info: info 类日志的变量, 与 Warning, Error, Slow, Debug, Statistics 类似.
var (
	Info       *log.Logger
	Warning    *log.Logger
	Error      *log.Logger
	Slow       *log.Logger
	Debug      *log.Logger
	Statistics *log.Logger
	errFile    *os.File
	allFile    *os.File
	slowFile   *os.File
	statFile   *os.File
)
var day = ""

func SetLoggerInTestMode() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Slow = log.New(os.Stderr, "SLOW: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Statistics = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// SetLogger 设置日志 -- 输出到日志文件.
func SetLogger() {
	level := "debug"
	logPath := "/home/sunwei03/go/src/gogogo/logs"
	currentDay := time.Now().Format("20060102")
	if currentDay == day {
		return
	} else {
		day = currentDay
	}

	if nil != errFile {
		_ = errFile.Close()
	}
	if nil != allFile {
		_ = allFile.Close()
	}
	if nil != slowFile {
		_ = slowFile.Close()
	}
	if nil != statFile {
		_ = statFile.Close()
	}

	var err error
	errFile, err := os.OpenFile(logPath+"/error.log."+day, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	allFile, err = os.OpenFile(logPath+"/all.log."+day, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open all log file:", err)
	}

	slowFile, err = os.OpenFile(logPath+"/slow.log."+day, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open slow log file:", err)
	}

	statFile, err = os.OpenFile(logPath+"/statistic.log."+day, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open statistic log file:", err)
	}

	debugWriter := ioutil.Discard
	infoWriter := ioutil.Discard
	warningWriter := ioutil.Discard
	errorWriter := io.MultiWriter(errFile, os.Stderr)
	slowWriter := slowFile
	statWriter := statFile

	switch level {
	case "debug":
		debugWriter = allFile
		infoWriter = allFile
		warningWriter = io.MultiWriter(errFile, os.Stderr)
	case "info":
		infoWriter = allFile
		warningWriter = io.MultiWriter(errFile, os.Stderr)
	case "warning":
		warningWriter = io.MultiWriter(errFile, os.Stderr)
	case "error":
	default:
		log.Fatalln("invalid level:", level)
	}

	Info = log.New(infoWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningWriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Slow = log.New(slowWriter, "SLOW: ", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(debugWriter, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Statistics = log.New(statWriter, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
