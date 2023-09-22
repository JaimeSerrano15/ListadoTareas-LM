package taskListRepo

import task "listadoTareas/internal/taskList"

type taskList struct {
	tasks []task.Task
}

func NewTaskList(tasks []task.Task) *taskList {
	return &taskList{tasks: tasks}
}

func (tl *taskList) AddNewTask(task task.Task) {
	tl.tasks = append(tl.tasks, task)
}

func (tl *taskList) GetAllTasks() []task.Task {
	return tl.tasks
}

func (tl *taskList) PendingTasks() int {
	var pending int

	for _, task := range tl.tasks {
		if task.State == 1 {
			pending++
		}
	}

	return pending
}

func (tl *taskList) UpdateTask(task task.Task) {
	for i, t := range tl.tasks {
		if t.Id == task.Id {
			tl.tasks[i] = task
			break
		}
	}
}

func (tl *taskList) DeleteTask(taskId int) {
	for i, task := range tl.tasks {
		if task.Id == taskId {
			tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
			break
		}
	}
}
