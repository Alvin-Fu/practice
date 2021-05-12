package option

import (
	"crypto/md5"
	"hash/crc32"
	"io"
	"os"
	"rpcx/log"
)


type Option struct {
	ID         int64
	ConfigFile string   `flag:"configFile"`
	HTTPHost   string	`flag:"HTTPHost" cfg:"HTTPHost"`
	HTTPPort   string	`flag:"HTTPPort" cfg:"HTTPPort"`
	APPName    string	`cfg:"APPName"`
}

func NewOption() *Option {
	hostName, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	h := md5.New()
	io.WriteString(h, hostName)
	defaultID := int64(crc32.ChecksumIEEE(h.Sum(nil)) % 1024)
	return &Option{
		ID:         defaultID,
		ConfigFile: "conf/app.conf",
		HTTPHost:   "",
		HTTPPort:   "8080",
		APPName:    "",
	}
}
