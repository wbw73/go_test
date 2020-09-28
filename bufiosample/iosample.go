package bufiosample

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IoSample()  {
	strReader:=strings.NewReader("hello world")  //读取字符串，返回地址*Reader  处理效率更高

	bufReader:=bufio.NewReader(strReader) //读取地址的值，传给缓存

	p,_:=bufReader.Peek(4)  //从缓存读几个字符
	fmt.Println(p,bufReader.Buffered())  //bufReader.Buffered()  缓存多少个字符
	fmt.Println(string(p))

	data,_:=bufReader.ReadString(' ') //字符串类型读取，读取到空格为止,包含空格
	fmt.Println(data,bufReader.Buffered())

	w:=bufio.NewWriter(os.Stdout)  //输出缓存
	fmt.Fprint(w,"aaaa")   //F 表示写入文件  unix一切皆文件
	fmt.Fprint(w,"bbbb")  //写入bufio
	w.Flush()   //提交给os.Stdout
}