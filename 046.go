package main

import "fmt"

func main() {
	var num int

	fmt.Print("入場する人数を入力してください: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fee := 600
	smallGroupNum := 5
	smallGroupFee := 550
	largeGroupNum := 20
	largeGroupFee := 500

	if num < smallGroupNum {
		fmt.Printf("料金: ￥%d", num * fee)
	} else if(num < largeGroupNum) {
		fmt.Printf("料金: ￥%d", num * smallGroupFee)
	} else {
		fmt.Printf("料金: ￥%d", num * largeGroupFee)
	}
}