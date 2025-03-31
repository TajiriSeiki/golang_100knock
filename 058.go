package main

import "fmt"

func main() {
	var num [5]int

	for i:=0; i < len(num); i++ {
		fmt.Printf("input data[%d]: ", i)
		_, err := fmt.Scanln(&num[i])
		if err != nil {
			fmt.Println("エラーが発生しました:", err)
			return
		}
	}
	row := ""
	for i:=0; i < len(num); i++ {
		row = ""
		for j:=0; j < num[i]; j++{
			row += "*"
			if j !=0 && j % 4 == 0{
				row += " "
			}
		}
		fmt.Printf("%d	:%s\n", num[i], row)
	}
}