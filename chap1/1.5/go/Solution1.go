package main

import (
	"algorithm-interview/structure/slice"
	"algorithm-interview/structure/util"
	"fmt"
)

func main() {
	sliceData := slice.SliceType{1, 2, 3, 4, 5, 6, 7}
	head := util.Slice2LinkList(sliceData)
	head.PrintLinkList()

	// 转为 slice 处理
	sliceConverted := util.LinkList2Slice(head)
	fmt.Println((*sliceConverted)[4])
}
