package taskListView

import (
	"listadoTareas/internal/taskList"
	task "listadoTareas/internal/taskList"
	"reflect"
	"testing"
	"time"
)

func TestNewTaskListView(t *testing.T) {
	tests := []struct {
		name string
		want *TaskListView
	}{
		{
			name: "Test Create a New Task List View",
			want: &TaskListView{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskListView(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskListView() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskListView_Input(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 taskList.Task
	}{
		{
			name: "Test Get Input of a New task",
			want: 0,
			want1: taskList.Task{
				Id:          0,
				Name:        "",
				Description: "",
				State:       0,
				StartDate:   time.Time{},
				EndDate:     time.Time{},
				Priority:    0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tv := &TaskListView{}
			got, got1 := tv.Input()
			if got != tt.want {
				t.Errorf("Input() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Input() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTaskListView_PrintTasks(t *testing.T) {
	type args struct {
		tasks []task.Task
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Print the tasks",
			args: args{
				tasks: []task.Task{
					{
						Id:          0,
						Name:        "",
						Description: "",
						State:       0,
						StartDate:   time.Time{},
						EndDate:     time.Time{},
						Priority:    0,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tv := &TaskListView{}
			tv.PrintTasks(tt.args.tasks)
		})
	}
}

func Test_getPriority(t *testing.T) {
	type args struct {
		priority int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPriority(tt.args.priority); got != tt.want {
				t.Errorf("getPriority() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getState(t *testing.T) {
	type args struct {
		state int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getState(tt.args.state); got != tt.want {
				t.Errorf("getState() = %v, want %v", got, tt.want)
			}
		})
	}
}
