package simplefactory

import "fmt"

type MobilePhone interface {
	Call(tel int)error
	Camera()
}

func NewPhone(t int)MobilePhone{
	switch t {
	case 1:
		return new(Mi)
	case 2:
		return new(Honor)
	default:
		return nil
	}
}

type Mi struct {

}
func (m *Mi) Call(tel int)error{
	fmt.Println("Call: ", tel)
	return nil
}

func (m *Mi) Camera(){
	fmt.Println("mi camera")
}

type Honor struct{}
 func (h *Honor) Call(tel int) error{
 	fmt.Println("Call: ", tel)
 	return nil
 }

func (h *Honor) Camera(){
	fmt.Println("honor camera")
}

