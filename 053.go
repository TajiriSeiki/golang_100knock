package main

import "fmt"

func main() {
	var num int

	fmt.Print("input number: ")
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}
	sum := ""
	for i:=2; true; i++ {
		for j:=0; true; j++ {
			if (num % i == 0){
				sum = fmt.Sprintf("%s %d", sum, i)
				num /= i
				continue
			}
			if num == 1 {
				fmt.Println(sum)
				return
			}
			break
		}
	}
}