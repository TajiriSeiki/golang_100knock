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
	for i := 1; i <  9 + 1; i++ {
		if i == inputNum || i == inputNum + 1 {
			continue
		}
		fmt.Println(i)
	}
	
}