package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("🚀 LeetCode Solutions in Go")
	fmt.Println("================================")
	
	if len(os.Args) < 2 {
		fmt.Println("使用方法:")
		fmt.Println("  go run main.go [problem_number]")
		fmt.Println("例:")
		fmt.Println("  go run main.go 1")
		return
	}
	
	problemNumber := os.Args[1]
	fmt.Printf("問題 %s を実行中...\n", problemNumber)
	
	// ここで問題番号に応じた関数を呼び出す
	// 例: runProblem1(), runProblem2() など
	fmt.Printf("問題 %s の実装は problems/ ディレクトリ内にあります\n", problemNumber)
} 
