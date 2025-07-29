package main

import (
	"testing"
)

func TestIsValidSolution(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1 - Simple brackets",
			s:    "[]",
			want: true,
		},
		{
			name: "Example 2 - Nested brackets",
			s:    "([{}])",
			want: true,
		},
		{
			name: "Example 3 - Invalid order",
			s:    "[(])",
			want: false,
		},
		{
			name: "Empty string",
			s:    "",
			want: true,
		},
		{
			name: "Single opening bracket",
			s:    "(",
			want: false,
		},
		{
			name: "Single closing bracket",
			s:    ")",
			want: false,
		},
		{
			name: "Simple parentheses",
			s:    "()",
			want: true,
		},
		{
			name: "Simple braces",
			s:    "{}",
			want: true,
		},
		{
			name: "Simple brackets",
			s:    "[]",
			want: true,
		},
		{
			name: "Mixed valid brackets",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "Nested valid brackets",
			s:    "({[]})",
			want: true,
		},
		{
			name: "Multiple nested",
			s:    "((()))",
			want: true,
		},
		{
			name: "Multiple nested mixed",
			s:    "({[()]})",
			want: true,
		},
		{
			name: "Invalid - wrong closing",
			s:    "(]",
			want: false,
		},
		{
			name: "Invalid - wrong closing 2",
			s:    "([)]",
			want: false,
		},
		{
			name: "Invalid - extra closing",
			s:    "())",
			want: false,
		},
		{
			name: "Invalid - extra opening",
			s:    "(()",
			want: false,
		},
		{
			name: "Invalid - mixed wrong order",
			s:    "{[}]",
			want: false,
		},
		{
			name: "Long valid string",
			s:    "(((((())))))",
			want: true,
		},
		{
			name: "Long invalid string",
			s:    "(((((()))))",
			want: false,
		},
		{
			name: "Complex valid",
			s:    "({[]})[()]{()}",
			want: true,
		},
		{
			name: "Complex invalid",
			s:    "({[]})[()]{()",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSolution(tt.s); got != tt.want {
				t.Errorf("isValidSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 全実装の比較テスト
func TestAllImplementations(t *testing.T) {
	testCases := []string{
		"",                    // 空文字列
		"()",                  // 短い有効
		"()[]{}",             // 短い混合有効
		"(((((())))))",       // 長い有効
		"({[()]})",           // 複雑な有効
		"({[()]})[()]{()}",   // 非常に複雑な有効
		"(",                   // 単一開き括弧
		")",                   // 単一閉じ括弧
		"(]",                  // 短い無効
		"([)]",               // 短い混合無効
		"(((((()))))",        // 長い無効
		"({[()]})[()]{()",    // 複雑な無効
	}

	for _, s := range testCases {
		t.Run(s, func(t *testing.T) {
			result1 := isValidSolution(s)
			result2 := isValidOptimized(s)
			result3 := isValidSimple(s)
			
			if result1 != result2 || result2 != result3 {
				t.Errorf("Different results for '%s': Solution=%v, Optimized=%v, Simple=%v", 
					s, result1, result2, result3)
			}
		})
	}
}

// ベンチマークテスト
func BenchmarkIsValidSolution(b *testing.B) {
	testCases := []string{
		"",                    // 空文字列
		"()",                  // 短い有効
		"()[]{}",             // 短い混合有効
		"(((((())))))",       // 長い有効
		"({[()]})",           // 複雑な有効
		"({[()]})[()]{()}",   // 非常に複雑な有効
	}

	for _, s := range testCases {
		b.Run("Solution_"+s, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				isValidSolution(s)
			}
		})
		
		b.Run("Optimized_"+s, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				isValidOptimized(s)
			}
		})
		
		b.Run("Simple_"+s, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				isValidSimple(s)
			}
		})
	}
} 
