package recursion

func IsPalindrome(i string) bool {
	if len(i) == 0 || len(i) == 1 {
		return true
	}

	if i[0] == i[len(i)-1] {
		i = i[1 : len(i)-1]
		return IsPalindrome(i)
	}
	return false
}
