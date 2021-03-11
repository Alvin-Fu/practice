package operate

import (
	"database/sql"
	"fmt"
	"log"
)

type Student struct {
	Id      int64 `json:"id" db:"id"`
	Uid     int64 `json:"uid" db:"uid"`
	ClassId int64 `json:"class_id" db:"class_id"`
	Score   int64 `json:"score" db:"score"`
}

func Insert(db *sql.DB) {
	for i := 1000; i < 1500; i++ {
		db.Exec(fmt.Sprintf("INSERT INTO `fuli`.`student`(`uid`, `class_id`, `score`) "+
			"VALUES (%d, %d, %d)", i, i%5, i%100))
	}
}

func OrderBy(db *sql.DB) {
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
