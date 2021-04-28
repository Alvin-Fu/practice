package options

import (
	"crypto/md5"
	"hash/crc32"
	"io"
	"log"
	"os"
)

type Option struct {
	ID         int64
	ConfigFile string
	HTTPHost   string
	HTTPPort   string
	APPName    string
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
		ID: defaultID,
		//ConfigFile: "../conf/app.conf",
		HTTPHost: "",
		HTTPPort: "8080",
		APPName:  "",
	}
}
