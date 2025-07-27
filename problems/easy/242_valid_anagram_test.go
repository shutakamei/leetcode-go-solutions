package main

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
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
			name: "One empty string",
			s:    "hello",
			t:    "",
			want: false,
		},
		{
			name: "Single character",
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			name: "Single character different",
			s:    "a",
			t:    "b",
			want: false,
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
		{
			name: "Not anagram with repeated characters",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "Long strings - anagram",
			s:    "abcdefghijklmnopqrstuvwxyz",
			t:    "zyxwvutsrqponmlkjihgfedcba",
			want: true,
		},
		{
			name: "Long strings - not anagram",
			s:    "abcdefghijklmnopqrstuvwxyz",
			t:    "zyxwvutsrqponmlkjihgfedcbaa",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

// ベンチマークテスト
func BenchmarkIsAnagram(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram(s, t)
	}
}

func BenchmarkIsAnagramNotAnagram(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcbaa"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram(s, t)
	}
}

func BenchmarkIsAnagramShort(b *testing.B) {
	s := "hello"
	t := "olleh"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram(s, t)
	}
} 
