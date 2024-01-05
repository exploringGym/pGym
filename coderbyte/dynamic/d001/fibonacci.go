package d001

// 1 1 2 3 5 7

func Fibonacci(target int) int {
	if target <= 2 {
		return 1
	}
	return Fibonacci(target-1) + Fibonacci(target-2)
}
