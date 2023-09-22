package taskListService

import (
	"fmt"
	"listadoTareas/internal/taskList"
	"sort"
)

type TaskListService struct {
	taskListRepo taskList.TaskListRepository
	taskListView taskList.TaskListView
}

func NewTaskListService(taskListRepo taskList.TaskListRepository, taskListView taskList.TaskListView) *TaskListService {
	return &TaskListService{
		taskListRepo: taskListRepo,
		taskListView: taskListView,
	}
}

func (t *TaskListService) TaskProcess() {
	var choice int

	for choice != 5 {
		fmt.Println("Seleccione una opcion:")
		fmt.Println("1. Agregar")
		fmt.Println("2. Mostrar")
		fmt.Println("3. Eliminar")
		fmt.Println("4. Actualizar")
		fmt.Println("5. Salir")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			t.addTask()
		case 2:
			t.showTasks()
		case 3:
			t.deleteTask()
		case 4:
			t.updateTask()
		case 5:
			return
		default:
			fmt.Println("Opcion invalida, intente de nuevo")
		}
	}
}

func (t *TaskListService) addTask() {
	_, task := t.taskListView.Input()
	t.taskListRepo.AddNewTask(task)
}

func (t *TaskListService) showTasks() {
	tasks := t.taskListRepo.GetAllTasks()
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Priority < tasks[j].Priority
	})

	t.taskListView.PrintTasks(tasks)
	fmt.Printf("Tareas Pendientes: %d\n", t.taskListRepo.PendingTasks())
}

func (t *TaskListService) updateTask() {
	id, task := t.taskListView.Input()
	task.Id = id
	t.taskListRepo.UpdateTask(task)
}

func (t *TaskListService) deleteTask() {
	id := 0
	fmt.Print("Ingrese el Id de la Tarea a eliminar: ")
	fmt.Scan(&id)
	t.taskListRepo.DeleteTask(id)
}
