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

	if (num > 0) {
		fmt.Println("positive")
	} else if (num == 0) {
		fmt.Println("zero")
	} else {
		fmt.Println("negative")
	}
}