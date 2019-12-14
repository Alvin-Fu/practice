package proxy

import "fmt"

type Subject interface {
	DoSomething()
}

type RealSubject struct{}

func (*RealSubject) DoSomething(){
	fmt.Println("do something")
}
func newRealSubject()*RealSubject{
	return &RealSubject{}
}
type ProxySubject struct{
	realSubject *RealSubject
}

func NewProxySubject()*ProxySubject{
	return &ProxySubject{
		realSubject: newRealSubject(),
	}
}
func (p *ProxySubject) DoSomething(){
	p.realSubject.DoSomething()
}

