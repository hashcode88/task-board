package entity

import "github.com/google/uuid"

type Board struct {
	id   string
	name string
}

func (b *Board) GetName() string {
	return b.name
}

func (b *Board) GetId() string {
	return b.id
}

func NewBoard(name string) *Board {
	return &Board{
		id:   uuid.New().String(),
		name: name,
	}
}
