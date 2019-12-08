package adapter

import (
	"fmt"
	"testing"
)

func TestPlayGameOSAdapter_PlayGame(t *testing.T) {
	p := NewOSAdapter()
	p.PlayGame(1)
	fmt.Println("-----")
	p.PlayGame(2)
	fmt.Println("-----")
	p.PlayGame(0)
}
