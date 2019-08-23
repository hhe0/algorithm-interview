package main

import (
	"algorithm-interview/structure/slice"
)

func main() {
	sliceData := slice.SliceType{1, 2, 3, 4, 5, 6}
	head := sliceData.Slice2LinkList()
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
	linkList := sliceLinkList.Slice2LinkList()
	linkList.PrintLinkList()
}
