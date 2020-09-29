package vsprojparse

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func GetAttributeValue(attr []xml.Attr,name string) string  {
	for _,a:=range attr{
		if a.Name.Local==name{
			return a.Value
		}
	}
	return ""
}

func ProjparseXml()  {
	data,err:=ioutil.ReadFile("01GZipStream.csproj") //读取的文件内容

	//bytes.NewBuffer 对二进制内容封装后可以读写 传递给xml解析 节点变成了token
	decoder:=xml.NewDecoder(bytes.NewBuffer(data))

	var t xml.Token
	var inItemGroup bool
	for t,err=decoder.Token();err==nil;t,err=decoder.Token() {   //decoder.Token() 获取节点名称
		switch token:=t.(type) {
		case xml.StartElement:
			name:=token.Name.Local
			//fmt.Println(name)
			if inItemGroup{
				if name=="Compile"{
					fmt.Println(name)
					fmt.Println(token.Attr)
					fmt.Println(GetAttributeValue(token.Attr,"Include"))
				}
			}else{
				if name=="ItemGroup"{
					inItemGroup=true
				}
			}
		case xml.EndElement:
			if inItemGroup{
				if token.Name.Local=="ItemGroup"{
					inItemGroup=false
				}
			}
		}
	}
}