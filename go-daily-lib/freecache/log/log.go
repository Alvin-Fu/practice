package log

import (
	"log"
	"fmt"
	"time"
	"strings"
	"runtime"
)

var (
	logRegister     map[string]*log.Logger = make(map[string]*log.Logger)
)

func init(){
	log.SetFlags( log.Lshortfile |log.LstdFlags| log.Llongfile)
}

//初始化logger操作
//func LogRegister(logLevelName string) *log.Logger {
//	flag := strings.HasSuffix(logPath, "/")
//	if flag {
//		defaultFileName = logPath + logLevelName + "-" + prefix + "-" + time.Now().Format("2006-01-02") + ".log"
//	} else {
//		defaultFileName = logPath + "/" + logLevelName + "-" + prefix + "-" + time.Now().Format("2006-01-02") + ".log"
//	}
//
//	if v, ok := logRegister[logLevelName]; ok {
//		return v
//	}
//
//	//设置默认的writer
//	writer := &lumberjack.Logger{
//		Filename:   defaultFileName,
//		MaxSize:    maxSize, // megabytes
//		MaxBackups: DEFAULTMAXBACKUPS,
//		MaxAge:     DEBAULTMAXAGE,   //days
//		Compress:   DEFAULTCOMPRESS, // disabled by default
//	}
//
//	//writer := os.Stdout
//
//	//writer 设置为标准输出
//	//l = log.New(writer, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
//	l = log.New(writer, "", 0) //Ldate显示形如 2009/01/23 的日期 //Lmicroseconds显示时分秒日期//Lshortfile显示文件名和行号: d.go:23
//	logRegister[logLevelName] = l
//	return l
//}

//只初始化一个logger写日志
func createMergeWrite() {
	logger := new(log.Logger)
	//logger := LogRegister("info")
	logRegister["core"] = logger
	logRegister["fatal"] = logger
	logRegister["warn"] = logger
	logRegister["debug"] = logger
	logRegister["info"] = logger
}

func setLogPrefix(level string) string {
	var prefix string
	switch level {
	case "info":
		prefix = "[INFO ]"
	case "debug":
		prefix = "[DEBUG]"
	case "warn":
		prefix = "[WARNING]"
	case "fatal":
		prefix = "[FATAL]"
	case "core":
		prefix = "[CORE ]"
	default:
		prefix = "[UNKNOWN]"
	}
	t := time.Now()
	y, m, d := t.Date()

	prefixWithTime := fmt.Sprintf("%-7s%d-%d-%d %02d:%02d:%02d.%03d %s ", prefix, y, m, d, t.Hour(), t.Minute(), t.Second(), t.Nanosecond()/1000000, logFileAndLine())
	return prefixWithTime
}

func logFileAndLine() string {
	pc, file, line, ok := runtime.Caller(4)
	funcFullName := runtime.FuncForPC(pc).Name()

	////获取文件名路径（剔除了GOPATH部分）
	//if !ok {
	//	file = "???"
	//	line = 0
	//}
	//gopath := os.Getenv("GOPATH")
	//if gopath == "" {
	//	gopath = build.Default.GOPATH
	//}
	//noPathFileName := strings.Replace(file, gopath+"/src/", "", 1)

	//获取文件名称
	var noPathFileName string
	if !ok {
		file = "???"
		line = 0
	}
	fileIndex := strings.LastIndex(file, "/")
	if len(file) > fileIndex {
		noPathFileName = file[fileIndex+1:]
	} else {
		noPathFileName = file[fileIndex:]
	}

	//获取方法名称
	var funcName string
	fIndex := strings.LastIndex(funcFullName, "/")
	if len(funcFullName) > fIndex {
		funcName = funcFullName[fIndex+1:]
	} else {
		funcName = funcFullName[fIndex:]
	}

	return fmt.Sprintf("%s [func:%s] [line:%d] ", noPathFileName, funcName, line)
}

func Debugf(format string, v ...interface{}){
	prefix := setLogPrefix("debug") //获取前缀信息文件名、方法名、行数
	logRegister["debug"].Output(3, prefix+fmt.Sprintf(format, v...))
}

func Fatalf(format string, v ...interface{}){
	prefix := setLogPrefix("fatal") //获取前缀信息文件名、方法名、行数
	logRegister["fatal"].Output(3, prefix+fmt.Sprintf(format, v...))
}
