package operate

import (
	"database/sql"
	"fmt"
	"log"
)

func GroupBy(db *sql.DB) {
	row, err := db.Query(fmt.Sprintf("SELECT `class_id`, COUNT(*) FROM `fuli`.`student` GROUP BY class_id"))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	c, _ := row.Columns()
	for row.Next() {
		rue := make([]int, len(c))
		row.Scan(&rue[0], &rue[1])
		fmt.Println(rue)
	}
}

// 注意order by需要放在后面，不然会报错
func GroupOrder(db *sql.DB) {
	row, err := db.Query("SELECT score, COUNT(*) FROM student  GROUP BY score order by score desc")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	c, _ := row.Columns()
	for row.Next() {
		rue := make([]int, len(c))
		row.Scan(&rue[0], &rue[1])
		fmt.Println(rue)
	}
}

func GroupWithRollup(db *sql.DB) {
	row, err := db.Query("SELECT score, count(*) FROM student  GROUP BY score WITH ROLLUP")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	c, _ := row.Columns()
	for row.Next() {
		rue := make([]interface{}, len(c))
		row.Scan(&rue[0], &rue[1])
		fmt.Println(rue)
	}
}
