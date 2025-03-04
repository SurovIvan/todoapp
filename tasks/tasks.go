// Файл: tasks/tasks.go
package tasks

import (
	"todoapp/storage"
)

// TaskManager управляет задачами
type TaskManager struct {
	storage *storage.Storage
}

// NewTaskManager создаёт новый менеджер задач
func NewTaskManager() *TaskManager {
	return &TaskManager{
		storage: storage.NewStorage(),
	}
}

// AddTask добавляет задачу
func (tm *TaskManager) AddTask(title string) storage.Task {
	return tm.storage.AddTask(title)
}

// ListTasks возвращает список задач
func (tm *TaskManager) ListTasks() []storage.Task {
	return tm.storage.GetTasks()
}

// DeleteTask удаляет задачу по ID
func (tm *TaskManager) DeleteTask(id int) error {
	return tm.storage.DeleteTask(id)
}
