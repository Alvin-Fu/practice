package abstractfactory

import "testing"

func TestMi(t *testing.T){
	m := NewMiFactory()
	m5 := m.CreateMi5GPhone()
	m5.Call("400")
	m5.Camera()
}

func TestHonor(t *testing.T){
	h := NewHonorFactory()
	h5 := h.CreateHonor5GPhone()
	h5.Call("100")
	h5.Camera()
}