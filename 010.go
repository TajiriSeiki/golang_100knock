package main

import "fmt"

func main() {
    var num int

	fmt.Print("整数を入力してください: ")
	_, err := fmt.Scanln(&num)
    if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	// math.Abs を使わずに実装する
	if (num < 0) {
		num = num * (-1)
	}
	fmt.Println("absolute value:", num)
}