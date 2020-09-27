package point_demo

import "fmt"

func Pointtest()  {

	//不支持指针运算
	
	var num int =1
	var numPoint *int  //指针类型
	numPoint=&num

	var num2 *int
	fmt.Printf("num的地址: %x \n",&num) // &取地址符
	fmt.Printf("numPonit的地址: %x \n",numPoint)
	fmt.Printf("numPoint的值: %d \n",*numPoint) //*取值
	fmt.Printf("num2: ",num2)  //nil  代表空

}

func TestPointArr()  {
	a,b:=1,2
	pointArr:=[...]*int{&a,&b}
	fmt.Println("指针数组 pointArr",pointArr)

	arr:=[...]int{3,4,5}
	arrPoint:=&arr
	fmt.Println("数组指针 arrPoint",arrPoint)

}
