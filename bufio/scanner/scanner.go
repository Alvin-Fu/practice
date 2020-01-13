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
		log.Fatalf("failed to open file: %s, err: %v", name, err)
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("file stat err: %v", err)
		return
	}

	bufio.NewScanner()

}
