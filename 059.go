package main

import "fmt"

func main() {
	var array1 [3][3]int
	var array2 [3][3]int
	var result [3][3]int
	
	fmt.Print("1つめの行列: \n")
	for i:=0; i < 3; i++ {
		_, err := fmt.Scanf("%d %d %d\n", &array1[i][0], &array1[i][1],&array1[i][2])
		if err != nil {
			fmt.Println("エラーが発生しました:", err)
			return
		}
	}
	fmt.Print("2つめの行列: \n")
	for i:=0; i < 3; i++ {
		_, err := fmt.Scanf("%d %d %d\n", &array2[i][0], &array2[i][1],&array2[i][2])
		if err != nil {
			fmt.Println("エラーが発生しました:", err)
			return
		}
	}

	for i:=0; i < len(array1); i++ {
		for j:=0; j < len(array2); j++ {
			result[i][j] = array1[i][j] + array2[i][j]
		}
	}

	fmt.Println("和")
	for k:=0; k < len(result); k++{
		fmt.Printf("%d	%d	%d\n", result[k][0], result[k][1], result[k][2])
	}
}