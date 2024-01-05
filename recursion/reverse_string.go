package recursion

func ReverseString(input string) string {
	if len(input) == 0 {
		return ""
	}

	sub := input[1:]
	char := string(input[0])
	return ReverseString(sub) + char
}
