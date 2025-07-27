# LeetCode Solutions in Go

このプロジェクトは、LeetCode の問題を Go 言語で解くためのリポジトリです。

## プロジェクト構造

```
leetcode/
├── problems/          # 問題別のソリューション
│   ├── easy/         # Easy難易度の問題
│   ├── medium/       # Medium難易度の問題
│   └── hard/         # Hard難易度の問題
├── utils/            # 共通ユーティリティ
├── tests/            # テストファイル
└── main.go           # メインファイル
```

## 使用方法

### 新しい問題を追加する場合

1. 適切な難易度フォルダに新しいファイルを作成
2. ファイル名は `problem_number_problem_name.go` の形式
3. テストファイルも同時に作成

### 例

```go
// problems/easy/1_two_sum.go
package main

func twoSum(nums []int, target int) []int {
    // 実装
}
```

## テストの実行

```bash
go test ./...
```

## 特定の問題のテスト実行

```bash
go test ./problems/easy -v
```

## 便利なコマンド

- `make test`: 全テストの実行
- `make new-problem`: 新しい問題ファイルの作成
- `make run`: メインアプリケーションの実行
