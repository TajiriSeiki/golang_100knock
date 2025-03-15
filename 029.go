package main

import "fmt"

func main() {
	var inputNum int
	sum := 0

	for i := 0; i <  5; i++  {
		fmt.Print("整数を入力してください: ")
		_, err := fmt.Scanln(&inputNum)
		if err != nil {
			fmt.Println("エラーが発生しました:", err)
			return
		}
		sum += inputNum
	}
	fmt.Println(sum)
}