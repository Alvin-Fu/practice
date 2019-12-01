package abstractfactory

import "fmt"

//四个角色
// 抽象工厂
type AbstractFactory interface {
	Phone5G() AbstractProduct
	Phone4G() AbstractProduct
}

//抽象产品
type AbstractProduct interface {
	Call(tel string)
	Camera()
}

// 具体产品
type Mi5GPhone struct {

}
func (m5 Mi5GPhone) Call(tel string){
	fmt.Println("Welcome to make phone calls with xiaomi 5g! Phone code: ", tel)
}

func (m5 Mi5GPhone) Camera(){
	fmt.Println("Welcome to take photos with xiaomi 5g mobile phone!")
}

type Mi4GPhone struct{

}
func (m4 Mi4GPhone) Call(tel string){
	fmt.Println("Welcome to make phone calls with xiaomi 4g! Phone code: ", tel)
}

func (m4 Mi4GPhone) Camera(){
	fmt.Println("Welcome to take photos with xiaomi 4g mobile phone!")
}

type Honor5GPhone struct {

}
func (h5 Honor5GPhone) Call(tel string){
	fmt.Println("Welcome to make phone calls with honor 5g! Phone code: ", tel)
}
func (h5 Honor5GPhone) Camera(){
	fmt.Println("Welcome to take photos with honor 5g mobile phone!")
}

type Honor4GPhone struct {

}
func (h4 Honor4GPhone) Call(tel string){
	fmt.Println("Welcome to make phone calls with honor 4g! Phone code: ", tel)
}
func (h4 Honor4GPhone) Camera(){
	fmt.Println("Welcome to take photos with honor 4g mobile phone!")
}

// 具体工厂
type MiFactory struct {

}
func NewMiFactory()MiFactory{
	return MiFactory{}
}

func (m MiFactory) CreateMi5GPhone()AbstractProduct{
	return Mi5GPhone{}
}
func (m MiFactory) CreateMi4GPhone()AbstractProduct{
	return Mi4GPhone{}
}

type HonorFactory struct{}

func NewHonorFactory()HonorFactory{
	return HonorFactory{}
}
func (h HonorFactory) CreateHonor5GPhone()AbstractProduct{
	return Honor5GPhone{}
}
func (h HonorFactory) CreateHonor4GPhone()AbstractProduct{
	return Honor4GPhone{}
}







