package facade

import (
	"fmt"
	"testing"
)

func TestMyPhone(t *testing.T){
	p := NewMyPhone()
	p.Boot()
	fmt.Println("-----")
	p.Shutdown()
}
