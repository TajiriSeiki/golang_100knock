package main

import "fmt"

func main() {
	var result string

	for i:=0; i < 3; i++ {
		for j:=0; j < 9; j++ {
			result += "とんで"
		}
		for k:=0; k < 3; k++ {
			result += "まわって"
		}
		result += "まわる\n"
	}
	fmt.Println(result)
}