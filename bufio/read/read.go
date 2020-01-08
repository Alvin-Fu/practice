package read

import (
	"bufio"
	"log"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func Read() {
	read := strings.NewReader("hello world")
	log.Printf("size: %v", read.Size())
	reader := bufio.NewReaderSize(read, 2)
	log.Printf("size: %d", reader.Size())
	line, _ := reader.ReadSlice(' ')
	log.Printf("line: %s", string(line))

	line, _ = reader.ReadSlice(' ')
	log.Printf("line: %s", string(line))
}
