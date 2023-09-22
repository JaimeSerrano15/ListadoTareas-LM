package taskListRepo

import (
	task "listadoTareas/internal/taskList"
	"reflect"
	"testing"
	"time"
)

func TestNewTaskList(t *testing.T) {
	type args struct {
		tasks []task.Task
	}
	tests := []struct {
		name string
		args args
		want *taskList
	}{
		{
			name: "Test New Task List Repo",
			args: args{
				tasks: []task.Task{},
			},
			want: &taskList{
				tasks: []task.Task{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskList(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskList_AddNewTask(t *testing.T) {
	type fields struct {
		tasks []task.Task
	}
	type args struct {
		task task.Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Add a New task",
			fields: fields{
				tasks: []task.Task{},
			},
			args: args{
				task: task.Task{
					Id:          0,
					Name:        "test",
					Description: "test",
					State:       1,
					StartDate:   time.Time{},
					EndDate:     time.Time{},
					Priority:    1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &taskList{
				tasks: tt.fields.tasks,
			}
			tl.AddNewTask(tt.args.task)
		})
	}
}

func Test_taskList_DeleteTask(t *testing.T) {
	type fields struct {
		tasks []task.Task
	}
	type args struct {
		taskId int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Delete a Task",
			fields: fields{
				tasks: []task.Task{
					{
						Id:          1,
						Name:        "",
						Description: "",
						State:       0,
						StartDate:   time.Time{},
						EndDate:     time.Time{},
						Priority:    0,
					},
				},
			},
			args: args{
				taskId: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &taskList{
				tasks: tt.fields.tasks,
			}
			tl.DeleteTask(tt.args.taskId)
		})
	}
}

func Test_taskList_GetAllTasks(t *testing.T) {
	type fields struct {
		tasks []task.Task
	}
	tests := []struct {
		name   string
		fields fields
		want   []task.Task
	}{
		{
			name: "Test Get All Tasks",
			fields: fields{
				tasks: []task.Task{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &taskList{
				tasks: tt.fields.tasks,
			}
			if got := tl.GetAllTasks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskList_PendingTasks(t *testing.T) {
	type fields struct {
		tasks []task.Task
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test Get Pending Tasks",
			fields: fields{
				tasks: []task.Task{
					{
						Id:          0,
						Name:        "",
						Description: "",
						State:       1,
						StartDate:   time.Time{},
						EndDate:     time.Time{},
						Priority:    0,
					},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &taskList{
				tasks: tt.fields.tasks,
			}
			if got := tl.PendingTasks(); got != tt.want {
				t.Errorf("PendingTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskList_UpdateTask(t *testing.T) {
	type fields struct {
		tasks []task.Task
	}
	type args struct {
		task task.Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Test Update a Task",
			fields: fields{
				tasks: []task.Task{
					{
						Id:          1,
						Name:        "",
						Description: "",
						State:       0,
						StartDate:   time.Time{},
						EndDate:     time.Time{},
						Priority:    0,
					},
				},
			},
			args: args{
				task: task.Task{
					Id:          1,
					Name:        "",
					Description: "",
					State:       0,
					StartDate:   time.Time{},
					EndDate:     time.Time{},
					Priority:    0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tl := &taskList{
				tasks: tt.fields.tasks,
			}
			tl.UpdateTask(tt.args.task)
		})
	}
}
