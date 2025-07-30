package main

import "reflect"

// Problem 49 - Group Anagrams (アナグラムのグループ化)
// https://leetcode.com/problems/group-anagrams/

// 問題の説明:
// 文字列の配列 strs が与えられ、すべてのアナグラムをサブリストにグループ化します。
// 出力の順序は任意です。
// 
// アナグラムとは、別の文字列と同じ文字を正確に含む文字列ですが、
// 文字の順序が異なる場合があります。
//
// 制約:
// - 1 <= strs.length <= 1000
// - 0 <= strs[i].length <= 100
// - strs[i] は小文字の英字で構成されています
//
// 例:
// 入力: ["act","pots","tops","cat","stop","hat"]
// 出力: [["hat"],["act", "cat"],["stop", "pots", "tops"]]

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
			}
		}
		if newCharFlag {
			charMapArr = append(charMapArr, charMap)
			anagramArr = append(anagramArr, []string{str})
		}
	}
	return anagramArr
}

// アナグラムが一致した配列に追加する処理の実現が難しい
