package readimage

import (
	"encoding/binary"
	"fmt"
	"os"
)

type BitMapReadInfo struct {  //C++提供读取顺序
	Size uint32
	Width int32
	Height int32
	Place uint16
	BitCount uint16
	Companiession uint32
	SizeImage uint32
	XperlsPerMeter int32
	YperlsPerMeter int32
	Clsrused uint32
	ClrImportant uint32

}

func ImgRead()  {
	file,_:=os.Open("1.bmp")
	defer file.Close()
	var headA,headB byte   //头按两个字节读取 结果 规定是BM  字节是66  77
	binary.Read(file,binary.LittleEndian,&headA)
	binary.Read(file,binary.LittleEndian,&headB)

	var size uint32
	binary.Read(file,binary.LittleEndian,&size)

	var reservedA,reservedB uint16
	binary.Read(file,binary.LittleEndian,&reservedA)
	binary.Read(file,binary.LittleEndian,&reservedB)

	var offbite uint32
	binary.Read(file,binary.LittleEndian,&offbite)

	fmt.Println(headA)
	fmt.Println(headB)
	fmt.Println(size)
	fmt.Println(reservedA)
	fmt.Println(reservedB)
	fmt.Println(offbite)

	inf:=new(BitMapReadInfo)
	binary.Read(file,binary.LittleEndian,inf)
	fmt.Println(*inf)
	fmt.Println(inf)

}
