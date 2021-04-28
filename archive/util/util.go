package util

import (
	"log"
	"os"
	"strings"
)

func IsContinue(str string) bool {
	last := len(str)
	if os.IsPathSeparator(str[last-1]) {
		err := os.MkdirAll(str, 654)
		if err != nil {
			log.Println(err)
		}
		return true
	}
	bytes := []byte(str)
	n := strings.LastIndex(str, string(os.PathSeparator))
	if n < 0 {
		n = strings.LastIndex(str, "/")
	}
	str = string(bytes[:n]) + string(os.PathSeparator)
	err := os.MkdirAll(str, 654)
	if err != nil {
		log.Println(err)
	}
	return false
}

func ParentDir(name string) string {
	last := len(name)
	if os.IsPathSeparator(name[last-1]) {
		return name
	}
	bytes := []byte(name)
	n := strings.LastIndex(name, string(os.PathSeparator))
	if n < 0 {
		n = strings.LastIndex(name, "/")
	}
	return string(bytes[:n]) + string(os.PathSeparator)
}

func IsExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsFile(name string) bool {
	return !IsDir(name)

}

func IsDir(name string) bool {
	f, err := os.Stat(name)
	if err != nil {
		return false
	}
	return f.IsDir()
}
