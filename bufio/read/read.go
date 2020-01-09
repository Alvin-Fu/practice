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

	lin, _ := reader.ReadSlice(' ')
	log.Printf("line: %s", string(line))
	log.Printf("lin: %s", string(lin))
}

func ReadString() {
	read := strings.NewReader("hello world")
	log.Printf("size: %v", read.Size())
	reader := bufio.NewReaderSize(read, 2)
	str1, _ := reader.ReadString(' ')
	log.Printf("str1: %s", str1)
	str2, _ := reader.ReadString(' ')
	log.Printf("str1: %s", str1)
	log.Printf("str2: %s", str2)
}

func ReadBytes() {
	read := strings.NewReader("hello world")
	log.Printf("size: %v", read.Size())
	reader := bufio.NewReaderSize(read, 2)
	bytes, _ := reader.ReadBytes(' ')
	log.Printf("bytes: %s", string(bytes))
	bs, _ := reader.ReadBytes(' ')
	log.Printf("bytes: %s", string(bytes))
	log.Printf("bs: %s", string(bs))
}
