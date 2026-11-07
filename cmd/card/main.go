package main

import "fmt"

func main() {
	var name string = "Аня"
	var group string = "ТЭП-04-26"
	const course = "Бэкенд-разработка на языке Go"
	completed := 1
	var interest bool = true

	fmt.Println("Участник:", name)
	fmt.Printf("Группа: %s\n", group)
	fmt.Println("Курс:", course)
	fmt.Printf("Завершено работ: %d\n", completed)
	fmt.Printf("Интерес к бэкенду: %t\n", interest)

	completed = 2
	fmt.Printf("После изменения: %d\n", completed)
}
