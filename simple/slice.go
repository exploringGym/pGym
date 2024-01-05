package simple

func SliceFront(x []int, cut int) []int {
	r := x[cut:] // it will include (keep) cut
	return r
}

func SliceEnd(x []int, cut int) []int {
	r := x[:cut] // it won't include cut
	return r
}
