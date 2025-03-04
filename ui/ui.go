// Файл: ui/ui.go
package ui

import (
	"fmt"
	"github.com/SurovIvan/todoapp/tasks"
	"github.com/fatih/color"
)

// ShowTasks отображает список задач
func ShowTasks(taskManager *tasks.TaskManager) {
	tasks := taskManager.ListTasks()
	if len(tasks) == 0 {
		color.Yellow("No tasks found.")
		return
	}
	color.Cyan("Your tasks:")
	for _, task := range tasks {
		fmt.Printf("%d: %s\n", task.ID, task.Title)
	}
}

// AddTaskPrompt запрашивает у пользователя новую задачу
func AddTaskPrompt(taskManager *tasks.TaskManager) {
	var title string
	color.Green("Enter task title:")
	fmt.Scanln(&title)
	task := taskManager.AddTask(title)
	color.Green("Task added with ID: %d", task.ID)
}

// DeleteTaskPrompt запрашивает ID задачи для удаления
func DeleteTaskPrompt(taskManager *tasks.TaskManager) {
	var id int
	color.Yellow("Enter task ID to delete:")
	fmt.Scanln(&id)
	err := taskManager.DeleteTask(id)
	if err != nil {
		color.Red("Error: %s", err)
	} else {
		color.Green("Task deleted successfully.")
	}
}
