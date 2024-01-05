package simple

func Append(i []int) []int {
	ret := []int{}
	for _, v := range i {
		ret = append(ret, v)
	}
	return ret
}

func AppendRev(i []int) []int {
	ret := []int{}
	for _, v := range i {
		ret = append([]int{v}, ret...)
	}
	return ret
}
