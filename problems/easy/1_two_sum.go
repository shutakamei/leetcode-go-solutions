package main

import "fmt"

// Two Sum - LeetCode Problem 1
// https://leetcode.com/problems/two-sum/
//
// 問題: 整数配列 nums と整数 target が与えられたとき、
// 合計が target になる2つの数のインデックスを返す。
// 各入力には正確に1つの解があると仮定する。

func twoSum(nums []int, target int) []int {
	// ハッシュマップを使用してO(n)の時間複雑度で解く
	numMap := make(map[int]int)
	
	for i, num := range nums {
		complement := target - num
		
		// 補数が既にマップに存在するかチェック
		if j, exists := numMap[complement]; exists {
			return []int{j, i}
		}
		
		// 現在の数とそのインデックスをマップに追加
		numMap[num] = i
	}
	
	// 解が見つからない場合（問題の制約上、実際には到達しない）
	return []int{}
}

// テスト用の関数
func runTwoSum() {
	fmt.Println("🔍 Two Sum Problem")
	fmt.Println("==================")
	
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}
	
	for i, tc := range testCases {
		result := twoSum(tc.nums, tc.target)
		fmt.Printf("テストケース %d: nums=%v, target=%d\n", i+1, tc.nums, tc.target)
		fmt.Printf("  結果: %v\n", result)
		fmt.Printf("  期待: %v\n", tc.want)
		fmt.Printf("  正解: %t\n\n", len(result) == 2 && 
			((result[0] == tc.want[0] && result[1] == tc.want[1]) ||
			 (result[0] == tc.want[1] && result[1] == tc.want[0])))
	}
} 
