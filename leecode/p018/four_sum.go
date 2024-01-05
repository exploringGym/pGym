package p018

import "sort"

// https://leetcode.com/problems/4sum/

/*
Given an array nums of n integers, return an array of all the unique quadruplets
[nums[a], nums[b], nums[c], nums[d]] such that:

	0 <= a, b, c, d < n
	a, b, c, and d are distinct.
	nums[a] + nums[b] + nums[c] + nums[d] == target

You may return the answer in any order.

Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

Example 2:

Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]

Constraints:

	1 <= nums.length <= 200
	-109 <= nums[i] <= 109
	-109 <= target <= 109
*/
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	ret := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < n; i++ {
		// no repeated elements for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n; j++ {
			// no repeated elements for j
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}

			l := j + 1
			r := n - 1

			for l < r {

				s := nums[i] + nums[j] + nums[l] + nums[r]

				if s < target {
					l++
					continue
				} else if s > target {
					r--
					continue
				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})

					l++
					r--

					for l < len(nums) && nums[l] == nums[l-1] {
						l++
					}
					for r > 0 && nums[r] == nums[r+1] {
						r--
					}
				}
			}
		}
	}
	return ret
}
