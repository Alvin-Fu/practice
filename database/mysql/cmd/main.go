package main

import (
	"log"
	"os"
	"practice/database/mysql"
	"practice/lib/mysql"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	conf := &mysql.MysqlConf{
		Host:     "192.168.200.144",
		Port:     "3388",
		User:     "root",
		Password: "public",
		BaseName: "fuli",
		IdleTime: 180,
	}
	svr, err := mysql.NewMysqlSever(conf)
	if err != nil {
		os.Exit(1)
	}
	svr.Breath(conf)
	db := svr.GetMysqlDb()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	operate.GroupWithRollup(db)
}
