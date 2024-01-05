package recursion

import "sort"

func BinarySearch(array []int, target int) int {
	sort.Ints(array)
	return BinarySearchRec(array, 0, len(array)-1, target)
}

func BinarySearchRec(array []int, left, right, target int) int {
	inboundL := left >= 0 && left <= right
	inboundR := right <= len(array)-1 && right >= left

	if !inboundL || !inboundR {
		return -1
	}

	mid := (left + right) / 2

	if target == array[mid] {
		return mid
	}

	if target < array[mid] {
		return BinarySearchRec(array, left, mid-1, target)
	} else {
		return BinarySearchRec(array, mid+1, right, target)
	}
}
