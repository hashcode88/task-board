package entity

import "github.com/google/uuid"

type TaskList struct {
	id    string
	title string
}

func newTaskList(title string) *TaskList {
	return &TaskList{
		id:    uuid.New().String(),
		title: title,
	}
}
