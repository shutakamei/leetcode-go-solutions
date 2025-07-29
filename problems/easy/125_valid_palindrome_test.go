package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1 - Valid palindrome",
			s:    "Was it a car or a cat I saw?",
			want: true,
		},
		{
			name: "Example 2 - Not palindrome",
			s:    "tab a cat",
			want: false,
		},
		{
			name: "Simple palindrome",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Empty string",
			s:    "",
			want: true,
		},
		{
			name: "Single character",
			s:    "a",
			want: true,
		},
		{
			name: "Single non-alphanumeric",
			s:    "!",
			want: true,
		},
		{
			name: "Only non-alphanumeric",
			s:    "!@#$%^&*()",
			want: true,
		},
		{
			name: "Mixed alphanumeric and symbols",
			s:    "race a car",
			want: false,
		},
		{
			name: "Numbers only",
			s:    "12321",
			want: true,
		},
		{
			name: "Numbers not palindrome",
			s:    "12345",
			want: false,
		},
		{
			name: "Mixed case palindrome",
			s:    "A1b2C3c2b1a",
			want: true,
		},
		{
			name: "Mixed case not palindrome",
			s:    "A1b2C3d2b1a",
			want: false,
		},
		{
			name: "With spaces and punctuation",
			s:    "Never odd or even",
			want: true,
		},
		{
			name: "With spaces and punctuation not palindrome",
			s:    "Never even or odd",
			want: false,
		},
		{
			name: "Single letter with symbols",
			s:    "a.",
			want: true,
		},
		{
			name: "Two different letters",
			s:    "ab",
			want: false,
		},
		{
			name: "Two same letters",
			s:    "aa",
			want: true,
		},
		{
			name: "Long palindrome",
			s:    "Do geese see God?",
			want: true,
		},
		{
			name: "Long not palindrome",
			s:    "Do geese see dogs?",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.s)
			if got != tt.want {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

// ベンチマークテスト
func BenchmarkIsPalindrome(b *testing.B) {
	s := "A man, a plan, a canal: Panama"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isPalindrome(s)
	}
}

func BenchmarkIsPalindromeNot(b *testing.B) {
	s := "This is not a palindrome"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isPalindrome(s)
	}
}

func BenchmarkIsPalindromeLong(b *testing.B) {
	s := "Do geese see God? This is a very long string to test performance of the palindrome function with many characters and symbols!"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isPalindrome(s)
	}
} 
