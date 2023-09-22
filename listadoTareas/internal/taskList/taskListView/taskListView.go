package taskListView

import (
	"fmt"
	task "listadoTareas/internal/taskList"
	"time"
)

type TaskListView struct {
}

func NewTaskListView() *TaskListView {
	return &TaskListView{}
}

func (tv *TaskListView) PrintTasks(tasks []task.Task) {
	for _, task := range tasks {
		fmt.Printf("%d %s %s %s %s %s %s\n", task.Id, task.Name, task.Description, getState(task.State), task.StartDate.Format("2006-01-02"), task.EndDate.Format("2006-01-02"), getPriority(task.Priority))
	}
}

func (tv *TaskListView) Input() (int, task.Task) {
	var startDate, endDate string
	var task task.Task
	fmt.Println("Detalles de la Tarea a Ingresar:")
	fmt.Print("Id: ")
	fmt.Scanln(&task.Id)
	fmt.Print("Nombre: ")
	fmt.Scanln(&task.Name)
	fmt.Print("Descripcion: ")
	fmt.Scanln(&task.Description)
	fmt.Print("Estado (1=Pendiente, 2=En Progreso, 3=Completado): ")
	fmt.Scanln(&task.State)
	fmt.Print("Fecha de Inicio (YYYY-MM-DD): ")
	//fmt.Scan(&task.StartDate)
	fmt.Scanln(&startDate)
	task.StartDate, _ = time.Parse("2006-01-02", startDate)
	fmt.Print("Fecha de Fin (YYYY-MM-DD): ")
	//fmt.Scan(&task.EndDate)
	fmt.Scanln(&endDate)
	task.EndDate, _ = time.Parse("2006-01-02", startDate)
	fmt.Print("Prioridad (1=Bajo, 2=Medio, 3=Alto): ")
	fmt.Scanln(&task.Priority)
	return task.Id, task
}

func getState(state int) string {
	switch state {
	case 1:
		return "Pendiente"
	case 2:
		return "En Progreso"
	case 3:
		return "Completado"
	default:
		return ""
	}
}

func getPriority(priority int) string {
	switch priority {
	case 1:
		return "Bajo"
	case 2:
		return "Medio"
	case 3:
		return "Alto"
	default:
		return ""
	}
}
