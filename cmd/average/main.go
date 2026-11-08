package main

import "fmt"

func main() {
	var m1 int = 20
	var m2 int = 25
	var m3 int = 35
	var sum int = m1 + m2 + m3
	var avg float64 = float64(sum) / 3

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Среднее арифмитическое: %.2f\n", avg)
}
