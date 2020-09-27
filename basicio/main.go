package basicio

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(read io.Reader,num int) ([]byte,error) {
 	p:=make([]byte,num)
 	n,err:=read.Read(p)
 	if n>0{
 		return p[:n],nil
	}
	return p,err

}

func SampleReadFromString()  {
	data,_:=ReadFrom(strings.NewReader("aAfrom string"),7)
	fmt.Println(string(data))
	fmt.Println(data)
}

//命令行读取
func SamlpeReadStdin()  {
	fmt.Println("please input from stdin:")

	data,_:=ReadFrom(os.Stdin,11)

	fmt.Println(data)
	fmt.Println(string(data))
}

//文件读取
func SamlpeReadFile()  {
	file,_:=os.Open("hello.go")
	defer file.Close()

	data,_:=ReadFrom(file,16)
	fmt.Println(data)
	fmt.Println(string(data))
}