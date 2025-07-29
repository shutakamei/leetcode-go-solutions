package main

import "strings"

// Valid Palindrome - LeetCode Problem 125
// https://leetcode.com/problems/valid-palindrome/
//
// 問題: 文字列 s が与えられたとき、
// それが回文である場合は true を返し、
// そうでなければ false を返す。
//
// 回文とは、前から読んでも後ろから読んでも同じ文字列のこと。
// 大文字小文字は区別せず、英数字以外の文字は無視する。
//
// 注意: 英数字は文字（A-Z, a-z）と数字（0-9）で構成される。
//
// 制約:
// - 1 <= s.length <= 1000
// - s は印刷可能なASCII文字のみで構成される

func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    
	var charArr []rune
    for _, char := range s {
        if (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') {
			charArr = append(charArr, char)
        }
    }
	sLen := len(charArr)
	if sLen == 0 {
		return true
	}
	sLenHalf := sLen / 2
	for i := 0; i <= sLenHalf; i++ {
		j := sLen - i - 1
		if charArr[i] == charArr[j] {
			continue
		}
		return false
	}
	return true
}
