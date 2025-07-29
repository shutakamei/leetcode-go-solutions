package main

// Valid Parentheses - 正解実装
// スタックを使った効率的な解法

func isValidSolution(s string) bool {
	// スタックを使って括弧の対応関係を管理
	var stack []rune
	
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			// 開き括弧はスタックにプッシュ
			stack = append(stack, char)
		case ')', '}', ']':
			// 閉じ括弧の場合
			if len(stack) == 0 {
				return false // 対応する開き括弧がない
			}
			
			// スタックの一番上を取り出して対応関係をチェック
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // ポップ
			
			// 対応関係のチェック
			if !isMatching(top, char) {
				return false
			}
		}
	}
	
	// スタックが空なら有効
	return len(stack) == 0
}

// 開き括弧と閉じ括弧の対応関係をチェック
func isMatching(open, close rune) bool {
	switch open {
	case '(':
		return close == ')'
	case '{':
		return close == '}'
	case '[':
		return close == ']'
	default:
		return false
	}
}

// より効率的な実装（マップを使用）
func isValidOptimized(s string) bool {
	// 括弧の対応関係をマップで管理
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	
	var stack []rune
	
	for _, char := range s {
		// 開き括弧の場合
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			// 閉じ括弧の場合
			if len(stack) == 0 {
				return false
			}
			
			// スタックの一番上を取り出して対応関係をチェック
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			
			if brackets[char] != top {
				return false
			}
		}
	}
	
	return len(stack) == 0
}

// さらに簡潔な実装
func isValidSimple(s string) bool {
	var stack []rune
	
	for _, char := range s {
		switch char {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	
	return len(stack) == 0
} 
