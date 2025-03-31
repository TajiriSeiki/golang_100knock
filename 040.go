package main

import "fmt"

func main() {
	var inputNum int
	fmt.Print("整数を入力してください: ")
	_, err := fmt.Scanln(&inputNum)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	if inputNum % 2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}