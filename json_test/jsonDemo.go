package json_test

import (
	"encoding/json"
	"fmt"
)

type JsonDemo struct {
	JsonNmae string `json:"name"`  //tag
	JsonAge int `json:"age"`
	JsonId int `json:"id"`
}
//结构序列化json  需要先定义一个结构
func SerializStruct()  {
	j:=new(JsonDemo)
	j.JsonAge=12
	j.JsonId=1
	j.JsonNmae="张三"

	b,err := json.Marshal(j) //序列化成json
	if err !=nil{
		fmt.Printf("转换json失败:",err.Error())
		return
	}
	fmt.Println(string(b))
}
//map序列化json
func SerializMap()  {
	j:=make(map[string]interface{})
	j["JsonAge"]=14
	j["JsonId"]=2
	j["jsonName"]="里斯"

	b,err := json.Marshal(j) //序列化成json
	if err !=nil{
		fmt.Printf("转换json失败:",err.Error())
		return
	}
	fmt.Println(string(b))
}

//结构反序列化json
func DeSerializStruct()  {
	str:=`{"age":14,"id":2,"name":"里斯"}`
	j:=new(JsonDemo)

	err := json.Unmarshal([]byte(str),&j) //序列化成json
	if err !=err{
		fmt.Printf("反序列化json失败:",err.Error())
		return
	}
	fmt.Println(j)
}
//map反序列化json
func DeSerializMap(){

	str:=`{"age":14,"id":2,"name":"里斯"}`
	j:=make(map[string]interface{})
	err := json.Unmarshal([]byte(str),&j) //序列化成json
	if err !=err{
		fmt.Printf("转换json失败:",err.Error())
		return
	}
	fmt.Println(j)
}