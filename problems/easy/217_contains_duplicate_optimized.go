package main

// Contains Duplicate - 最適化版
// 複数のアプローチを提供

// アプローチ1: 現在の実装（ハッシュマップ + bool）
func containsDuplicate1(nums []int) bool {
	intMap := make(map[int]bool)
	for _, num := range nums {
		if intMap[num] {
			return true
		}
		intMap[num] = true
	}
	return false
}

// アプローチ2: メモリ最適化版（struct{}を使用）
func containsDuplicate2(nums []int) bool {
	intMap := make(map[int]struct{})
	for _, num := range nums {
		if _, exists := intMap[num]; exists {
			return true
		}
		intMap[num] = struct{}{}
	}
	return false
}

// アプローチ3: ソートベース（メモリ使用量削減、時間複雑度はO(n log n)）
func containsDuplicate3(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	
	// ソート（in-place）
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	
	// 隣接要素をチェック
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// アプローチ4: 最適化されたハッシュマップ（初期容量指定）
func containsDuplicate4(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	
	// 初期容量を指定してメモリ割り当てを最適化
	intMap := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, exists := intMap[num]; exists {
			return true
		}
		intMap[num] = struct{}{}
	}
	return false
} 
