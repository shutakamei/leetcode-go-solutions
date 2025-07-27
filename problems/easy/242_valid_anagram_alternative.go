package main

import (
	"sort"
	"strings"
)

// Valid Anagram - 別解集
// 複数のアプローチを提供

// アプローチ1: 単一マップ法（最推奨）
// 時間複雑度: O(n), 空間複雑度: O(n)
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	charCount := make(map[rune]int)
	
	// sの文字をカウントアップ
	for _, char := range s {
		charCount[char]++
	}
	
	// tの文字をカウントダウン
	for _, char := range t {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}
	
	return true
}

// アプローチ2: 配列法（小文字のみの場合、最も効率的）
// 時間複雑度: O(n), 空間複雑度: O(1)
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	charCount := make([]int, 26)
	
	for _, char := range s {
		charCount[char-'a']++
	}
	
	for _, char := range t {
		charCount[char-'a']--
		if charCount[char-'a'] < 0 {
			return false
		}
	}
	
	return true
}

// アプローチ3: ソート法（メモリ使用量削減、時間複雑度はO(n log n)）
// 時間複雑度: O(n log n), 空間複雑度: O(n)
func isAnagram3(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	// 文字列をソート
	sSorted := sortString(s)
	tSorted := sortString(t)
	
	return sSorted == tSorted
}

// 文字列をソートするヘルパー関数
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// アプローチ4: ダブルマップ法（現在の実装の改善版）
// 時間複雑度: O(n), 空間複雑度: O(n)
func isAnagram4(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	
	for _, char := range s {
		sMap[char]++
	}
	
	for _, char := range t {
		tMap[char]++
	}
	
	// マップの内容を比較
	if len(sMap) != len(tMap) {
		return false
	}
	
	for char, count := range sMap {
		if tMap[char] != count {
			return false
		}
	}
	
	return true
}

// アプローチ5: ビットマスク法（特殊なケース用）
// 制約: 各文字が最大1回しか出現しない場合
// 時間複雑度: O(n), 空間複雑度: O(1)
func isAnagram5(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	var sMask, tMask int32
	
	for _, char := range s {
		sMask |= 1 << (char - 'a')
	}
	
	for _, char := range t {
		tMask |= 1 << (char - 'a')
	}
	
	return sMask == tMask
}

// アプローチ6: 文字列分割法（現在の実装の最適化版）
// 時間複雑度: O(n), 空間複雑度: O(n)
func isAnagram6(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	
	sChars := strings.Split(s, "")
	tChars := strings.Split(t, "")
	
	sort.Strings(sChars)
	sort.Strings(tChars)
	
	return strings.Join(sChars, "") == strings.Join(tChars, "")
} 
