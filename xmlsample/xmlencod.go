package xmlsample

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string `xml:"name,attr"`  //转换成属性
	Age int
}

func XmlSample(){
	p:=Person{Name: "里斯",Age: 12}
	//data,err:=xml.Marshal(p)
	data,err:=xml.MarshalIndent(p,"#"," ")  //prefix某某前缀  Indent 某某缩进
	if err!=nil{
		fmt.Println("xml转换失败:",err)
		return
	}
	//fmt.Println(data)
	fmt.Println(string(data))

	p2:=new(Person)
	err1:=xml.Unmarshal(data,p2)
	if err!=nil{
		fmt.Println("xml转换失败:",err1)
		return
	}
	//fmt.Println(data)
	fmt.Println(p2)
}