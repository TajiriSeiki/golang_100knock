package main

import "fmt"

func main() {
	array := [10]int{3, 7, 0, 8, 4, 1, 9, 6, 5, 2};
	index := 0
	for i := 0; i < 10; i++ {
		fmt.Println(array[index])
		index = array[index]
	}
}