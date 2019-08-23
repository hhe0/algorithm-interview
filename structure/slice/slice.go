package slice

import (
	"algorithm-interview/structure/linklist"
	"fmt"
)

type SliceType []interface{}

// 判断某个元素是否在 slice 中
func (s SliceType) InSlice(a interface{}) bool {
	for _, data := range s {
		if data == a {
			return true
		}
	}

	return false
}

// 切片转链表
func (s SliceType) Slice2LinkList() *linklist.LinkNode {
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

// 反转切片
func (s SliceType) ReverseSlice() SliceType {
	temp := s[0]
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		temp = s[i]
		s[i] = s[j]
		s[j] = temp
	}

	return s
}

// 打印切片
func (s SliceType) PrintSlice() {
	fmt.Println(s)
}
