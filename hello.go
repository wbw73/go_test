package main

import (
	"fmt"
	"github.com/wbw73/go_test/readimage"
	"github.com/wbw73/go_test/show"
	"reflect"
)

//常量 名称  类型
const NAME string = "name"

//全局变量
var name int = 60

//一般类型声明
type lerar int

//结构声明
type Imooc struct {
}

//数据类型别名
type muke int32

//声明接口
//type Ilearn interface {
//}

//函数声明
func learmImooc() {
	//fmt.Printf("aaa")
	//imooc.Printf("BBBBBB")

}

var(
	d int
	f int
	g string
)

const(
	z=iota
	_
	x=iota
)

const(
	y,l=iota,iota*4
	u,m
)

type Data struct {

}

func (self Data) String() string {
	return "data"
}
//main()函数
func main() {
	//fmt.Printf("%v",Data{})

	//io
	//basicio.SampleReadFromString()
	//basicio.SamlpeReadStdin()
	//basicio.SamlpeReadFile()

	//缓冲io
	//bufiosample.IoSample()

	//统计文件行数
	//codecount.CodeCount()

	//读取bmp位图
	readimage.ImgRead()














	//learmImooc()
	//show.Sho()
	/*goto One
	fmt.Print(122)*/
	/*var i muke = 9
	fmt.Println(i)
	//查询占内存几个字节
	fmt.Println(unsafe.Sizeof(i))
	//变量数据类型
    fmt.Println(reflect.TypeOf(i))*/
	//fmt.Printf("hello world")
	//time.Sleep(3 * time.Second)
	/*a:=3

	var b=4
	var c int=9*/
	//var zx,zc,zv=1,"string",3
	//var zs,zd,zf int=3,4,5
	//zs,zd,zf :=3,4,5 全局不能使用
/*	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(x)
	fmt.Println(m)*/
	/*One:
		fmt.Print(1)*/
	//for i, i2 := range collection {
	//
	//}
	//make
	//makeSlice()
	//makeMap()
	//makeChan()


	//makeSliceCopy()

	//makeMapDelete()

	//new
	//NewMap()  //*map[int]string

	//panicRecover()
	//rePanic()

	//d:=new(struct_demo.Dog)
	//d.Age=3
	//d.Color="blue"
	//
	//d.Run()
	//d.Eat()
	//fmt.Print("\n")
	//
	//
	//struct_demo.TestForStruct()

	/*var b show.Behavior
	b=new(struct_demo.Dog)
	b.Eat()*/
	//多态使用接口
	//dog:=new(struct_demo.Dog)
	//action(dog)

	//并发  加入关键字 go

	//fmt.Println(runtime.NumCPU())  //获取当前系统核心数
	//核心数             或者输入数字
	/*runtime.GOMAXPROCS(runtime.NumCPU())
	go gototine.Loop()
	go gototine.Loop()

	time.Sleep(5*time.Second)*/


	//协程通信  channel  select
	/*go gototine.Send()
	go gototine.Receive()
	time.Sleep(5*time.Second)*/


	//协程同步  系统工具sync.waitgroup
	//Add(dela int)添加协程记录
	//Done()移除协程记录
	//Wait()同步等待所有记录的协程全部结束
	/*gototine.Read()
	go gototine.Write()
	gototine.WG.Wait()
	fmt.Println("等待")*/

	//指针
	//point_demo.Pointtest()
	//point_demo.TestPointArr()

	//json
	//json_test.SerializStruct()
	//json_test.SerializMap()
	//json_test.DeSerializStruct()
	//json_test.DeSerializMap()

	//go mod 包管理工具
	//go mod init 初始化mod 生成 go.mod文件  go.sum文件（外部依赖目录，没有依赖不生成）
	//go mod download 下载依赖，下载go.mod文件中required
	//go mod graph 输出使用的所有依赖
	//go mod tidy  删除不使用的依赖，加载包含的依赖（go问价使用了第三方依赖，go.mod文件没有包含）
	//go mod verify 验证包是否正确，源码是否修改
	//go mod why -m "包"  被谁使用
	//go mod edit [-fmt|-module|-require=path@version|...]  go help mod edit  编辑go.mod
	//go mod vendor 当前项目生成verdor目录  保存第三方包，不在下载带pkg目录

	//go get "包" 下载包
	//go build 编译并加载包
	//go list -m all 列所所有依赖

	/*str:=fmt.Sprintf("float %f",3.1415926)
	fmt.Printf(str)*/

	//fmt.Fprintln(os.Stdout,"A")
}

func action(b show.Behavior)  string{
	b.Eat()
	b.Run()
	return ""
}

func  NewMap(){
	newMap:=new(map[int]string)
	fmt.Print(reflect.TypeOf(newMap))
}

//创建切片  slice->append  slice->copy
func makeSlice(){
	mSlice:=make([]string,3)
	mSlice[0]="dog"
	mSlice[1]="cat"
	mSlice[2]="tiger"
	mSlice=append(mSlice,"ti") //追加数据
	fmt.Println(mSlice)
	fmt.Println("真实容量：len=",len(mSlice))
	fmt.Println("系统扩容：cap=",cap(mSlice))
}

//copy数据 数据替换 只替换已有数据 比如：目标数据有2个 ，src3个 只输出前两个
//                                    目标数据3个  src2个  输出src2个，目标数据最后一个
func makeSliceCopy(){
	mSlice:=make([]string,3)
	mSlice[0]="src-dog"
	mSlice[1]="src-cat"
	mSlice[2]="src-1cat"

	mSliceDst:=make([]string,2)
	mSliceDst[0]="dst-dog"
	mSliceDst[1]="dst-cat"
	//mSliceDst[2]="dst-1cat"

	copy(mSliceDst,mSlice)

	fmt.Println(mSliceDst)
}

//创建map  map->delete
func makeMap(){
	mMap:=make(map[int]string)
	mMap[10]="dog"
	mMap[100]="cat"
	fmt.Print(mMap)
}

func makeMapDelete(){
	mMapDe:=make(map[int]string)
	mMapDe[10]="dog"
	mMapDe[100]="cat"
	delete(mMapDe,10)
	fmt.Print(mMapDe)
}

//创建chan 管道
func makeChan(){
	mChan:=make(chan int,3)
	mChan<-1
	fmt.Print(mChan)
	close(mChan)
}
//一般实际写法 打开之后，直接写defer 防止后来忘记关闭
func closeChan()  {
	mChan:=make(chan string,3)
	defer close(mChan)
	mChan<-"hhh"
}

//panic & recover 处理异常
//panic 抛出异常
//recover 捕获异常
//第一种写法
func panicRecover(){
	defer func() {
		message:=recover()
		fmt.Print("panic message:",message)
	}()
	panic("first:sssss")
}
//第二种
func recoverPanic(){
	message:=recover()
	fmt.Print("panic message:",message)
}

func rePanic(){
	defer recoverPanic()
	panic("second:tttttsssss")
}


//len -> string array slice map chan
//cap -> slice array chan
//close -> chan