package codecount

import (
	"bufio"
	"fmt"
	"os"
)

func CodeCount()  {
	if len(os.Args)<2{
		fmt.Println("输入错误")
		return
	}
	file,err:=os.Open(os.Args[1])
	if err!=nil{
		fmt.Println("文件打开失败：",err)
		return
	}
	defer file.Close()
	reader:=bufio.NewReader(file)
	var i int
	for{
		_,isPrefix,err:=reader.ReadLine()
		if err!=nil{
			break
		}
		//fmt.Println(isPrefix)  isPrefix==false 是有效的读取
		if !isPrefix{
			i++
		}

	}
	fmt.Println(i)
}