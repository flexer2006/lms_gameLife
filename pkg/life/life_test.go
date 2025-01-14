package life

import (
	"testing"
)

func TestBlock(t *testing.T) {
	game := NewGameOfLife(4, 4)
	game.grid[1][1] = true
	game.grid[1][2] = true
	game.grid[2][1] = true
	game.grid[2][2] = true
	game.Step()
	if !game.grid[1][1] || !game.grid[1][2] || !game.grid[2][1] || !game.grid[2][2] {
		t.Error("Expected block configuration to remain stable")
	}
}

func TestBlinker(t *testing.T) {
	game := NewGameOfLife(5, 5)
	game.grid[1][0] = true
	game.grid[1][1] = true
	game.grid[1][2] = true
	game.Step()
	if !game.grid[0][1] || !game.grid[1][1] || !game.grid[2][1] {
		t.Error("Expected blinker configuration to oscillate")
	}
	game.Step()
	if !game.grid[1][0] || !game.grid[1][1] || !game.grid[1][2] {
		t.Error("Expected blinker configuration to oscillate back")
	}
}

func TestBoat(t *testing.T) {
	game := NewGameOfLife(5, 5)
	game.grid[1][1] = true
	game.grid[1][2] = true
	game.grid[2][1] = true
	game.grid[2][3] = true
	game.grid[3][2] = true
	game.Step()
	if !game.grid[1][1] || !game.grid[1][2] || !game.grid[2][1] || !game.grid[2][3] || !game.grid[3][2] {
		t.Error("Expected boat configuration to remain stable")
	}
}

func TestBeehive(t *testing.T) {
	game := NewGameOfLife(5, 5)
	game.grid[1][2] = true
	game.grid[1][3] = true
	game.grid[2][1] = true
	game.grid[2][4] = true
	game.grid[3][2] = true
	game.grid[3][3] = true
	game.Step()
	if !game.grid[1][2] || !game.grid[1][3] || !game.grid[2][1] || !game.grid[2][4] || !game.grid[3][2] || !game.grid[3][3] {
		t.Error("Expected beehive configuration to remain stable")
	}
}

func TestLoaf(t *testing.T) {
	game := NewGameOfLife(6, 6)
	game.grid[1][2] = true
	game.grid[1][3] = true
	game.grid[2][1] = true
	game.grid[2][4] = true
	game.grid[3][2] = true
	game.grid[3][4] = true
	game.grid[4][3] = true
	game.Step()
	if !game.grid[1][2] || !game.grid[1][3] || !game.grid[2][1] || !game.grid[2][4] || !game.grid[3][2] || !game.grid[3][4] || !game.grid[4][3] {
		t.Error("Expected loaf configuration to remain stable")
	}
}

func TestTub(t *testing.T) {
	game := NewGameOfLife(5, 5)
	game.grid[1][2] = true
	game.grid[2][1] = true
	game.grid[2][3] = true
	game.grid[3][2] = true
	game.Step()
	if !game.grid[1][2] || !game.grid[2][1] || !game.grid[2][3] || !game.grid[3][2] {
		t.Error("Expected tub configuration to remain stable")
	}
}

func TestBeacon(t *testing.T) {
	game := NewGameOfLife(6, 6)
	game.grid[1][1] = true
	game.grid[1][2] = true
	game.grid[2][1] = true
	game.grid[3][4] = true
	game.grid[4][3] = true
	game.grid[4][4] = true
	game.Step()
	if !game.grid[1][1] || !game.grid[1][2] || !game.grid[2][1] || !game.grid[3][4] || !game.grid[4][3] || !game.grid[4][4] {
		t.Error("Expected beacon configuration to oscillate")
	}
	game.Step()
	if !game.grid[1][1] || !game.grid[1][2] || !game.grid[2][1] || !game.grid[3][4] || !game.grid[4][3] || !game.grid[4][4] {
		t.Error("Expected beacon configuration to oscillate back")
	}
}

func TestPond(t *testing.T) {
	game := NewGameOfLife(6, 6)
	game.grid[1][2] = true
	game.grid[1][3] = true
	game.grid[2][1] = true
	game.grid[2][4] = true
	game.grid[3][1] = true
	game.grid[3][4] = true
	game.grid[4][2] = true
	game.grid[4][3] = true
	game.Step()
	if !game.grid[1][2] || !game.grid[1][3] || !game.grid[2][1] || !game.grid[2][4] || !game.grid[3][1] || !game.grid[3][4] || !game.grid[4][2] || !game.grid[4][3] {
		t.Error("Expected pond configuration to remain stable")
	}
}

func TestHertz(t *testing.T) {
	game := NewGameOfLife(12, 12)

	game.grid[1][2] = true
	game.grid[1][3] = true
	game.grid[2][1] = true
	game.grid[2][4] = true
	game.grid[3][1] = true
	game.grid[3][4] = true
	game.grid[4][2] = true
	game.grid[4][3] = true
	game.grid[6][2] = true
	game.grid[6][3] = true
	game.grid[7][1] = true
	game.grid[7][4] = true
	game.grid[8][1] = true
	game.grid[8][4] = true
	game.grid[9][2] = true
	game.grid[9][3] = true

	for i := 0; i < 8; i++ {
		game.Step()
	}

	if !game.grid[1][2] || !game.grid[1][3] || !game.grid[2][1] || !game.grid[2][4] || !game.grid[3][1] || !game.grid[3][4] || !game.grid[4][2] || !game.grid[4][3] || !game.grid[6][2] || !game.grid[6][3] || !game.grid[7][1] || !game.grid[7][4] || !game.grid[8][1] || !game.grid[8][4] || !game.grid[9][2] || !game.grid[9][3] {
		t.Error("Expected Hertz configuration to oscillate back")
	}
}

func TestRPentomino(t *testing.T) {
	game := NewGameOfLife(5, 5)

	game.grid[1][2] = true
	game.grid[2][1] = true
	game.grid[2][2] = true
	game.grid[2][3] = true
	game.grid[3][1] = true

	for i := 0; i < 1103; i++ {
		game.Step()
	}

	if !isStable(game, 1103) {
		t.Error("Expected R-pentomino to stabilize after 1103 generations")
	}
}

func TestAcorn(t *testing.T) {
	game := NewGameOfLife(10, 10)

	game.grid[1][3] = true
	game.grid[2][1] = true
	game.grid[2][3] = true
	game.grid[3][2] = true
	game.grid[3][3] = true
	game.grid[3][4] = true
	game.grid[3][5] = true

	for i := 0; i < 5206; i++ {
		game.Step()
	}

	if !isStable(game, 5206) {
		t.Error("Expected Acorn to stabilize after 5206 generations")
	}
}

func isStable(game *GameOfLife, steps int) bool {

	originalGrid := make([][]bool, len(game.grid))
	for i := range game.grid {
		originalGrid[i] = make([]bool, len(game.grid[i]))
		copy(originalGrid[i], game.grid[i])
	}

	for i := 0; i < steps; i++ {
		game.Step()
	}

	for i := range game.grid {
		for j := range game.grid[i] {
			if game.grid[i][j] != originalGrid[i][j] {
				return false
			}
		}
	}
	return true
}

func isDead(grid [][]bool) bool {
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				return false
			}
		}
	}
	return true
}
