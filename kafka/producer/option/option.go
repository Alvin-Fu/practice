package option

import (
    "os"
    "log"
    "crypto/md5"
    "io"
    "hash/crc32"
)

type Options struct {
    ID int64 `flag:"node-id" cfg:"node-id"`
    ConfigFile string `flag:"config"`
    KafkaAddr string  `flag:"kafka-addr" cfg:"kafka-addr"`
    HTTPAddr string `flag:"http-addr" cfg:"http-addr"`
}

func NewOptions()*Options{
    hostname, err := os.Hostname()
    if err != nil {
        log.Fatal(err)
    }
    h := md5.New()
    io.WriteString( h, hostname)
    defaultID := int64(crc32.ChecksumIEEE(h.Sum(nil)) % 1024)
    return &Options{
        ID: defaultID,
        ConfigFile: "conf.ini",
        KafkaAddr: "127.0.0.1:9092",
        HTTPAddr: "0.0.0.0:8080",
    }
}


