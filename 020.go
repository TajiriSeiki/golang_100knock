package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Print("1つ目の整数を入力してください: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fmt.Print("2つ目の整数を入力してください: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fmt.Println("result: ", num1 / num2)
	fmt.Println("result: ", (num1 / num2) * num2)
}