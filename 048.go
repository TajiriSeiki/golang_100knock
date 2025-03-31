package main

import "fmt"

func main() {
	var a int

	fmt.Print("input a: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	if a < 2 {
		fmt.Println("2以下です")
		return
	}

	for i:=1; true; i++ {
		if (a == 1){
			break
		} else if (a % 2 == 0) {
			a = a / 2
		} else {
			a = a * 3 + 1
		}
		fmt.Printf("%d: %d\n", i, a)
	}
}