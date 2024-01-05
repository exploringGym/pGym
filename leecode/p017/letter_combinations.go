package p017

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/

/*

Given a string containing digits from 2-9 inclusive, return all possible letter
combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on the telephone buttons) is given
below. Note that 1 does not map to any letters.

Example 1:

Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:

Input: digits = ""
Output: []

Example 3:

Input: digits = "2"
Output: ["a","b","c"]

Constraints:

    0 <= digits.length <= 4
    digits[i] is a digit in the range ['2', '9'].

*/

func letterCombinations(digits string) []string {

	if digits == "" {
		return []string{}
	}

	phone := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	digitsSlice := strings.Split(digits, "")

	result := make([]string, 0)
	for _, digit := range phone[digitsSlice[0]] {
		result = append(result, digit)
	}

	for i := 1; i < len(digitsSlice); i++ {
		codeDigitsSlice := phone[digitsSlice[i]]

		tmp := []string{}

		for j := 0; j < len(result); j++ {
			for k := 0; k < len(codeDigitsSlice); k++ {
				n := []string{fmt.Sprintf("%s%s", result[j], codeDigitsSlice[k])}
				tmp = append(tmp, n...)
			}
		}

		result = tmp
	}

	return result
}
