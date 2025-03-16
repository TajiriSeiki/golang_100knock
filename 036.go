package main

import "fmt"

func main() {
	array := [10]int{3, 7, 0, 8, 4, 1, 9, 6, 5, 2};
	var inputNum1 int
	var inputNum2 int
	fmt.Print("1つ目の整数を入力してください: ")
	_, err := fmt.Scanln(&inputNum1)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fmt.Print("2つ目の整数を入力してください: ")
	_, err = fmt.Scanln(&inputNum2)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fmt.Printf("%d * %d = %d", array[inputNum1], array[inputNum2], array[inputNum1] * array[inputNum2])
}