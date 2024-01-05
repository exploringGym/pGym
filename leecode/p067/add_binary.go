package p067

// https://leetcode.com/problems/add-binary/

/*
Given two binary strings a and b, return their sum as a binary string.

Example 1:

Input: a = "11", b = "1" Output: "100"

Example 2:

Input: a = "1010", b = "1011" Output: "10101"

Constraints:

1 <= a.length, b.length <= 104 a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.
*/

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	res, carry := []byte(a), 0

	for idA, idB := len(a)-1, len(b)-1; idA >= 0; idA, idB = idA-1, idB-1 {
		if a[idA] == '1' {
			carry++
		}
		if idB >= 0 && b[idB] == '1' {
			carry++
		}
		switch carry {
		case 1:
			res[idA], carry = '1', 0
		case 2:
			res[idA], carry = '0', 1
		case 3:
			res[idA], carry = '1', 1
		}
	}

	if carry == 1 {
		tmp := res
		res := make([]byte, len(tmp)+1)
		copy(res[1:], tmp)
		res[0] = '1'
		return string(res)
	}

	return string(res)
}
