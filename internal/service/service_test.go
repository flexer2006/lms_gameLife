package service

import (
	"testing"
)

func TestNewGameService(t *testing.T) {
	service := NewGameService(5, 5)
	if service == nil {
		t.Error("Expected NewGameService to return a non-nil value")
	}
	if len(service.game.GetGrid()) != 5 {
		t.Errorf("Expected grid to have 5 rows, got %d", len(service.game.GetGrid()))
	}
	if len(service.game.GetGrid()[0]) != 5 {
		t.Errorf("Expected grid to have 5 columns, got %d", len(service.game.GetGrid()[0]))
	}
}

func TestSetInitialState(t *testing.T) {
	service := NewGameService(5, 5)
	pattern := [][]bool{
		{false, true, false, false, false},
		{false, false, true, false, false},
		{true, true, true, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}
	service.SetInitialState(pattern)
	grid := service.game.GetGrid()
	for i := range pattern {
		for j := range pattern[i] {
			if grid[i][j] != pattern[i][j] {
				t.Errorf("Expected cell (%d, %d) to be %v, got %v", i, j, pattern[i][j], grid[i][j])
			}
		}
	}
}

func TestGetNextState(t *testing.T) {
	service := NewGameService(5, 5)
	pattern := [][]bool{
		{false, true, false, false, false},
		{false, false, true, false, false},
		{true, true, true, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}
	service.SetInitialState(pattern)
	nextState := service.GetNextState()
	expectedState := [][]bool{
		{false, false, false, false, false},
		{true, false, true, false, false},
		{false, true, true, false, false},
		{false, true, false, false, false},
		{false, false, false, false, false},
	}
	for i := range expectedState {
		for j := range expectedState[i] {
			if nextState[i][j] != expectedState[i][j] {
				t.Errorf("Expected cell (%d, %d) to be %v, got %v", i, j, expectedState[i][j], nextState[i][j])
			}
		}
	}
}
