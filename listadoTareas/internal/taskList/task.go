package taskList

import "time"

type TaskListRepository interface {
	AddNewTask(task Task)
	GetAllTasks() []Task
	PendingTasks() int
	UpdateTask(task Task)
	DeleteTask(int)
}

type TaskListView interface {
	PrintTasks([]Task)
	Input() (int, Task)
}

type TaskListProcess interface {
	TaskProcess()
}

type Task struct {
	Id          int
	Name        string
	Description string
	State       int
	StartDate   time.Time
	EndDate     time.Time
	Priority    int
}
