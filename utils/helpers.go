package utils

import (
	"fmt"
	"time"
)

// 実行時間を計測するヘルパー関数
func MeasureExecutionTime(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

// 配列をプリントするヘルパー関数
func PrintArray(arr []int, name string) {
	fmt.Printf("%s: [", name)
	for i, val := range arr {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d", val)
	}
	fmt.Println("]")
}

// 2次元配列をプリントするヘルパー関数
func Print2DArray(arr [][]int, name string) {
	fmt.Printf("%s:\n", name)
	for i, row := range arr {
		fmt.Printf("  [%d]: [", i)
		for j, val := range row {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d", val)
		}
		fmt.Println("]")
	}
}

// 文字列配列をプリントするヘルパー関数
func PrintStringArray(arr []string, name string) {
	fmt.Printf("%s: [", name)
	for i, val := range arr {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("\"%s\"", val)
	}
	fmt.Println("]")
}

// 結果を比較するヘルパー関数
func CompareResults(got, want interface{}, testName string) bool {
	equal := fmt.Sprintf("%v", got) == fmt.Sprintf("%v", want)
	
	fmt.Printf("テスト: %s\n", testName)
	fmt.Printf("  結果: %v\n", got)
	fmt.Printf("  期待: %v\n", want)
	fmt.Printf("  正解: %t\n\n", equal)
	
	return equal
}

// 実行時間を表示するヘルパー関数
func PrintExecutionTime(duration time.Duration, functionName string) {
	fmt.Printf("⏱️  %s の実行時間: %v\n", functionName, duration)
} 
