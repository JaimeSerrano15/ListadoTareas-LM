package bootstrap

import (
	task "listadoTareas/internal/taskList"
	taskListRepo2 "listadoTareas/internal/taskList/taskListRepo"
	"listadoTareas/internal/taskList/taskListService"
	taskListView2 "listadoTareas/internal/taskList/taskListView"
)

func Run() {
	tasks := []task.Task{}

	taskListRepo := taskListRepo2.NewTaskList(tasks)
	taskListView := taskListView2.NewTaskListView()

	taskListSvc := taskListService.NewTaskListService(taskListRepo, taskListView)

	taskListSvc.TaskProcess()
}
