package main

import (
	"algorithm-interview/structure/linklist"
	"fmt"
	"math"
	"strconv"
)

func main() {
	headA := linklist.NewLinkNode(0)
	headA.AddNode(linklist.LinkNode{
		Data: 3, Next: nil,
	}).AddNode(linklist.LinkNode{
		Data: 1, Next: nil,
	}).AddNode(linklist.LinkNode{
		Data: 5, Next: nil,
	})

	headB := linklist.NewLinkNode(0)
	headB.AddNode(linklist.LinkNode{
		Data: 7, Next: nil,
	}).AddNode(linklist.LinkNode{
		Data: 9, Next: nil,
	}).AddNode(linklist.LinkNode{
		Data: 2, Next: nil,
	})

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
		fmt.Println(num)
		tailNode := tempNode.AddNode(linklist.LinkNode{
			Data: num, Next: nil,
		})

		tempNode = tailNode
	}

	headA.PrintLinkList()
	headB.PrintLinkList()
	headC.PrintLinkList()
}
