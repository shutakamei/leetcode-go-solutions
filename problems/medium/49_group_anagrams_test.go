package main

import (
	"reflect"
	"sort"
	"testing"
)

// 結果を比較するためのヘルパー関数
// グループとグループ内の文字列の順序は重要ではないため、ソートして比較します
func sortResult(result [][]string) [][]string {
	// 各グループ内の文字列をソート
	for i := range result {
		sort.Strings(result[i])
	}
	
	// グループを最初の文字列でソート
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) == 0 && len(result[j]) == 0 {
			return false
		}
		if len(result[i]) == 0 {
			return true
		}
		if len(result[j]) == 0 {
			return false
		}
		return result[i][0] < result[j][0]
	})
	
	return result
}

// 2つの結果が等しいかチェックするヘルパー関数
func areResultsEqual(result1, result2 [][]string) bool {
	sorted1 := sortResult(result1)
	sorted2 := sortResult(result2)
	return reflect.DeepEqual(sorted1, sorted2)
}

// groupAnagrams関数のテスト
func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "例1",
			input:    []string{"act", "pots", "tops", "cat", "stop", "hat"},
			expected: [][]string{{"hat"}, {"act", "cat"}, {"stop", "pots", "tops"}},
		},
		{
			name:     "例2",
			input:    []string{"x"},
			expected: [][]string{{"x"}},
		},
		{
			name:     "例3",
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			name:     "空の入力",
			input:    []string{},
			expected: [][]string{},
		},
		{
			name:     "すべて同じ文字列",
			input:    []string{"abc", "abc", "abc"},
			expected: [][]string{{"abc", "abc", "abc"}},
		},
		{
			name:     "アナグラムなし",
			input:    []string{"abc", "def", "ghi"},
			expected: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name:     "複数のアナグラムグループ",
			input:    []string{"eat", "tea", "ate", "nat", "bat", "tan"},
			expected: [][]string{{"eat", "tea", "ate"}, {"nat", "tan"}, {"bat"}},
		},
		{
			name:     "単一文字の文字列",
			input:    []string{"a", "b", "c", "a", "b"},
			expected: [][]string{{"a", "a"}, {"b", "b"}, {"c"}},
		},
		{
			name:     "空文字列が混在",
			input:    []string{"", "a", "", "b"},
			expected: [][]string{{"", ""}, {"a"}, {"b"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.input)
			
			if !areResultsEqual(result, tt.expected) {
				t.Errorf("groupAnagrams() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// ベンチマークテスト（実装後に実行）
func BenchmarkGroupAnagrams(b *testing.B) {
	input := []string{"act", "pots", "tops", "cat", "stop", "hat", "eat", "tea", "ate", "nat", "bat", "tan"}
	for i := 0; i < b.N; i++ {
		groupAnagrams(input)
	}
} 
