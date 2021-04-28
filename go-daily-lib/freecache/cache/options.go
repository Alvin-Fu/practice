package cache

import (
	"time"
	"os"

	"log"
	"crypto/md5"
	"io"
	"hash/crc32"
)

type Options struct{
	ID         int64
	ConfigFile string   `flag:"config-file"`
	TCPAddr   string	`flag:"tcp-addr" cfg:"tcp-addr"`
	APPName    string	`cfg:"app-name"`
	FirstReadTimeOut time.Duration `flag:"first-read-timeout" cfg:"firstReadTimeout"`
}

func NewOptions()*Options{
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	h := md5.New()
	io.WriteString(h, hostName)
	id := int64(crc32.ChecksumIEEE(h.Sum(nil)) % 1024)

	return &Options{
		ID: id,
		ConfigFile: "conf/conf.ini",
		TCPAddr: "127.0.0.1:8080",
		APPName: "cache",
		FirstReadTimeOut: 1,
	}
}

