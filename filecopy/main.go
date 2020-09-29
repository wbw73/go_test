package filecopy

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func FileCopy()  {
	var f,v bool
	flag.BoolVar(&f,"f",false,"强制拷贝")
	flag.BoolVar(&v,"v",false,"显示拷贝过程")

	flag.Parse()

	if flag.NArg()<2{
		flag.Usage()
		fmt.Println("输入错误")
		return
	}

	ActionCopyFile(flag.Arg(0),flag.Arg(1),f,v)
}

func ActionCopyFile(src,des string,f,v bool) {
	if !f{
		if FileExits(des){
			fmt.Printf("%s目标文件已存在，确定覆盖吗? y/n",des)
			reader:=bufio.NewReader(os.Stdin)
			data,_,_:=reader.ReadLine()
			if strings.TrimSpace(string(data))!="y"{
				return
			}

		}
	}

	CopyFile(src,des)

	if v{
		fmt.Printf("%s-----%s",src,des)
	}
}

func CopyFile(src,des string) (w int64,err error)   {
	srcFile,err:=os.Open(src)
	if err!=nil{
		fmt.Println("源文件出错：",err.Error())
	}
	defer srcFile.Close()
	desFile,err:=os.Create(des)
	if err!=nil{
		fmt.Println("源文件出错：",err.Error())
	}
	desFile.Close()

	return io.Copy(desFile,srcFile)
}

func FileExits(filename string) bool  {
	_,err:=os.Stat(filename)
	/*if err==nil{
		return true
	}else{
		return false
	}*/
	return err==nil || os.IsExist(err)
}


