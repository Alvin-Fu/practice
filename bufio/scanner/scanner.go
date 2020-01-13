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
	name := "../readme.md"
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("failed to open file: %s", name)
		return
	}
	defer file.Close()

	bufio.NewScanner()

}
