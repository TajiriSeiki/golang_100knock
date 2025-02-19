package main

import "fmt"

func main() {
	var num [10]int
	num = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	
	for i := 0; i <  len(num); i++  {
		fmt.Println(num[i])
	}
}