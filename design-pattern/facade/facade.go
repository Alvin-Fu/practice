package facade

import "fmt"

// 手机是CPU，内存，硬盘等的外观，有了外观后，启动手机和关闭手机都简化了
// 直接new一个手机，对外暴露的方法（开启和关闭）
type Phone interface {
	Boot()
	Shutdown()
}

type CPU struct {

}
func (c *CPU)Open(){
	fmt.Println("open CPU")
}

func (c *CPU) Close(){
	fmt.Println("close CPU")
}

type Ddr struct {

}
func (d *Ddr) Open(){
	fmt.Println("open Ddr")
}

func (d *Ddr) Close(){
	fmt.Println("close Ddr")
}

type Disk struct{

}
func (d *Disk) Open(){
	fmt.Println("open disk")
}

func (d *Disk) Close(){
	fmt.Println("close disk")
}

type MyPhone struct {
	CPU
	Disk
	Ddr
}

func NewMyPhone()*MyPhone{
	return &MyPhone{}
}

func (ob *MyPhone)Boot(){
	ob.CPU.Open()
	ob.Ddr.Open()
	ob.Disk.Open()
}

func (ob *MyPhone) Shutdown(){
	ob.Disk.Close()
	ob.Ddr.Close()
	ob.CPU.Close()
}
