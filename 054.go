package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 標準入力を読み取るスキャナ
	scanner := bufio.NewScanner(os.Stdin)

	// 1行目を読み取り、データ数を取得
	scanner.Scan()
	numData, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("入力エラー")
		return
	}

	data := make([]int, numData)

	for i := 0; i < numData; i++ {
		if !scanner.Scan() {
			fmt.Println("データの読み取りエラー")
			return
		}
		data[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("無効な整数が入力されました")
			return
		}
	}

	max := data[1]
	min := data[1]
	for j:=1; j < len(data); j++ {
		if (max < data[j]) {
			max = data[j]
		}
		if (min > data[j]) {
			min = data[j]
		}
	}
	fmt.Printf("最小値: %d, 最大値: %d", min, max)
}