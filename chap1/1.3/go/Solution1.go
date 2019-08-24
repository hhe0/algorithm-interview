package main

import (
	"algorithm-interview/structure/linklist"
	"algorithm-interview/structure/slice"
	"algorithm-interview/structure/util"
	"math"
	"strconv"
)

func main() {
	sliceData := slice.SliceType{3, 1, 5}
	headA := util.Slice2LinkList(sliceData)

	sliceData = slice.SliceType{7, 9, 2}
	headB := util.Slice2LinkList(sliceData)

	// 链表转为数字
	numA := 0
	countA := 0
	tempA := headA.Next
	for tempA != nil {
		numA += tempA.Data.(int) * int(math.Pow10(countA))
		tempA = tempA.Next
		countA++
	}

	numB := 0
	tempB := headB.Next
	countB := 0
	for tempB != nil {
		numB += tempB.Data.(int) * int(math.Pow10(countB))
		tempB = tempB.Next
		countB++
	}

	// 数字求和
	res := numA + numB

	// 数字转链表
	headC := linklist.NewLinkNode(0)
	tempNode := headC
	resString := strconv.Itoa(res)
	for i := len(resString) - 1; i >= 0; i-- {
		num, _ := strconv.Atoi(string(resString[i]))
		tailNode := tempNode.AddNode(linklist.LinkNode{
			Data: num, Next: nil,
		})

		tempNode = tailNode
	}

	headA.PrintLinkList()
	headB.PrintLinkList()
	headC.PrintLinkList()
}
