package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"practice/lib/mysql"
)

type Student struct {
	Id      int64 `json:"id" db:"id"`
	Uid     int64 `json:"uid" db:"uid"`
	ClassId int64 `json:"class_id" db:"class_id"`
	Score   int64 `json:"score" db:"score"`
}

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
	getAll(db)
	//insert(db)
}

func insert(db *sql.DB) {
	for i := 1000; i < 1500; i++ {
		db.Exec(fmt.Sprintf("INSERT INTO `fuli`.`student`(`uid`, `class_id`, `score`) "+
			"VALUES (%d, %d, %d)", i, i%5, i%100))
	}
}

func getAll(db *sql.DB) {
	row, err := db.Query(fmt.Sprintf("SELECT * FROM `fuli`.`student` order by class_id asc, score desc"))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() {
		st := new(Student)
		row.Scan(&st.Id, &st.Uid, &st.ClassId, &st.Score)
		fmt.Println(*st)
	}
}
