package conf

import (
	"crypto/md5"
	"hash/crc32"
	"io"
	"log"
	"os"
)

type Options struct {
	ID int64 `flag:"node-id" cfg:"nodeID"`

	Verbose bool   `flag:"verbose"`
	PIDFile string `flag:"pid-file" cfg:"pidFile"` // path for pid file, default is "./matchsvr_<ID>.pid"

	// http server address, if set, will start a http server use to admin service
	HTTPAddr string `flag:"http-addr" cfg:"httpAddr"`
}

// NewOptions create a default options.
func NewOptions() *Options {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	h := md5.New()
	io.WriteString(h, hostname)
	defaultID := int64(crc32.ChecksumIEEE(h.Sum(nil)) % 1024) // 0-127
	//defaultConfig := fmt.Sprintf("/data/gameserver/gamehall/%s/%s@%d.ini", APPName, APPName, defaultID)

	return &Options{
		ID:      defaultID,
		Verbose: false,
		PIDFile: "",
		// rpc agent conf

		// http server
		HTTPAddr: "0.0.0.0:8081", // default do not start a http server
	}
}
