package main

/**
每日一题   1   两数之和
leetcode
给定一个整数数组 nums和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{i, k}
		} else {
			m[v] = k
		}
	}
	return nil
}
