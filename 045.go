package main

import "fmt"
import "math"

func main() {
	var distance float64

	fmt.Print("乗車距離をを入力してください: ")
	_, err := fmt.Scanln(&distance)
	if err != nil {
		fmt.Println("エラーが発生しました:", err)
		return
	}

	initialDistance := 1700.0
	initialFee := 610.0
	runnningDistance := 313.0
	runnningFee := 80.0

	if distance <= initialDistance {
		fmt.Printf("金額: %g", initialFee)
	} else {
		fmt.Printf("金額: %g", initialFee + math.Ceil((distance - initialDistance) / runnningDistance) * runnningFee)
	}
}