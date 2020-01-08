package read

import (
	"bufio"
	"log"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func ReadSlice() {
	read := strings.NewReader("hello world")
	log.Printf("size: %v", read.Size())
	reader := bufio.NewReaderSize(read, 2)
	log.Printf("size: %d", reader.Size())
	line, _ := reader.ReadSlice(' ')
	log.Printf("line: %s", string(line))

	line, _ = reader.ReadSlice(' ')
	log.Printf("line: %s", string(line))
}

func ReadString() {
	read := strings.NewReader("hello world")
	log.Printf("size: %v", read.Size())
	reader := bufio.NewReaderSize(read, 2)
	str, _ := reader.ReadString(' ')
	log.Printf("str: %s", str)
	str, _ = reader.ReadString(' ')
	log.Printf("str: %s", str)

}
