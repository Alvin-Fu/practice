package write

import (
	"bufio"
	"log"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type MyWrite struct {
}

func (*MyWrite) Write(b []byte) (int, error) {
	log.Println(len(b), string(b))
	return len(b), nil
}

func BufIOWrite() {
	w := new(MyWrite)
	w.Write([]byte("1"))
	w.Write([]byte("2"))
	w.Write([]byte("3"))
	w.Write([]byte("4"))
	log.Println("-----")
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte("1"))
	bw.Write([]byte("2"))
	bw.Write([]byte("34556fffffffffffffffffffffffffffffffff"))
	bw.Write([]byte("4"))
	// 将当前缓存中的数据落地
	err := bw.Flush()
	if err != nil {
		log.Fatalf("flush err: %v", err)
		return
	}

}

func Reset() {

}

func ReadFrom() {
	read := strings.NewReader("hello world!")
	w := new(MyWrite)
	bw := bufio.NewWriterSize(w, 3)
	bw.ReadFrom(read)
	err := bw.Flush()
	if err != nil {
		log.Fatalf("flush err: %v", err)
		return
	}
}
