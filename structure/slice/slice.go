package slice

import (
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
