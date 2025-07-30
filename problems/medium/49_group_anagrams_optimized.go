package main

import (
	"sort"
	"strings"
)

// 最適化された解法1: ソート方式
// 時間計算量: O(n * k * log k) ここで n は文字列の数、k は最大文字列長
// 空間計算量: O(n * k)
// 利点: 実装がシンプル、短い文字列に効率的
func groupAnagramsSorted(strs []string) [][]string {
	if len(strs) <= 1 {
		if len(strs) == 0 {
			return [][]string{}
		}
		return [][]string{strs}
	}

	// マップを使用してグループ化
	// キー: ソートされた文字列、値: 元の文字列のスライス
	groups := make(map[string][]string)
	
	for _, str := range strs {
		// 文字列をソートして正規形を作成
		canonical := getSortedString(str)
		groups[canonical] = append(groups[canonical], str)
	}
	
	// マップの値を結果スライスに変換
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

// 最適化された解法2: 文字カウント文字列化方式
// 時間計算量: O(n * k) ここで n は文字列の数、k は最大文字列長
// 空間計算量: O(n * k)
// 利点: 長い文字列に効率的、ソート不要
func groupAnagramsCountString(strs []string) [][]string {
	if len(strs) <= 1 {
		if len(strs) == 0 {
			return [][]string{}
		}
		return [][]string{strs}
	}

	groups := make(map[string][]string)
	
	for _, str := range strs {
		// 文字カウントを文字列として表現
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

// 最適化された解法3: 文字カウント配列方式（最も効率的）
// 時間計算量: O(n * k)
// 空間計算量: O(n * k)
// 利点: 最も効率的、メモリ使用量も最適
func groupAnagramsCountArray(strs []string) [][]string {
	if len(strs) <= 1 {
		if len(strs) == 0 {
			return [][]string{}
		}
		return [][]string{strs}
	}

	groups := make(map[string][]string)
	
	for _, str := range strs {
		// 文字カウント配列を文字列として表現
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
