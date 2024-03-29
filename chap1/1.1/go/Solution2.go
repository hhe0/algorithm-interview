package main

import (
	"algorithm-interview/structure/slice"
	"algorithm-interview/structure/util"
)

func main() {
	sliceData := slice.SliceType{1, 2, 3, 4, 5, 6}
	head := util.Slice2LinkList(sliceData)
	head.PrintLinkList()

	tempINode := head.Next
	for tempINode != nil {
		tempJNode := tempINode.Next
		for tempJNode != nil {
			tempData := tempINode.Data
			tempINode.Data = tempJNode.Data
			tempJNode.Data = tempData

			tempJNode = tempJNode.Next
		}

		tempINode = tempINode.Next
	}

	head.PrintLinkList()
}
