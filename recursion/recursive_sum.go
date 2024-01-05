package recursion

func RecursiveSum(input int) int {
	if input <= 1 {
		return input
	}
	return input + RecursiveSum(input-1)
}
