package simplefactory

import "testing"

func TestMi(t *testing.T){
	mi := NewPhone(1)
	err := mi.Call(6666)
	if err != nil {
		t.Fatal(err)
	}
	mi.Camera()
}

func TestHonor(t *testing.T){
	honor := NewPhone(2)
	err := honor.Call(8888)
	if err != nil {
		t.Fatal(err)
	}
	honor.Camera()
}