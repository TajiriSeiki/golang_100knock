package main

import "fmt"

func main() {
	var array [10]int
	var num int

	fmt.Print("整数を入力してください: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	for i := 0; i <  len(array); i++  {
		array[i] = num
		fmt.Println(array[i])
	}
}