package factory

import "fmt"

// 手机的公共接口
type Phone interface {
	Call(tel string)
	Camera()
}

type Mi struct{}
func (m *Mi) Call(tel string)   {
	fmt.Printf("Please call Mi! ")
	fmt.Printf("For further details call this number: %s\n", tel)
}
func (m *Mi) Camera() {}
type Honor struct {}
func (h *Honor) Call(tel string){
	fmt.Printf("Please call Honor! ")
	fmt.Printf("For further details call this number: %s\n", tel)
}
func (h *Honor) Camera(){}

type ProductPhone interface {
	ProductOnePhone() Phone
}

type ProductMi struct {}

func (pm *ProductMi)ProductOnePhone() Phone  {
	return new(Mi)
}

// 表示一个honor的工厂
type ProductHonor struct {}

func (ph *ProductHonor) ProductOnePhone() Phone{
	return new(Honor)
}