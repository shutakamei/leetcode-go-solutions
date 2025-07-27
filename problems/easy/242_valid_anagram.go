package main

import "strings"


// Valid Anagram - LeetCode Problem 242
// https://leetcode.com/problems/valid-anagram/
//
// 問題: 2つの文字列 s と t が与えられたとき、
// これらがアナグラムである場合は true を返し、
// そうでなければ false を返す。
//
// アナグラムとは、文字列が他の文字列と同じ文字を
// 同じ回数だけ含んでいるが、文字の順序が異なる文字列のこと。
//
// 制約:
// - s と t は小文字の英字のみで構成される
// - 1 <= s.length, t.length <= 5 * 10^4

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[string]int)
	tMap := make(map[string]int)
	
	sSplit := strings.Split(s, "")
	tSplit := strings.Split(t, "")
	for _, sc := range sSplit {
		sMap[sc] = sMap[sc] + 1
	}
	for _, tc := range tSplit {
		if sMap[tc] == 0 {
			return false
		}
		tMap[tc] = tMap[tc] + 1
	}
	for _, tc := range tSplit {
		if sMap[tc] != tMap[tc] {
			return false
		}
	}
	return true
}
