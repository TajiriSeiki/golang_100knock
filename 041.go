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

	if inputNum >= 10 || inputNum <= 0 {
		fmt.Printf("%d is not a single figure.", inputNum)
	} else {
		fmt.Printf("%d is a single figure.", inputNum)
	}
}