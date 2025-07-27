package main

import (
	"testing"
)

// ベンチマークテスト - 各アプローチの性能比較
func BenchmarkContainsDuplicate1(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1} // 重複あり
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate1(nums)
	}
}

func BenchmarkContainsDuplicate2(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1} // 重複あり
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate2(nums)
	}
}

func BenchmarkContainsDuplicate3(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1} // 重複あり
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// ソート版は配列を変更するため、コピーを作成
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		containsDuplicate3(numsCopy)
	}
}

func BenchmarkContainsDuplicate4(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1} // 重複あり
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate4(nums)
	}
}

// 重複なしの場合のベンチマーク
func BenchmarkContainsDuplicate1NoDuplicate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 重複なし
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate1(nums)
	}
}

func BenchmarkContainsDuplicate2NoDuplicate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 重複なし
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate2(nums)
	}
}

func BenchmarkContainsDuplicate4NoDuplicate(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 重複なし
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate4(nums)
	}
}

// 大きな配列でのベンチマーク
func BenchmarkContainsDuplicateLargeArray(b *testing.B) {
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i
	}
	nums[9999] = 0 // 重複を作成
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate4(nums)
	}
} 
