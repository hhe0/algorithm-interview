package slice

// 判断某个元素是否在 slice 中
func InSlice(a interface{}, s []interface{}) bool {
	for _, data := range s {
		if data == a {
			return true
		}
	}

	return false
}
