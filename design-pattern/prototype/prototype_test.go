package prototype

import (
	"fmt"
	"testing"
	"time"
)


func TestProto_CloneProto(t *testing.T) {
	p := NewProto()
	cp := p.CloneProto()
	cp.Tmp = 10
	go func(){
		<- p.Ch
		fmt.Println("hello p")
	}()
	go func(){
		<- cp.Ch
		fmt.Println("hello cp")
	}()
	cp.Ch <- struct{}{}
	time.Sleep(2 * time.Second)
}
func TestProto_CloneProto2(t *testing.T) {

}

func BenchmarkNewProto(b *testing.B) {
	for i := 0; i < b.N; i++{
		NewProto()
	}
}

func BenchmarkProto_CloneProto(b *testing.B) {
	p := NewProto()
	for i := 0; i < b.N; i++{
		p.CloneProto()
	}
}