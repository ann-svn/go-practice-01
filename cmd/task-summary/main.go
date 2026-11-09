package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	count := 0
	for _, status := range statuses {
		if status == target {
			count++
		}
	}
	return count
}

// findStatus находит индекс первого совпадения или -1, если статус не найден.
func findStatus(statuses []string, target string) int {
	for i, status := range statuses {
		if status == target {
			return i
		}
	}
	return -1
}

// completionPercent находит долю выполненных задач.
func completionPercent(statuses []string) float64 {
	if len(statuses) == 0 {
		return 0
	}
	return 100 * float64(countStatus(statuses, "done")) / float64(len(statuses))
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")
	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}
	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d\n", counts["new"], counts["in_progress"], counts["done"])

	fmt.Printf("Индекс статуса done: %d\n", findStatus(statuses, "new"))
	fmt.Printf("Индекс статуса blocked: %d\n", findStatus(statuses, "blocked"))
	fmt.Printf("Индекс статуса new в пустом списке: %d\n", findStatus([]string{}, "new"))

	fmt.Printf("Доля выполненных задач: %.1f\n", completionPercent(statuses))
	fmt.Printf("Доля выполненных задач на пустом списке: %.1f\n", completionPercent([]string{}))
}
