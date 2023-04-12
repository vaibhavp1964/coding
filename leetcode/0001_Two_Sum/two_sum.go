package twosum

/**
Problem can be found at - https://leetcode.com/problems/two-sum/description/

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {
	mapping := make(map[int]int)

	for index, val := range nums {
		toSearch := target - val
		searchIndex, ok := mapping[toSearch]
		if ok {
			return []int{searchIndex, index}
		}
		mapping[val] = index
	}

	return []int{-1, -1}
}
