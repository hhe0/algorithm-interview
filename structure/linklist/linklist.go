package linklist

import (
	"fmt"
)

type DataType interface{}

type LinkNode struct {
	Data DataType
	Next *LinkNode
}

// 新建链表结点
func NewLinkNode(data DataType) *LinkNode {
	return &LinkNode{
		Data: data,
		Next: nil,
	}
}

// 为链表新增结点
func (ln *LinkNode) AddNode(node LinkNode) *LinkNode {
	ln.Next = &node

	return &node
}

// 打印链表（包括头结点）
func (ln *LinkNode) PrintLinkListWithHeadNode() {
	tmpList := ln
	flag := true
	for {
		if !flag {
			break
		}

		fmt.Printf("%v ", tmpList.Data)
		if tmpList.Next == nil {
			flag = false
		}

		tmpList = tmpList.Next
	}

	fmt.Println()
}

// 打印链表（不包括头结点）
func (ln *LinkNode) PrintLinkList() {
	ln.Next.PrintLinkListWithHeadNode()
}
