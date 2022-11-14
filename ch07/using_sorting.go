package ch07

import (
	"fmt"
	"sort"

	"github.com/google/uuid"
)

// references
// ---------------------------------------------------------------------------------
// http://c.biancheng.net/view/81.html

func UsingSorting() {
	var taskList TaskList
	taskList = append(taskList, newTask("Check out source codes"))
	taskList = append(taskList, newTask("Unit test"))
	taskList = append(taskList, newTask("Build"))
	taskList = append(taskList, newTask("Package"))
	taskList = append(taskList, newTask("Deploy"))

	for i, task := range taskList {
		fmt.Printf("%d. %s > %s\n", i, task.id, task.name)
	}
	fmt.Println("--------------------------------------------------------------------------")

	sort.Sort(taskList)
	for i, task := range taskList {
		fmt.Printf("%d. %s > %s\n", i, task.id, task.name)
	}
}

// -------------------------------------------------------------------------------------
type Task struct {
	id   string
	name string
}

func newTask(name string) *Task {
	return &Task{
		id:   uuid.NewString(),
		name: name,
	}
}

type TaskList []*Task

func (list TaskList) Len() int {
	return len(list)
}

func (list TaskList) Less(i int, j int) bool {
	left := list[i]
	right := list[j]
	return left.id < right.id
}

func (list TaskList) Swap(i int, j int) {
	list[i], list[j] = list[j], list[i]
}
