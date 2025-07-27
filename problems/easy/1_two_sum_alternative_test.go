package main

import (
	"testing"
)

// 全アプローチのテスト
func TestAllTwoSumApproaches(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1",
			nums:   []int{3, 4, 5, 6},
			target: 7,
			want:   []int{0, 1},
		},
		{
			name:   "Example 2",
			nums:   []int{4, 5, 6},
			target: 10,
			want:   []int{0, 2},
		},
		{
			name:   "Example 3",
			nums:   []int{5, 5},
			target: 10,
			want:   []int{0, 1},
		},
		{
			name:   "Custom test 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Negative numbers",
			nums:   []int{-1, -2, -3, -4},
			target: -7,
			want:   []int{2, 3},
		},
		{
			name:   "Large numbers",
			nums:   []int{1000000, 2000000, 3000000},
			target: 4000000,
			want:   []int{0, 2},
		},
		{
			name:   "Same numbers",
			nums:   []int{3, 3, 3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "Edge case - two elements",
			nums:   []int{1, 2},
			target: 3,
			want:   []int{0, 1},
		},
		{
			name:   "Edge case - large array",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 19,
			want:   []int{8, 9},
		},
	}

	approaches := []struct {
		name string
		fn   func([]int, int) []int
	}{
		{"Hash Map", twoSum1},
		{"Sort + Two Pointers", twoSum2},
		{"Brute Force", twoSum3},
		{"Hash Map with Order", twoSum4},
		{"Early Return", twoSum5},
		{"Struct Based", twoSum6},
	}

	for _, approach := range approaches {
		t.Run(approach.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := approach.fn(tt.nums, tt.target)
					
					// 結果の長さをチェック
					if len(got) != 2 {
						t.Errorf("%s: %s() = %v, want length 2", 
							approach.name, approach.name, got)
						return
					}
					
					// 実際に合計がtargetになるかチェック
					sum := tt.nums[got[0]] + tt.nums[got[1]]
					if sum != tt.target {
						t.Errorf("%s: %s() returned indices %v, but nums[%d] + nums[%d] = %d + %d = %d, want %d",
							approach.name, approach.name, got, got[0], got[1], 
							tt.nums[got[0]], tt.nums[got[1]], sum, tt.target)
					}
					
					// インデックスが異なることをチェック
					if got[0] == got[1] {
						t.Errorf("%s: %s() returned same indices %v", 
							approach.name, approach.name, got)
					}
				})
			}
		})
	}
}

// 各アプローチのベンチマークテスト
func BenchmarkTwoSum1(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 2, 4, 3, 3, 1, 5, 8, 10, 13}
	target := 9
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum1(nums, target)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 2, 4, 3, 3, 1, 5, 8, 10, 13}
	target := 9
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum2(nums, target)
	}
}

func BenchmarkTwoSum3(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 2, 4, 3, 3, 1, 5, 8, 10, 13}
	target := 9
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum3(nums, target)
	}
}

func BenchmarkTwoSum4(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 2, 4, 3, 3, 1, 5, 8, 10, 13}
	target := 9
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum4(nums, target)
	}
}

func BenchmarkTwoSum5(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 2, 4, 3, 3, 1, 5, 8, 10, 13}
	target := 9
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum5(nums, target)
	}
}

func BenchmarkTwoSum6(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 2, 4, 3, 3, 1, 5, 8, 10, 13}
	target := 9
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum6(nums, target)
	}
}

// 大きな配列でのベンチマーク
func BenchmarkTwoSumLarge1(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	target := 1998
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum1(nums, target)
	}
}

func BenchmarkTwoSumLarge2(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	target := 1998
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum2(nums, target)
	}
}

func BenchmarkTwoSumLarge3(b *testing.B) {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	target := 1998
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum3(nums, target)
	}
} 
