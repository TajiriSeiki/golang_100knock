package main

import "fmt"

func main() {
	var inputNum int
	sum := 0
	fmt.Print("整数を入力してください: ")
	_, err := fmt.Scanln(&inputNum)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	for i := 1; i <=  inputNum; i++ {
		sum += i
	}
	fmt.Println(sum)
}