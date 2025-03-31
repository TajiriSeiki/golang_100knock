package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("標準入力の読み取りに失敗しました:", err)
		return
	}

	numStudents := len(lines) - 1
	if numStudents > 100 {
		fmt.Println("受験者数が100人を超えています。")
		return
	}

	var englishTotal, mathTotal, japaneseTotal int
	scores := make([]int, numStudents)

	for i, line := range lines[1:] {
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) != 3 {
			fmt.Println("データ形式が不正です:", line)
			return
		}

		english, _ := strconv.Atoi(fields[0])
		math, _ := strconv.Atoi(fields[1])
		japanese, _ := strconv.Atoi(fields[2])

		englishTotal += english
		mathTotal += math
		japaneseTotal += japanese
		scores[i] = english + math + japanese
	}

	// 各科目の平均点を計算
	englishAverage := float64(englishTotal) / float64(numStudents)
	mathAverage := float64(mathTotal) / float64(numStudents)
	japaneseAverage := float64(japaneseTotal) / float64(numStudents)

	fmt.Printf("英語の平均点: %.2f\n", englishAverage)
	fmt.Printf("数学の平均点: %.2f\n", mathAverage)
	fmt.Printf("国語の平均点: %.2f\n", japaneseAverage)

	// 各受験生の合計点を表示
	fmt.Println("各受験生の合計点:")
	for i, total := range scores {
		fmt.Printf("%d: %d\n", i, total)
	}
}
