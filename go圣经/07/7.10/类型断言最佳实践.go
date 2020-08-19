package main

import (
	"fmt"
)

type Usb interface {
	Connect()
	Disconnect()
}

type Phone struct{ Name string }
type Camera struct{ Name string }

func (p Phone) Connect()     { fmt.Printf("%s连接中...\n", p.Name) }
func (p Phone) Disconnect()  { fmt.Printf("%s断开连接中...\n", p.Name) }
func (c Camera) Connect()    { fmt.Printf("%s连接中...\n", c.Name) }
func (c Camera) Disconnect() { fmt.Printf("%s断开连接中...\n", c.Name) }

// Phone结构体特有的方法Call
func (p Phone) Call() { fmt.Printf("正在使用%s打电话...\n", p.Name) }
// Camera
func (c Camera) Photograph() { fmt.Printf("正在使用%s拍照...\n",c.Name)}
type Computer struct{}

func (c Computer) Working(usb Usb) {
	usb.Connect()
	usb.Disconnect()
	// 如果 usb 是指向 Phone 结构体变量，则还需要调用 Call 方法
	//phone, ok := usb.(Phone) // 类型断言
	//if ok {
	//	phone.Call()
	//}
	switch v := usb.(type){
	case Phone:
		v.Call()
	case Camera:
		v.Photograph()
	default:
		fmt.Println("不支持的设备")
	}
}

func main() {
	// 定义一个 Usb 接口数组，可以存放 Phone 和 Camera 结构体的实例
	var usbArr [2]Usb
	usbArr[0] = Phone{"苹果"}
	usbArr[1] = Camera{"尼康"}
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
	}
	fmt.Println()
	fmt.Println(usbArr)
}
