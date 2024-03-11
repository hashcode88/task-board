package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitub.com/hashcode88/task-board/entity"
)

func TestCreateNewBoard(t *testing.T) {
	board := entity.NewBoard("Aplicação web")
	assert.NotNil(t, board)
	assert.Equal(t, board.GetName(), "Aplicação web")
}

func TestAddNewTaskList(t *testing.T) {
	board := entity.NewBoard("Aplicação web")
	board.AddNewTaskList("Não iniciadas")
	assert.Equal(t, len(board.GetTaskList()), 1)
}
