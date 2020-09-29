package flagsample

import (
	"flag"
	"fmt"
	"os"
)

func OsSample()  {
	fmt.Println(os.Args)

	for i, j:= range os.Args {
		fmt.Println(i,"---->",j)
	}
}

func FlagSample()  {

	//hello.exe -method jjj -value 8  调用命令
	//格式化定义
	methodPointer:=flag.String("method","default","method of sample")
	valuePoint:=flag.Int("value",-1,"value of sample")
	//解析
	flag.Parse()
	//输出
	fmt.Println(*methodPointer)
	fmt.Println(*valuePoint)

}

func Style()  {

	//hello.exe -aa ss -value 88

	var method string
	var value int
	flag.StringVar(&method,"aa","default","method of sample")
	flag.IntVar(&value,"value",-1,"value of sample")
	flag.Parse()

	fmt.Println(method,value)

}