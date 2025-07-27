package main

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1 - Contains duplicate",
			nums: []int{1, 2, 3, 3},
			want: true,
		},
		{
			name: "Example 2 - No duplicate",
			nums: []int{1, 2, 3, 4},
			want: false,
		},
		{
			name: "Single element",
			nums: []int{1},
			want: false,
		},
		{
			name: "Two same elements",
			nums: []int{1, 1},
			want: true,
		},
		{
			name: "Multiple duplicates",
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			want: true,
		},
		{
			name: "Negative numbers with duplicate",
			nums: []int{-1, -2, -1, 3},
			want: true,
		},
		{
			name: "Negative numbers without duplicate",
			nums: []int{-1, -2, -3, -4},
			want: false,
		},
		{
			name: "Large numbers",
			nums: []int{1000000000, -1000000000, 1000000000},
			want: true,
		},
		{
			name: "Empty array",
			nums: []int{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ベンチマークテスト
func BenchmarkContainsDuplicate(b *testing.B) {
	// 重複ありの配列
	numsWithDuplicate := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate(numsWithDuplicate)
	}
}

func BenchmarkContainsDuplicateNoDuplicate(b *testing.B) {
	// 重複なしの配列
	numsWithoutDuplicate := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		containsDuplicate(numsWithoutDuplicate)
	}
} 
