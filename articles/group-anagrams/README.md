# 【LeetCode】Group Anagrams (Medium) - アナグラムのグループ化問題を Go で解く

## 目次

1. [問題の概要](#問題の概要)
2. [制約条件](#制約条件)
3. [例](#例)
4. [実装](#実装)
5. [解法のアプローチ](#解法のアプローチ)
6. [性能比較とベンチマーク結果](#性能比較とベンチマーク結果)
7. [詳細な性能分析](#詳細な性能分析)
8. [メモリ使用量の分析](#メモリ使用量の分析)
9. [実装選択のガイドライン](#実装選択のガイドライン)
10. [まとめ](#まとめ)

## 問題の概要

文字列の配列 `strs` が与えられ、すべてのアナグラムをサブリストにグループ化する問題です。

**アナグラム**とは、別の文字列と同じ文字を正確に含む文字列ですが、文字の順序が異なる場合があります。

## 制約条件

- `1 <= strs.length <= 1000`
- `0 <= strs[i].length <= 100`
- `strs[i]` は小文字の英字で構成されています

## 例

### 例 1

```
入力: strs = ["act","pots","tops","cat","stop","hat"]
出力: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
```

### 例 2

```
入力: strs = ["x"]
出力: [["x"]]
```

### 例 3

```
入力: strs = [""]
出力: [[""]]
```

## 実装

### 1. 私が実装した解法（文字カウント + マップ比較）

#### アプローチ

Valid Anagram 問題で効率的に動作したマップを使った処理を応用しようと考えましたが、
アナグラムが一致した際に同じ配列に追加する方法が思いつかず、結果的に強引な実装になってしまいました。

- 各文字列の文字カウントをマップで管理
- 既存のマップと比較してグループ化
- **時間計算量**: O(n² × k)
- **空間計算量**: O(n × k)

#### 実装

```go
func groupAnagrams(strs []string) [][]string {
    if len(strs) <= 1 {
        if len(strs) == 0 {
            return [][]string{}
        }
        return [][]string{strs}
    }

    var anagramArr [][]string
    var charMapArr []map[rune]int

    for _, str := range strs {
        charMap := make(map[rune]int)
        for _, char := range str {
            charMap[char]++
        }

        newCharFlag := true
        for i, cm := range charMapArr {
            if reflect.DeepEqual(cm, charMap) {
                anagramArr[i] = append(anagramArr[i], str)
                newCharFlag = false
                break
            }
        }

        if newCharFlag {
            charMapArr = append(charMapArr, charMap)
            anagramArr = append(anagramArr, []string{str})
        }
    }

    return anagramArr
}
```

### 2. ソート方式

#### アプローチ

- 各文字列をソートして正規形を作成
- ソートされた文字列をキーとしてマップでグループ化
- **時間計算量**: O(n × k × log k)
- **空間計算量**: O(n × k)

#### 実装

```go
func groupAnagramsSorted(strs []string) [][]string {
    if len(strs) <= 1 {
        if len(strs) == 0 {
            return [][]string{}
        }
        return [][]string{strs}
    }

    groups := make(map[string][]string)

    for _, str := range strs {
        canonical := getSortedString(str)
        groups[canonical] = append(groups[canonical], str)
    }

    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }

    return result
}

func getSortedString(s string) string {
    chars := []rune(s)
    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })
    return string(chars)
}
```

### 3. 文字カウント文字列化方式（推奨）

#### アプローチ

- 各文字の出現回数をカウント
- カウント結果を文字列として表現してキー化
- **時間計算量**: O(n × k)
- **空間計算量**: O(n × k)

#### 実装

```go
func groupAnagramsCountString(strs []string) [][]string {
    if len(strs) <= 1 {
        if len(strs) == 0 {
            return [][]string{}
        }
        return [][]string{strs}
    }

    groups := make(map[string][]string)

    for _, str := range strs {
        canonical := getCountString(str)
        groups[canonical] = append(groups[canonical], str)
    }

    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }

    return result
}

func getCountString(s string) string {
    count := make([]int, 26)

    for _, char := range s {
        count[char-'a']++
    }

    var result strings.Builder
    for i := 0; i < 26; i++ {
        if count[i] > 0 {
            result.WriteByte(byte('a' + i))
            result.WriteString(string(rune(count[i] + '0')))
        }
    }

    return result.String()
}
```

### 4. 文字カウント配列方式

#### アプローチ

- 26 文字の配列で各文字の出現回数を管理
- 配列を文字列として表現してキー化
- **時間計算量**: O(n × k)
- **空間計算量**: O(n × k)

#### 実装

```go
func groupAnagramsCountArray(strs []string) [][]string {
    if len(strs) <= 1 {
        if len(strs) == 0 {
            return [][]string{}
        }
        return [][]string{strs}
    }

    groups := make(map[string][]string)

    for _, str := range strs {
        canonical := getCountArrayString(str)
        groups[canonical] = append(groups[canonical], str)
    }

    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }

    return result
}

func getCountArrayString(s string) string {
    count := make([]int, 26)

    for _, char := range s {
        count[char-'a']++
    }

    var result strings.Builder
    for i := 0; i < 26; i++ {
        if i > 0 {
            result.WriteByte('#')
        }
        result.WriteString(string(rune(count[i] + '0')))
    }

    return result.String()
}
```

## 性能比較とベンチマーク結果

### 実行環境

- **OS**: Linux 6.6.87.2-microsoft-standard-WSL2
- **CPU**: 12th Gen Intel(R) Core(TM) i7-12700H
- **Go Version**: 1.21+
- **アーキテクチャ**: amd64

### ベンチマーク結果

#### 短い文字列（12 要素）での性能比較

| 実装方法                 | 実行回数 | 平均時間 (ns/op) | 性能向上率     | 評価       |
| ------------------------ | -------- | ---------------- | -------------- | ---------- |
| **私の実装**             | 63,482   | 17,246           | 基準           | ⭐         |
| **ソート方式**           | 325,282  | 3,542            | **4.9 倍高速** | ⭐⭐⭐     |
| **文字カウント文字列化** | 648,145  | 1,830            | **9.4 倍高速** | ⭐⭐⭐⭐⭐ |
| **文字カウント配列**     | 206,853  | 5,066            | **3.4 倍高速** | ⭐⭐⭐⭐   |

#### 長い文字列（26 文字 × 6 要素）での性能比較

| 実装方法                 | 実行回数 | 平均時間 (ns/op) | 性能向上率     | 評価       |
| ------------------------ | -------- | ---------------- | -------------- | ---------- |
| **ソート方式**           | 304,905  | 3,889            | 基準           | ⭐⭐⭐     |
| **文字カウント文字列化** | 493,472  | 2,146            | **1.8 倍高速** | ⭐⭐⭐⭐⭐ |
| **文字カウント配列**     | 439,186  | 2,707            | **1.4 倍高速** | ⭐⭐⭐⭐   |

### 生のベンチマーク出力

```
BenchmarkGroupAnagrams-20                          63482             17246 ns/op
BenchmarkGroupAnagramsSorted-20                   325282              3542 ns/op
BenchmarkGroupAnagramsCountString-20              648145              1830 ns/op
BenchmarkGroupAnagramsCountArray-20               206853              5066 ns/op
BenchmarkGroupAnagramsLongStrings/Sorted-20       304905              3889 ns/op
BenchmarkGroupAnagramsLongStrings/CountString-20  493472              2146 ns/op
BenchmarkGroupAnagramsLongStrings/CountArray-20   439186              2707 ns/op
```

## 詳細な性能分析

### 1. 私が実装した解法（文字カウント + マップ比較）

- **時間計算量**: O(n² × k)
- **特徴**:
  - 各文字列を既存のマップと比較するため、二重ループが必要
  - `reflect.DeepEqual`によるマップ比較が重い
  - Valid Anagram 問題の解法を応用しようとしたが、グループ化の処理が非効率
- **性能**: 最も遅いが、理解しやすい実装
- **ベンチマーク**: 17,246 ns/op（基準）

### 2. ソート方式

- **時間計算量**: O(n × k × log k)
- **特徴**:
  - 文字列のソートがボトルネック
  - 短い文字列では効率的
  - 長い文字列ではソートのオーバーヘッドが大きい
- **性能**: 短い文字列で良好、長い文字列では劣化
- **ベンチマーク**:
  - 短い文字列: 3,542 ns/op（4.9 倍高速）
  - 長い文字列: 3,889 ns/op（基準）

### 3. 文字カウント文字列化方式（推奨）

- **時間計算量**: O(n × k)
- **特徴**:
  - ソート不要
  - 文字カウントを効率的に文字列化
  - ハッシュマップのキーとして最適
  - `strings.Builder`による効率的な文字列構築
- **性能**: すべてのケースで最高性能
- **ベンチマーク**:
  - 短い文字列: 1,830 ns/op（9.4 倍高速）
  - 長い文字列: 2,146 ns/op（1.8 倍高速）

### 4. 文字カウント配列方式

- **時間計算量**: O(n × k)
- **特徴**:
  - 26 文字の配列を文字列化
  - 短い文字列ではオーバーヘッドが大きい
  - 長い文字列では効率的
  - 固定長の配列を使用
- **性能**: 長い文字列で良好、短い文字列では劣る
- **ベンチマーク**:
  - 短い文字列: 5,066 ns/op（3.4 倍高速）
  - 長い文字列: 2,707 ns/op（1.4 倍高速）

## メモリ使用量の分析

### 空間計算量の比較

| 実装方法                 | 空間計算量 | メモリ効率 | 説明                         |
| ------------------------ | ---------- | ---------- | ---------------------------- |
| **私の実装**             | O(n × k)   | 低         | 各文字列のマップを保持       |
| **ソート方式**           | O(n × k)   | 中         | ソートされた文字列を保持     |
| **文字カウント文字列化** | O(n × k)   | 高         | カウント文字列を保持（短い） |
| **文字カウント配列**     | O(n × k)   | 中         | 配列文字列を保持（固定長）   |

### 実際のメモリ効率ランキング

1. **文字カウント文字列化**: 最も効率的

   - 短い文字列ではカウント文字列が短い
   - 例: "act" → "a1c1t1"（6 文字）

2. **ソート方式**: 中程度

   - ソートされた文字列の長さは元と同じ
   - 例: "act" → "act"（3 文字）

3. **文字カウント配列**: 中程度

   - 26 文字の配列を文字列化
   - 例: "act" → "1#0#1#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#0#1"（51 文字）

4. **私の実装**: 最も非効率
   - 各文字列のマップを保持
   - 例: "act" → map[rune]int{'a':1, 'c':1, 't':1}

## 実装選択のガイドライン

**文字カウント文字列化方式**を推奨します：

✅ **最高性能**: 元の実装の 9.4 倍高速  
✅ **一貫性**: 短い文字列でも長い文字列でも良好な性能  
✅ **可読性**: 実装が理解しやすい  
✅ **メモリ効率**: 適切なメモリ使用量

### 使用場面別の選択

| 使用場面           | 推奨実装             | 理由               | 性能向上率　(私の実装比) |
| ------------------ | -------------------- | ------------------ | ------------------------ |
| **一般的な用途**   | 文字カウント文字列化 | 最高性能、一貫性   | 9.4 倍                   |
| **短い文字列中心** | ソート方式           | シンプル、効率的   | 4.9 倍                   |
| **長い文字列中心** | 文字カウント配列     | 長い文字列で効率的 | 1.4 倍                   |

### パフォーマンス要件別の選択

#### 高パフォーマンス要件

- **推奨**: 文字カウント文字列化方式
- **理由**: 最高性能、一貫性
- **期待性能**: 9.4 倍の高速化

#### バランス重視

- **推奨**: ソート方式
- **理由**: シンプル、十分な性能
- **期待性能**: 4.9 倍の高速化

#### メモリ効率重視

- **推奨**: 文字カウント文字列化方式
- **理由**: 最も効率的なメモリ使用
- **期待性能**: 9.4 倍の高速化

## まとめ

Group Anagrams 問題は、文字列の正規化とハッシュマップを組み合わせることで効率的に解決できます。

### 🏆 **最優秀解法**: 文字カウント文字列化方式

- **一貫性**: 短い文字列でも長い文字列でも良好な性能
- **実装**: 理解しやすく、メンテナンスしやすい
- **メモリ**: 効率的なメモリ使用

### 📊 **性能向上の実証**

- 実際のベンチマークで私の実装の 9.4 倍の性能向上を実現
- 様々な入力サイズでの一貫した性能
- メモリ効率も最適化

### 🎯 **実用的な価値**

- LeetCode の制約内で十分に効率的
- 本番環境でも使用可能な品質
- 学習から実用まで幅広く対応

この解法により、様々な入力サイズに対応できる汎用的なソリューションを提供できます。

## 参考リンク

- [LeetCode - Group Anagrams](https://leetcode.com/problems/group-anagrams/)
- [Go 言語公式ドキュメント](https://golang.org/)
- [Go 言語の strings.Builder](https://golang.org/pkg/strings/#Builder)
- [Go 言語のベンチマーク](https://golang.org/pkg/testing/#hdr-Benchmarks)

---

**タグ**: #Go #LeetCode #アルゴリズム #文字列処理 #ハッシュマップ #パフォーマンス最適化 #ベンチマーク #性能分析
