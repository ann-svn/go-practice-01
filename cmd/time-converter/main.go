package main

import "fmt"

func main() {
	var totalMinutes int
	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&totalMinutes)

	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	const number_minutes_in_one_hour int = 60
	var hours int = totalMinutes / number_minutes_in_one_hour
	var minutes int = totalMinutes % number_minutes_in_one_hour

	fmt.Printf("Полных часов: %d\n", hours)
	fmt.Printf("Оставшихся минут: %d\n", minutes)
}
