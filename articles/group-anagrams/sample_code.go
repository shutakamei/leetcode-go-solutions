package main

import (
	"reflect"
	"sort"
	"strings"
)

// 基本的な解法（文字カウント + マップ比較）
// 時間計算量: O(n² × k) ここで n は文字列の数、k は最大文字列長
// 空間計算量: O(n × k)
func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 {
		if len(strs) == 0 {
			return [][]string{}
		}
		return [][]string{strs}
	}
	
	var anagramArr [][]string
	var charMapArr []map[rune]int
	
	for _, str := range strs {
		charMap := make(map[rune]int)
		for _, char := range str {
			charMap[char]++
		}
		
		newCharFlag := true
		for i, cm := range charMapArr {
			if reflect.DeepEqual(cm, charMap) {
				anagramArr[i] = append(anagramArr[i], str)
				newCharFlag = false
				break
			}
		}
		
		if newCharFlag {
			charMapArr = append(charMapArr, charMap)
			anagramArr = append(anagramArr, []string{str})
		}
	}
	
	return anagramArr
}

// 最適化版1: ソート方式
// 時間計算量: O(n × k × log k) ここで n は文字列の数、k は最大文字列長
// 空間計算量: O(n × k)
func groupAnagramsSorted(strs []string) [][]string {
	if len(strs) <= 1 {
		if len(strs) == 0 {
			return [][]string{}
		}
		return [][]string{strs}
	}

	groups := make(map[string][]string)
	
	for _, str := range strs {
		canonical := getSortedString(str)
		groups[canonical] = append(groups[canonical], str)
	}
	
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

// 文字列をソートして正規形を取得
func getSortedString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

// 最適化版2: 文字カウント文字列化方式（推奨）
// 時間計算量: O(n × k) ここで n は文字列の数、k は最大文字列長
// 空間計算量: O(n × k)
func groupAnagramsCountString(strs []string) [][]string {
	if len(strs) <= 1 {
		if len(strs) == 0 {
			return [][]string{}
		}
		return [][]string{strs}
	}

	groups := make(map[string][]string)
	
	for _, str := range strs {
		canonical := getCountString(str)
		groups[canonical] = append(groups[canonical], str)
	}
	
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

// 文字カウントを文字列として表現（例: "a1b2c3"）
func getCountString(s string) string {
	count := make([]int, 26)
	
	for _, char := range s {
		count[char-'a']++
	}
	
	var result strings.Builder
	for i := 0; i < 26; i++ {
		if count[i] > 0 {
			result.WriteByte(byte('a' + i))
			result.WriteString(string(rune(count[i] + '0')))
		}
	}
	
	return result.String()
}

// 最適化版3: 文字カウント配列方式
// 時間計算量: O(n × k) ここで n は文字列の数、k は最大文字列長
// 空間計算量: O(n × k)
func groupAnagramsCountArray(strs []string) [][]string {
	if len(strs) <= 1 {
		if len(strs) == 0 {
			return [][]string{}
		}
		return [][]string{strs}
	}

	groups := make(map[string][]string)
	
	for _, str := range strs {
		canonical := getCountArrayString(str)
		groups[canonical] = append(groups[canonical], str)
	}
	
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}
	
	return result
}

// 文字カウント配列を文字列として表現（例: "1#0#2#..."）
func getCountArrayString(s string) string {
	count := make([]int, 26)
	
	for _, char := range s {
		count[char-'a']++
	}
	
	var result strings.Builder
	for i := 0; i < 26; i++ {
		if i > 0 {
			result.WriteByte('#')
		}
		result.WriteString(string(rune(count[i] + '0')))
	}
	
	return result.String()
}

// 使用例
func main() {
	// テストケース
	testCases := []struct {
		name string
		input []string
	}{
		{
			name: "例1",
			input: []string{"act", "pots", "tops", "cat", "stop", "hat"},
		},
		{
			name: "例2", 
			input: []string{"x"},
		},
		{
			name: "例3",
			input: []string{""},
		},
		{
			name: "複数のアナグラムグループ",
			input: []string{"eat", "tea", "ate", "nat", "bat", "tan"},
		},
	}

	for _, tc := range testCases {
		println("=== " + tc.name + " ===")
		println("入力:", strings.Join(tc.input, ", "))
		
		// 各解法で実行
		result1 := groupAnagrams(tc.input)
		result2 := groupAnagramsSorted(tc.input)
		result3 := groupAnagramsCountString(tc.input)
		result4 := groupAnagramsCountArray(tc.input)
		
		println("基本実装:", formatResult(result1))
		println("ソート方式:", formatResult(result2))
		println("文字カウント文字列化:", formatResult(result3))
		println("文字カウント配列:", formatResult(result4))
		println()
	}
}

// 結果をフォーマットするヘルパー関数
func formatResult(result [][]string) string {
	if len(result) == 0 {
		return "[]"
	}
	
	var parts []string
	for _, group := range result {
		parts = append(parts, "["+strings.Join(group, ", ")+"]")
	}
	return "[" + strings.Join(parts, ", ") + "]"
} 
