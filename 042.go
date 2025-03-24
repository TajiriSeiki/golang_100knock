package main

import "fmt"

func main() {
	var inputNum int

	lastNum := 0
	ascFlg := true

	for i := 0; i < 3; i++ {
		fmt.Print("整数を入力してください: ")
		_, err := fmt.Scanln(&inputNum)
		if err != nil {
			fmt.Println("エラーが発生しました:", err)
			return
		}
		if i == 0 {
			lastNum = inputNum
			continue
		}
		if ascFlg {
			ascFlg = lastNum <= inputNum
			lastNum = inputNum
		}
	}
	if ascFlg {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
	
}