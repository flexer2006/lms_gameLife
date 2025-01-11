package service

import (
	"github.com/flexer2006/lms_gameLife/pkg/life"
)

type GameService struct {
	game *life.GameOfLife
}

// NewGameService creates a new game service with the specified grid size
func NewGameService(rows, cols int) *GameService {
	service := &GameService{
		game: life.NewGameOfLife(rows, cols),
	}
	service.game.SetRandomState() // Устанавливаем случайное начальное состояние
	return service
}

// GetNextState executes one step of the game and returns the current state of the grid
func (s *GameService) GetNextState() [][]bool {
	s.game.Step()
	return s.game.GetGrid()
}
