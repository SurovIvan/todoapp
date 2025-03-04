// Файл: storage/storage.go
package storage

import "errors"

// Task представляет собой структуру задачи
type Task struct {
	ID    int
	Title string
}

// Storage представляет собой хранилище задач
type Storage struct {
	tasks  []Task
	nextID int
}

// NewStorage создаёт новое хранилище
func NewStorage() *Storage {
	return &Storage{
		tasks:  []Task{},
		nextID: 1,
	}
}

// AddTask добавляет задачу в хранилище
func (s *Storage) AddTask(title string) Task {
	task := Task{
		ID:    s.nextID,
		Title: title,
	}
	s.tasks = append(s.tasks, task)
	s.nextID++
	return task
}

// GetTasks возвращает список всех задач
func (s *Storage) GetTasks() []Task {
	return s.tasks
}

// DeleteTask удаляет задачу по ID
func (s *Storage) DeleteTask(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
