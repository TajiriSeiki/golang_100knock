package main

import "fmt"

func main() {
	for i:=1; i<=100; i++ {
		if i % 15 == 0 {
			fmt.Printf("foobar\n")
		} else if i % 5 == 0 {
			fmt.Printf("bar\n")
		} else if i % 3 == 0 {
			fmt.Printf("foo\n")
		} else {
			fmt.Println(i)
		}
	}
}