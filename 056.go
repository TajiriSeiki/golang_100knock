package main

import "fmt"

func main() {
	var num int

	fmt.Print("0〜65535の整数を入力してください: ")
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("入力エラー: 整数を入力してください")
		return
	}

	if num < 0 || num > 65535 {
		fmt.Println("エラー: 0〜65535の範囲で入力してください")
		return
	}

	binaryStr := ""
	for i := 15; i >= 0; i-- {
		if (num & (1 << i)) != 0 {
			binaryStr += "1"
		} else {
			binaryStr += "0"
		}
	}

	fmt.Println(binaryStr)
}