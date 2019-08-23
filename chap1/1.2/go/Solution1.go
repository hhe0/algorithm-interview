package main

import (
	"algorithm-interview/structure/linklist"
	"algorithm-interview/structure/slice"
)

func main() {
	sliceData := slice.SliceType{1, 3, 1, 5, 5, 7}
	head := sliceData.Slice2LinkList()
	head.PrintLinkList()

	// 映射到 slice 去重（没办法使用map，map会无序）
	linkListSlice := make(slice.SliceType, 0)
	firstNode := head.Next
	for firstNode != nil {
		if !linkListSlice.InSlice(firstNode.Data) {
			linkListSlice = append(linkListSlice, firstNode.Data)
		}

		firstNode = firstNode.Next
	}

	// 重组
	tempNode := head
	for _, linkNode := range linkListSlice {
		tailNode := tempNode.AddNode(linklist.LinkNode{
			Data: linkNode, Next: nil,
		})
		tempNode = tailNode
	}

	head.PrintLinkList()
}
