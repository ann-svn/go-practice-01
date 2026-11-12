package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
}

func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

func tasksByStatus(tasks []Task, status string) []Task {
	new_tasks := []Task{}
	for _, task := range tasks {
		if task.Status == status {
			new_tasks = append(new_tasks, task)
		}
	}
	return new_tasks
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	found := findTaskByID(tasks, 2)
	if found != nil {
		found.Complete()
		found.SetPriority(3)
	}

	for _, task := range tasks {
		fmt.Println(task.Summary())
	}

	cnt_done := 0
	for _, task := range tasks {
		if task.Status == "done" {
			cnt_done++
		}
	}
	fmt.Println("Задач завершено:", cnt_done)

	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}

	status := "done"
	fmt.Printf("Задачи со статусом %s:\n", status)
	tasks_done := tasksByStatus(tasks, status)
	if len(tasks_done) != 0 {
		for _, task := range tasks_done {
			fmt.Printf("#%d %s\n", task.ID, task.Title)
		}
	} else {
		fmt.Println("Нет задач с таким статусом.")
	}
}
