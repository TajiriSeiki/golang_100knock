package main

import "fmt"

func main() {
	var year int

	fmt.Print("input year: ")
	_, err := fmt.Scanln(&year)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	if ((year % 4 == 0) && (year % 100 != 0)) || ((year % 4 == 0) && (year % 400 == 0)) {
		fmt.Printf("%d年は閏年である", year)
	} else {
		fmt.Printf("%d年は閏年でない", year)
	}
}