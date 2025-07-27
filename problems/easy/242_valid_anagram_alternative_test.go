package main

import (
	"testing"
)

// 全アプローチのテスト
func TestAllAnagramApproaches(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "Example 1 - Valid anagram",
			s:    "racecar",
			t:    "carrace",
			want: true,
		},
		{
			name: "Example 2 - Not anagram",
			s:    "jar",
			t:    "jam",
			want: false,
		},
		{
			name: "Same string",
			s:    "hello",
			t:    "hello",
			want: true,
		},
		{
			name: "Different lengths",
			s:    "hello",
			t:    "world",
			want: false,
		},
		{
			name: "Empty strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "Classic anagram",
			s:    "listen",
			t:    "silent",
			want: true,
		},
		{
			name: "Anagram with repeated characters",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
	}

	approaches := []struct {
		name string
		fn   func(string, string) bool
	}{
		{"Single Map", isAnagram1},
		{"Array Method", isAnagram2},
		{"Sort Method", isAnagram3},
		{"Double Map", isAnagram4},
		{"Bitmask", isAnagram5},
		{"String Split", isAnagram6},
	}

	for _, approach := range approaches {
		t.Run(approach.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					got := approach.fn(tt.s, tt.t)
					if got != tt.want {
						t.Errorf("%s: %s(%q, %q) = %v, want %v", 
							approach.name, approach.name, tt.s, tt.t, got, tt.want)
					}
				})
			}
		})
	}
}

// 各アプローチのベンチマークテスト
func BenchmarkAnagram1(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram1(s, t)
	}
}

func BenchmarkAnagram2(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram2(s, t)
	}
}

func BenchmarkAnagram3(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram3(s, t)
	}
}

func BenchmarkAnagram4(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram4(s, t)
	}
}

func BenchmarkAnagram5(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram5(s, t)
	}
}

func BenchmarkAnagram6(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram6(s, t)
	}
}

// 短い文字列でのベンチマーク
func BenchmarkAnagramShort1(b *testing.B) {
	s := "hello"
	t := "olleh"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram1(s, t)
	}
}

func BenchmarkAnagramShort2(b *testing.B) {
	s := "hello"
	t := "olleh"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram2(s, t)
	}
}

func BenchmarkAnagramShort3(b *testing.B) {
	s := "hello"
	t := "olleh"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram3(s, t)
	}
} 
