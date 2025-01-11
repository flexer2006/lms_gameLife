package service

import (
	"github.com/flexer2006/lms_gameLife/pkg/life"
	"github.com/flexer2006/lms_gameLife/pkg/random"
)

type GameService struct {
	game *life.GameOfLife
}

// NewGameService Creates a new game service with a random grid size
func NewGameService() *GameService {
	rows, cols := random.GenerateRandomSize(50, 255)
	service := &GameService{
		game: life.NewGameOfLife(rows, cols),
	}
	service.game.SetRandomState()
	return service
}

// GetNextState executes one step of the game and returns the current state of the grid
func (s *GameService) GetNextState() [][]bool {
	s.game.Step()
	return s.game.GetGrid()
}
