package main

import "sort"

// Two Sum - 別解集
// 複数のアプローチを提供

// アプローチ1: ハッシュマップ法（最推奨）
// 時間複雑度: O(n), 空間複雑度: O(n)
func twoSum1(nums []int, target int) []int {
	indexMap := make(map[int]int)
	
	for i, num := range nums {
		complement := target - num
		if j, exists := indexMap[complement]; exists {
			return []int{j, i}
		}
		indexMap[num] = i
	}
	
	return []int{}
}

// アプローチ2: ソート + 双方向ポインタ法
// 時間複雑度: O(n log n), 空間複雑度: O(n)
func twoSum2(nums []int, target int) []int {
	// インデックス付きのスライスを作成
	type indexedNum struct {
		value int
		index int
	}
	
	indexedNums := make([]indexedNum, len(nums))
	for i, num := range nums {
		indexedNums[i] = indexedNum{value: num, index: i}
	}
	
	// 値でソート
	sort.Slice(indexedNums, func(i, j int) bool {
		return indexedNums[i].value < indexedNums[j].value
	})
	
	// 双方向ポインタで検索
	left, right := 0, len(indexedNums)-1
	for left < right {
		sum := indexedNums[left].value + indexedNums[right].value
		if sum == target {
			// インデックスの順序を保証
			idx1, idx2 := indexedNums[left].index, indexedNums[right].index
			if idx1 < idx2 {
				return []int{idx1, idx2}
			} else {
				return []int{idx2, idx1}
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	
	return []int{}
}

// アプローチ3: ブルートフォース法（学習用）
// 時間複雑度: O(n²), 空間複雑度: O(1)
func twoSum3(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// アプローチ4: ハッシュマップ法（インデックス順序保証）
// 時間複雑度: O(n), 空間複雑度: O(n)
func twoSum4(nums []int, target int) []int {
	indexMap := make(map[int]int)
	
	for i, num := range nums {
		complement := target - num
		if j, exists := indexMap[complement]; exists {
			// 小さいインデックスを先に返す
			if j < i {
				return []int{j, i}
			} else {
				return []int{i, j}
			}
		}
		indexMap[num] = i
	}
	
	return []int{}
}

// アプローチ5: 早期リターン最適化版
// 時間複雑度: O(n), 空間複雑度: O(n)
func twoSum5(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}
	
	indexMap := make(map[int]int, len(nums))
	
	for i, num := range nums {
		complement := target - num
		if j, exists := indexMap[complement]; exists {
			return []int{j, i}
		}
		indexMap[num] = i
	}
	
	return []int{}
}

// アプローチ6: 構造体を使用した版（現在の実装の改善版）
// 時間複雑度: O(n), 空間複雑度: O(n)
func twoSum6(nums []int, target int) []int {
	type numInfo struct {
		index int
		exists bool
	}
	
	indexMap := make(map[int]numInfo)
	
	for i, num := range nums {
		complement := target - num
		if info, exists := indexMap[complement]; exists && info.exists {
			return []int{info.index, i}
		}
		indexMap[num] = numInfo{index: i, exists: true}
	}
	
	return []int{}
} 
