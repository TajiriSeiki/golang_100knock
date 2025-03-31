package main

import "fmt"

func main() {
	space := ""
	for i:=1; i<=9; i++ {
		for j:=1; j<=9; j++ {
			if i*j < 10 {
				space = " 	"
			} else {
				space = "	"
			}
			fmt.Printf("%d%s", i * j, space)
		}
		fmt.Printf("\n")
	}
}