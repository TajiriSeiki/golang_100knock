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
	
	if ((num1 >= -10) && (num1 < 0)) || (num1 >= 10) {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}