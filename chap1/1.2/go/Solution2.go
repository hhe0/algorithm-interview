package main

import (
	"algorithm-interview/structure/slice"
)

func main() {
	sliceData := slice.SliceType{1, 3, 1, 5, 5}
	head := sliceData.Slice2LinkList()
	head.PrintLinkList()

	// 映射到 slice 去重（没办法使用map，map会无序）
	linkListSlice := make(slice.SliceType, 0)
	curNode := head
	for curNode.Next != nil {
		if !linkListSlice.InSlice(curNode.Next.Data) {
			linkListSlice = append(linkListSlice, curNode.Next.Data)

			curNode = curNode.Next
		} else {
			// 将重复的结点直接删除
			curNode.Next = curNode.Next.Next
		}
	}

	head.PrintLinkList()
}
