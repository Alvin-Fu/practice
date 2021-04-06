package probability

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tmp := make(map[int]int, 10)
	for i := 0; i < 100000000; i++ {
		tmp[rand10()]++
	}
	fmt.Printf("%v\n", tmp)
}
