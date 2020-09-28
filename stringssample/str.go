package stringssample

import (
	"fmt"
	"strings"
)

func StrSample()  {
	s:="hello world"

	//是否包含
	fmt.Println(strings.Contains(s,"e"),strings.Contains(s,"?"))

	//索引(下标)  首次出现
	fmt.Println(strings.Index(s,"o"))

	ss:="1@!@4@54"

	p:=strings.Split(ss,"@") //转换数组
	fmt.Println(p)

	fmt.Println(strings.Join(p,"#")) //拼接字符串

	//是否包含某个前缀
	fmt.Println(strings.HasPrefix(s,"h"))

	//是否包含某个后缀
	fmt.Println(strings.HasSuffix(s,"h"))
}