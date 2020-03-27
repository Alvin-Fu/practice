package log

import "log"

func Debugf(format string, v ...interface{}){
	log.Printf(format, v)
}

func Fatalf(format string, v ...interface{}){
	log.Fatalf(format, v)
}
