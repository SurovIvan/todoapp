// Файл: main.go
package main

import (
	"fmt"
	"todoapp/tasks"
	"todoapp/ui"
)

func main() {
	taskManager := tasks.NewTaskManager()

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. View tasks")
		fmt.Println("2. Add task")
		fmt.Println("3. Delete task")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ui.ShowTasks(taskManager)
		case 2:
			ui.AddTaskPrompt(taskManager)
		case 3:
			ui.DeleteTaskPrompt(taskManager)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
