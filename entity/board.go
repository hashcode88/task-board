package entity

import "github.com/google/uuid"

type Board struct {
	id       string
	name     string
	taskList []TaskList
}

func (b *Board) GetName() string {
	return b.name
}

func (b *Board) GetId() string {
	return b.id
}

func (b *Board) GetTaskList() []TaskList {
	return b.taskList
}

func (b *Board) AddNewTaskList(title string) {
	taskList := newTaskList(title)
	b.taskList = append(b.taskList, *taskList)
}

func NewBoard(name string) *Board {
	return &Board{
		id:   uuid.New().String(),
		name: name,
	}
}
