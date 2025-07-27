package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "Example 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "Custom test 1",
			nums:   []int{1, 5, 8, 10, 13},
			target: 18,
			want:   []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			
			// 結果の長さをチェック
			if len(got) != 2 {
				t.Errorf("twoSum() = %v, want length 2", got)
				return
			}
			
			// 結果が正しいかチェック（順序は関係ない）
			valid := (reflect.DeepEqual(got, tt.want) ||
				reflect.DeepEqual(got, []int{tt.want[1], tt.want[0]}))
			
			if !valid {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
			
			// 実際に合計がtargetになるかチェック
			sum := tt.nums[got[0]] + tt.nums[got[1]]
			if sum != tt.target {
				t.Errorf("twoSum() returned indices %v, but nums[%d] + nums[%d] = %d + %d = %d, want %d",
					got, got[0], got[1], tt.nums[got[0]], tt.nums[got[1]], sum, tt.target)
			}
		})
	}
}

// ベンチマークテスト
func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15, 3, 2, 4, 3, 3, 1, 5, 8, 10, 13}
	target := 9
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
} 
