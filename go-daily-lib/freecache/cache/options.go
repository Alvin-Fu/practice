package cache

import "time"

type Options struct{
	ID         int64
	ConfigFile string   `flag:"config"`
	TCPAddr   string	`flag:"tcp-addr" cfg:"tcp-addr"`
	APPName    string	`cfg:"app-name"`
	FirstReadTimeOut time.Duration `flag:"first-read-timeout" cfg:"firstReadTimeOut"`
}

func NewOptions()*Options{
	return &Options{}
}

