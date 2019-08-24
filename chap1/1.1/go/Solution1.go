package main

import (
	"algorithm-interview/structure/slice"
	"algorithm-interview/structure/util"
)

func main() {
	sliceData := slice.SliceType{1, 2, 3, 4, 5, 6}
	head := util.Slice2LinkList(sliceData)
	head.PrintLinkList()

	// 转 slice 处理
	tempNode := head.Next
	sliceLinkList := slice.SliceType{}
	for tempNode != nil {
		sliceLinkList = append(sliceLinkList, tempNode.Data)
		tempNode = tempNode.Next
	}

	// 反转 slice 元素的顺序
	sliceLinkList = sliceLinkList.ReverseSlice()

	// 转 linklist
	linkList := util.Slice2LinkList(sliceLinkList)
	linkList.PrintLinkList()
}
