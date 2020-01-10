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
	line1, _ := reader.ReadSlice(' ')
	log.Printf("line1: %s", string(line1))

	line2, _ := reader.ReadSlice(' ')
	log.Printf("line1: %s", string(line1))
	log.Printf("line2: %s", string(line2))
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
	bytes1, _ := reader.ReadBytes(' ')
	log.Printf("bytes1: %s", string(bytes1))
	bytes2, _ := reader.ReadBytes(' ')
	log.Printf("bytes1: %s", string(bytes1))
	log.Printf("bytes2: %s", string(bytes2))
}

func Reader() {
	read := strings.NewReader("hello world")
	log.Printf("size: %v", read.Size())
	reader := bufio.NewReaderSize(read, 2)
	var bytes = make([]byte, 0)
	n, err := reader.Read(bytes)
	if err != nil {
		log.Fatalf("read err: %v", err)
		return
	}
	log.Printf("read num: %d", n)
	log.Printf("bytes: %s", string(bytes))
}
