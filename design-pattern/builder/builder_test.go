package builder

import (
	"fmt"
	"testing"
)

func TestCar1(t *testing.T){
	c := NewCar1()
	d := NewDirector(c)
	d.Construct()
	fmt.Println(c.car)
}

func TestCar2(t *testing.T){
	c := NewCar2()
	d := NewDirector(c)
	d.Construct()
	fmt.Println(c.car)
}