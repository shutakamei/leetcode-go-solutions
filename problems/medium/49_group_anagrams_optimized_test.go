package main

import (
	"reflect"
	"sort"
	"testing"
)

// 結果を比較するためのヘルパー関数
func sortResultOptimized(result [][]string) [][]string {
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
func areResultsEqualOptimized(result1, result2 [][]string) bool {
	sorted1 := sortResultOptimized(result1)
	sorted2 := sortResultOptimized(result2)
	return reflect.DeepEqual(sorted1, sorted2)
}

// 最適化された解法1のテスト
func TestGroupAnagramsSorted(t *testing.T) {
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
			result := groupAnagramsSorted(tt.input)
			
			if !areResultsEqualOptimized(result, tt.expected) {
				t.Errorf("groupAnagramsSorted() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// 最適化された解法2のテスト
func TestGroupAnagramsCountString(t *testing.T) {
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
			result := groupAnagramsCountString(tt.input)
			
			if !areResultsEqualOptimized(result, tt.expected) {
				t.Errorf("groupAnagramsCountString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// 最適化された解法3のテスト
func TestGroupAnagramsCountArray(t *testing.T) {
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
			result := groupAnagramsCountArray(tt.input)
			
			if !areResultsEqualOptimized(result, tt.expected) {
				t.Errorf("groupAnagramsCountArray() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// ヘルパー関数のテスト
func TestGetSortedString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"act", "act"},
		{"cat", "act"},
		{"pots", "opst"},
		{"tops", "opst"},
		{"stop", "opst"},
		{"", ""},
		{"a", "a"},
		{"abc", "abc"},
		{"cba", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := getSortedString(tt.input)
			if result != tt.expected {
				t.Errorf("getSortedString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetCountString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"act", "a1c1t1"},
		{"cat", "a1c1t1"},
		{"pots", "o1p1s1t1"},
		{"tops", "o1p1s1t1"},
		{"stop", "o1p1s1t1"},
		{"", ""},
		{"a", "a1"},
		{"abc", "a1b1c1"},
		{"cba", "a1b1c1"},
		{"aaa", "a3"},
		{"aab", "a2b1"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := getCountString(tt.input)
			if result != tt.expected {
				t.Errorf("getCountString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetCountArrayString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"aaa", "3#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0"},
		{"", "0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := getCountArrayString(tt.input)
			if result != tt.expected {
				t.Errorf("getCountArrayString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// ベンチマークテスト
func BenchmarkGroupAnagramsSorted(b *testing.B) {
	input := []string{"act", "pots", "tops", "cat", "stop", "hat", "eat", "tea", "ate", "nat", "bat", "tan"}
	for i := 0; i < b.N; i++ {
		groupAnagramsSorted(input)
	}
}

func BenchmarkGroupAnagramsCountString(b *testing.B) {
	input := []string{"act", "pots", "tops", "cat", "stop", "hat", "eat", "tea", "ate", "nat", "bat", "tan"}
	for i := 0; i < b.N; i++ {
		groupAnagramsCountString(input)
	}
}

func BenchmarkGroupAnagramsCountArray(b *testing.B) {
	input := []string{"act", "pots", "tops", "cat", "stop", "hat", "eat", "tea", "ate", "nat", "bat", "tan"}
	for i := 0; i < b.N; i++ {
		groupAnagramsCountArray(input)
	}
}

// 長い文字列でのベンチマーク
func BenchmarkGroupAnagramsLongStrings(b *testing.B) {
	input := []string{
		"abcdefghijklmnopqrstuvwxyz",
		"zyxwvutsrqponmlkjihgfedcba",
		"abcdefghijklmnopqrstuvwxyz",
		"aaaaaaaaaaaaaaaaaaaaaaaaaa",
		"zzzzzzzzzzzzzzzzzzzzzzzzzz",
		"abcdefghijklmnopqrstuvwxyz",
	}
	
	b.Run("Sorted", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			groupAnagramsSorted(input)
		}
	})
	
	b.Run("CountString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			groupAnagramsCountString(input)
		}
	})
	
	b.Run("CountArray", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			groupAnagramsCountArray(input)
		}
	})
} 
