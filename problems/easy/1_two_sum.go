package main

// Two Sum - LeetCode Problem 1
// https://leetcode.com/problems/two-sum/
//
// 問題: 整数配列 nums と整数 target が与えられたとき、
// nums[i] + nums[j] == target かつ i != j となる
// インデックス i と j を返す。
// 各入力には正確に1つの解があると仮定する。
// 答えは小さいインデックスを先に返す。
//
// 制約:
// - 2 <= nums.length <= 1000
// - -10,000,000 <= nums[i] <= 10,000,000
// - -10,000,000 <= target <= 10,000,000

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]struct{Index int; Exist bool})
	for i, num := range nums {
		if indexMap[target-num].Exist {
			return []int{indexMap[target-num].Index, i}
		}
		indexMap[num] = struct{Index int; Exist bool}{Index: i, Exist: true}
	}

	return []int{}
}

// 完全にこっちのほうがいい
// func twoSum(nums []int, target int) []int {
// 	indexMap := make(map[int]int)
// 	for i, num := range nums {
// 		if j, exists := indexMap[target-num]; exists {
// 			return []int{j, i}
// 		}
// 		indexMap[num] = i
// 	}
// 	return []int{}
// } 
