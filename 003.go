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

	fmt.Println("入力された値は:", num)
}