package main

import "fmt"

func main() {
	var inputNum int
	array := [10]int{3, 7, 0, 8, 4, 1, 9, 6, 5, 2};
	fmt.Print("整数を入力してください: ")
	_, err := fmt.Scanln(&inputNum)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fmt.Printf("array[%d] = %d", inputNum, array[inputNum])
}