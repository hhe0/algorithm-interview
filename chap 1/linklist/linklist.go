package linklist

type Element interface{}

type LinkNode struct {
	Data Element
	Next *LinkNode
}

func NewLinkNode(data Element) *LinkNode {
	return &LinkNode{
		Data: data,
		Next: nil,
	}
}