package main

// Valid Parentheses - LeetCode Problem 20
// https://leetcode.com/problems/valid-parentheses/
//
// 問題: 文字列 s が与えられたとき、
// それが有効な括弧の組み合わせである場合は true を返し、
// そうでなければ false を返す。
//
// 有効な括弧の組み合わせとは：
// - すべての開き括弧が同じ種類の閉じ括弧で閉じられる
// - 開き括弧が正しい順序で閉じられる
// - すべての閉じ括弧に対応する同じ種類の開き括弧がある
//
// 括弧の種類: '(', ')', '{', '}', '[', ']'
//
// 制約:
// - 1 <= s.length <= 1000
// - s は '(', ')', '{', '}', '[', ']' のみで構成される

func isValid(s string) bool {
	// if len(s) % 2 == 1 {
	// 	return false
	// }
	// var charArr []rune
	// for _, char := range s {
	// 	switch char {
	// 	case "(", "{", "[":
	// 			charArr = append(charArr, char)
	// 	case ")":
	// 			for i := len(charArr) - 1; 0 >= i; i-- {
	// 				switch char {
	// 				case "{", "[":
	// 						return false
	// 				}
	// 			}
	// 			charArr = append(charArr, char)
	// 		case "}":
	// 			charArr = append(charArr, char)
	// 		case "]":
	// 			charArr = append(charArr, char)
	// 	}
	// }
	return true
}
