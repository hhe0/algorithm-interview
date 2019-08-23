package main

import (
	linklist2 "algorithm-interview/structure/linklist"
)

func main() {
	head := linklist2.NewLinkNode(0)

	head.AddNode(linklist2.LinkNode{
		Data: 1, Next: nil,
	}).AddNode(linklist2.LinkNode{
		Data: 2, Next: nil,
	}).AddNode(linklist2.LinkNode{
		Data: 3, Next: nil,
	}).AddNode(linklist2.LinkNode{
		Data: 4, Next: nil,
	}).AddNode(linklist2.LinkNode{
		Data: 5, Next: nil,
	}).AddNode(linklist2.LinkNode{
		Data: 6, Next: nil,
	})

	// 0 -> 1 -> 2
	head.PrintLinkList()
}
