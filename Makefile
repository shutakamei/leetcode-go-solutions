.PHONY: test test-easy test-medium test-hard run clean new-problem help

# デフォルトターゲット
help:
	@echo "🚀 LeetCode Go Solutions - 利用可能なコマンド"
	@echo "=========================================="
	@echo "make test        - 全テストを実行"
	@echo "make test-easy   - Easy問題のテストを実行"
	@echo "make test-medium - Medium問題のテストを実行"
	@echo "make test-hard   - Hard問題のテストを実行"
	@echo "make run         - メインアプリケーションを実行"
	@echo "make clean       - ビルドファイルをクリーンアップ"
	@echo "make new-problem - 新しい問題ファイルを作成"
	@echo "make help        - このヘルプを表示"

# 全テストを実行
test:
	@echo "🧪 全テストを実行中..."
	go test ./... -v

# Easy問題のテストを実行
test-easy:
	@echo "🧪 Easy問題のテストを実行中..."
	go test ./problems/easy -v

# Medium問題のテストを実行
test-medium:
	@echo "🧪 Medium問題のテストを実行中..."
	go test ./problems/medium -v

# Hard問題のテストを実行
test-hard:
	@echo "🧪 Hard問題のテストを実行中..."
	go test ./problems/hard -v

# メインアプリケーションを実行
run:
	@echo "🚀 メインアプリケーションを実行中..."
	go run main.go

# ビルドファイルをクリーンアップ
clean:
	@echo "🧹 ビルドファイルをクリーンアップ中..."
	go clean
	rm -f leetcode

# 新しい問題ファイルを作成
new-problem:
	@echo "📝 新しい問題ファイルを作成します"
	@read -p "問題番号を入力してください: " number; \
	read -p "問題名を入力してください (例: two_sum): " name; \
	read -p "難易度を選択してください (easy/medium/hard): " difficulty; \
	file_path="problems/$$difficulty/$${number}_$${name}.go"; \
	test_file_path="problems/$$difficulty/$${number}_$${name}_test.go"; \
	echo "package main" > $$file_path; \
	echo "" >> $$file_path; \
	echo "// Problem $$number - $$name" >> $$file_path; \
	echo "// https://leetcode.com/problems/$$name/" >> $$file_path; \
	echo "" >> $$file_path; \
	echo "func solution$$number() {" >> $$file_path; \
	echo "	// TODO: 実装してください" >> $$file_path; \
	echo "}" >> $$file_path; \
	echo "package main" > $$test_file_path; \
	echo "" >> $$test_file_path; \
	echo "import \"testing\"" >> $$test_file_path; \
	echo "" >> $$test_file_path; \
	echo "func TestProblem$$number(t *testing.T) {" >> $$test_file_path; \
	echo "	// TODO: テストケースを追加してください" >> $$test_file_path; \
	echo "}" >> $$test_file_path; \
	echo "✅ ファイルを作成しました:"; \
	echo "   - $$file_path"; \
	echo "   - $$test_file_path"

# ベンチマークテストを実行
bench:
	@echo "⚡ ベンチマークテストを実行中..."
	go test ./... -bench=.

# コードフォーマット
fmt:
	@echo "🎨 コードをフォーマット中..."
	go fmt ./...

# 静的解析
vet:
	@echo "🔍 静的解析を実行中..."
	go vet ./...

# 依存関係を整理
mod-tidy:
	@echo "📦 依存関係を整理中..."
	go mod tidy 
