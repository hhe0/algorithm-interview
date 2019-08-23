package main

import "algorithm-interview/structure/slice"

func main() {
	sliceData := slice.SliceType{1, 2, 3, 4, 5, 6}
	head := sliceData.Slice2LinkList()
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
