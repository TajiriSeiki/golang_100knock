package main

import "fmt"

func main() {
	var money int

	fmt.Print("input money: ")
	_, err := fmt.Scanln(&money)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	num100 := money / 100
	num10 := money / 10 - 10*num100
	num1 := money % 10

	fmt.Printf("100円玉: %d, 10円玉: %d, 1円玉: %d", num100, num10, num1)
}