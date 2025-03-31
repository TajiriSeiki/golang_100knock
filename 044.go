package main

import "fmt"

func main() {
	var yen int
	var dollar float64

	fmt.Print("換算金額を入力してください: ")
	_, err := fmt.Scanln(&yen)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	fmt.Print("1ドル何円か入力してください: ")
	_, err = fmt.Scanln(&dollar)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	calculatedDollar := int(float64(yen) / dollar)
	calculatedCent := int((float64(yen) / dollar - float64(calculatedDollar)) * 100)
	fmt.Printf("%d円は%dドル%dセント", yen, calculatedDollar, calculatedCent)
}