package main

import (
	"algorithm-interview/structure/slice"
)

func main() {
	sliceData := slice.SliceType{1, 2, 3, 4, 5, 6}
	head := sliceData.Slice2LinkList()
	head.PrintLinkList()
}
