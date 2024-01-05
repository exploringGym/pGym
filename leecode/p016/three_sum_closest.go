package p016

import (
	"math"
	"sort"
)

// https://leetcode.com/problems/3sum-closest/

/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.

Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

Example 2:

Input: nums = [0,0,0], target = 1
Output: 0
Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).

Constraints:

    3 <= nums.length <= 500
    -1000 <= nums[i] <= 1000
    -104 <= target <= 104
*/

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	closestSum := nums[0] + nums[1] + nums[n-1]
	minDiff := math.MaxInt32

	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		l := i + 1
		r := n - 1
		for l < r {
			s := nums[i] + nums[l] + nums[r]

			if s == target {
				return s
			} else if s > target {
				r--
			} else {
				l++
			}

			d := int(math.Abs(float64(s - target)))

			if d < minDiff {
				minDiff = d
				closestSum = s
			}
		}
	}
	return closestSum
}
