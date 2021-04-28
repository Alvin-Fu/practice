package factory

import "testing"

func TestMi(t *testing.T)  {
	pm := new(ProductMi)
	mi := pm.ProductOnePhone()
	mi.Call("123456789")
	mi.Camera()
}

func TestHonor(t *testing.T) {
	ph := new(ProductHonor)
	honor := ph.ProductOnePhone()
	honor.Call("987654321")
	honor.Camera()
}
