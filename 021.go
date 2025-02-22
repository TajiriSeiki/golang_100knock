package main

import "fmt"

func main() {
	var num1 int

	fmt.Print("整数を入力してください: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	if (num1 > 5 && num1 < 20) {
		fmt.Println("OK")
	}
}