package main

import "fmt"

func main() {
	var a int
	var b int
	var tmp int

	fmt.Print("input a: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fmt.Print("input b: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	tmp = a
	a = b
	b = tmp

	fmt.Printf("a = %d, b = %d", a, b)
}