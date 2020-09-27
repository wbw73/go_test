package struct_demo

import (
	"fmt"
)

type Animal struct {
	Color string
}

type Dog struct {
	Animal //继承
	Id int
	Name string
	Age int
}

func TestForStruct()  {
	//方式1
	/*var dog Dog
	dog.Id=1
	dog.Name="kiki"
	dog.Age=3*/

	//方式2
	//dog:=Dog{Id:2,Name:"Yaya",Age:2}

	//方式3
	dog:=new(Dog)
	dog.Id=3
	dog.Name="Uiui"
	dog.Age=6
	dog.Color="red"

	fmt.Print(dog)
}
//首字母大写，外部才能调用  小写只能内部调用 结构体内部属性同理  string返回值类型
func (d *Dog) Run() string{
	fmt.Println(d.Age,"jjjjj")
	return "run"
}

func (d *Dog)Eat() string{
	fmt.Println(d.Color,"hhhh")
	return "eat"
}