package scanner

import (
	"bufio"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}
func Scanner() {
	file, err := os.Open("../readme.md")
	if err != nil {
		log.Fatalf("failed ")
	}
	bufio.NewScanner()

}
