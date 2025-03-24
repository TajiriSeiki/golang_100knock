package main

import "fmt"

func main() {
	var inputNumA int
	var inputNumB int
	var inputNumC int

	fmt.Print("整数を入力してください a: ")
	_, err := fmt.Scanln(&inputNumA)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fmt.Print("整数を入力してください b: ")
	_, err = fmt.Scanln(&inputNumB)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fmt.Print("整数を入力してください c: ")
	_, err = fmt.Scanln(&inputNumC)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}

	condition := inputNumB * inputNumB - 4 * inputNumA * inputNumC 
	if condition > 0 {
		fmt.Println("2つの実数解")
	} else if condition < 0 {
		fmt.Println("2つの虚数")
	} else {
		fmt.Println("重解")
	}
}