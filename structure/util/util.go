package util

import (
	"algorithm-interview/structure/linklist"
	"algorithm-interview/structure/slice"
)

// slice 转链表
func Slice2LinkList(s slice.SliceType) *linklist.LinkNode {
	head := linklist.NewLinkNode(0)

	tempNode := head
	for _, element := range s {
		tailNode := tempNode.AddNode(linklist.LinkNode{
			Data: element, Next: nil,
		})

		tempNode = tailNode
	}

	return head
}

// 链表转 slice
func LinkList2Slice(ln *linklist.LinkNode) *slice.SliceType {
	tempNode := ln.Next
	sliceData := slice.SliceType{}
	for tempNode != nil {
		sliceData = append(sliceData, tempNode.Data)
		tempNode = tempNode.Next
	}

	return &sliceData
}
