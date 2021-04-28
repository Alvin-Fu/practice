package log

import (
	"log"
	"os"
	"io"
)

var (
	logRegister map[string]*log.Logger = make(map[string]*log.Logger)
)

type LogLevel int

const (
	DEBUG = LogLevel(1)
	INFO  = LogLevel(2)
	WARN  = LogLevel(3)
	ERROR = LogLevel(4)
	FATAL = LogLevel(5)
)

type Logger struct {
	FileName string
	Dir string

}



func (l *Logger)LogRegister() {
	file, err := os.OpenFile(l.Dir + l.FileName, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("open file err: %v", err)
	}
	writer := io.MultiWriter(os.Stdout, file)
	flag := log.Ldate | log.Lmicroseconds | log.Lshortfile
	logRegister["debug"] = log.New(writer, "debug", flag)
	logRegister["info"] = log.New(writer, "info", flag)
	logRegister["fatal"] = log.New(writer, "fatal", flag)
}

func Logf(level LogLevel, f string, args ...interface{}) {
	switch level {
	case DEBUG:
		logRegister["debug"].Printf(f,args...)
	}
}
