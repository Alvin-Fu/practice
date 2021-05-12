package singleton

import "testing"

func TestGetSingleton(t *testing.T) {
	s1 := GetSingleton()
	s2 := GetSingleton()
	if s1 != s2{
		t.Fatal("not equal")
	}
}
