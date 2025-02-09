package main

import (
    "fmt"
)

func main() {
    var num1 int
    var num2 int

	fmt.Print("1つ目の整数を入力してください: ")
	_, err := fmt.Scanln(&num1)
	fmt.Print("2つ目の整数を入力してください: ")
	_, err = fmt.Scanln(&num2)
    if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}

	fmt.Println("和：", num1 + num2)
	fmt.Println("差：", num1 - num2)
	fmt.Println("積：", num1 * num2)
	fmt.Printf("商：%d, 余り：%d", num1/num2 , num1 % num2)

}