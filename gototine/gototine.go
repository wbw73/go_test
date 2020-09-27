package gototine

import (
	"fmt"
	"sync"
	"time"
)

var chanNum chan int  =make(chan int,10)
var timeOut chan bool= make(chan bool)

var WG sync.WaitGroup

func Loop()  {
	for i:=1;i<11;i++ {
		time.Sleep(3*time.Microsecond)
		fmt.Printf("%d,", i)
	}
}

//发送数据到chan
func Send()  {
	time.Sleep(1*time.Second)
	chanNum<-1
	time.Sleep(1*time.Second)
	chanNum<-2
	time.Sleep(1*time.Second)
	chanNum<-3
	time.Sleep(1*time.Second)
	timeOut<-true
}

//接受chan数据
func Receive()  {

	//num:=chanNum  输出栈地址

	/*num:=<-chanNum
	fmt.Println(num)
	num=<-chanNum
	fmt.Println(num)
	num=<-chanNum
	fmt.Println(num)*/
	//for true一直接收
	for  {
		select {
			case num:=<-chanNum:
			fmt.Println(num)
			case <-timeOut:
			fmt.Println("timeout...")
		}
	}
}

func Read()  {
	for i:=0;i<3;i++ {
		WG.Add(i)
	}
}

func Write()  {
	for i:=0;i<3;i++ {
		time.Sleep(2*time.Second)
		fmt.Println(i)
		WG.Done()
	}
}