package strconvsample

import (
	"fmt"
	"strconv"
)

func StrConvSample()  {
	//数值转换
	fmt.Println(strconv.Itoa(9999))  //返回字符串类型
	fmt.Println(strconv.Atoi("78665")) //字符串只能是纯数字  字符串转整型

	//解析
	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseFloat("3.14",32))

	//格式化
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(12,2))  //2是转换到2进制
	fmt.Println(strconv.FormatInt(12,8))  //2是转换到8进制
}
