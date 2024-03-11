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
