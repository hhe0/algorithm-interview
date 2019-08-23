package main

import (
	"algorithm-interview/structure/linklist"
	"algorithm-interview/structure/slice"
)

func main() {
	head := linklist.NewLinkNode(0)

	head.AddNode(linklist.LinkNode{
		Data: 1, Next: nil,
	}).AddNode(linklist.LinkNode{
		Data: 3, Next: nil,
	}).AddNode(linklist.LinkNode{
		Data: 1, Next: nil,
	}).AddNode(linklist.LinkNode{
		Data: 5, Next: nil,
	}).AddNode(linklist.LinkNode{
		Data: 5, Next: nil,
	}).AddNode(linklist.LinkNode{
		Data: 7, Next: nil,
	})
	head.PrintLinkList()

	// 映射到 slice 去重（没办法使用map，map会无序）
	linkListSlice := make([]interface{}, 0)
	firstNode := head.Next
	for firstNode != nil {
		if !slice.InSlice(firstNode.Data, linkListSlice) {
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
