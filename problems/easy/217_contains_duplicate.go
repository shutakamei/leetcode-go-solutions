package main

// Contains Duplicate - LeetCode Problem 217
// https://leetcode.com/problems/contains-duplicate/
//
// 問題: 整数配列 nums が与えられたとき、
// 配列内に重複する値が存在する場合は true を返し、
// そうでなければ false を返す。
//
// 制約:
// - 1 <= nums.length <= 10^5
// - -10^9 <= nums[i] <= 10^9

func containsDuplicate(nums []int) bool {
	intMap := make(map[int]bool)
	for _, num := range nums {
		if intMap[num] {
			return true
		}
		intMap[num] = true
	}
	return false
}

