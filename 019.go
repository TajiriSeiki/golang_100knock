package main

import "fmt"

func main() {
	var inputNum int
	var array [5]int

	for i := 0; i <  5; i++  {
		fmt.Print("整数を入力してください: ")
		_, err := fmt.Scanln(&inputNum)
		if err != nil {
			fmt.Println("エラーが発生しました:", err)
			return
		}
		array[i] = inputNum
	}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}