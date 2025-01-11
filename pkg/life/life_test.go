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

/*
func TestGlider(t *testing.T) {

		game := NewGameOfLife(5, 5)

		gliderPattern := [][]bool{
			{false, false, false, false, false},
			{false, true, true, false, false},
			{false, false, true, true, false},
			{false, true, false, false, false},
			{false, false, false, false, false},
		}
		game.SetGrid(gliderPattern)

		game.Run(4)

		expected := [][]bool{
			{false, false, false, false, false},
			{false, false, true, true, false},
			{false, true, false, true, false},
			{false, true, false, false, false},
			{false, false, false, false, false},
		}

		for i := 0; i < len(expected); i++ {
			for j := 0; j < len(expected[i]); j++ {
				if game.grid[i][j] != expected[i][j] {
					t.Errorf("After 4 steps: expected %v at [%d][%d], got %v", expected[i][j], i, j, game.grid[i][j])
				}
			}
		}
	}
*/
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

/*
	func TestToad(t *testing.T) {
		game := NewGameOfLife(5, 5)
		game.grid[1][1] = true
		game.grid[1][2] = true
		game.grid[1][3] = true
		game.grid[2][0] = true
		game.grid[2][1] = true
		game.grid[2][2] = true
		game.Step()
		if !game.grid[0][2] || !game.grid[1][3] || !game.grid[2][0] || !game.grid[3][1] || !game.grid[3][2] {
			t.Error("Expected toad configuration to oscillate")
		}
		game.Step()
		if !game.grid[1][1] || !game.grid[1][2] || !game.grid[1][3] || !game.grid[2][0] || !game.grid[2][1] || !game.grid[2][2] {
			t.Error("Expected toad configuration to oscillate back")
		}
	}
*/
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

/*
func TestPulsar(t *testing.T) {
	game := NewGameOfLife(10, 10)

	for i := 2; i <= 4; i++ {
		game.grid[2][i] = true
		game.grid[7][i] = true
		game.grid[2][i+5] = true
		game.grid[7][i+5] = true
		game.grid[i][2] = true
		game.grid[i][7] = true
		game.grid[i+5][2] = true
		game.grid[i+5][7] = true
	}
	game.Step()

	if !game.grid[1][3] || !game.grid[1][4] || !game.grid[1][5] || !game.grid[6][3] || !game.grid[6][4] || !game.grid[6][5] {
		t.Error("Expected pulsar configuration to oscillate")
	}
	game.Step()
	game.Step()

	for i := 2; i <= 4; i++ {
		if !game.grid[2][i] || !game.grid[7][i] || !game.grid[2][i+5] || !game.grid[7][i+5] || !game.grid[i][2] || !game.grid[i][7] || !game.grid[i+5][2] || !game.grid[i+5][7] {
			t.Error("Expected pulsar configuration to oscillate back")
		}
	}
}
*/
/*
func TestLWSS(t *testing.T) {
	game := NewGameOfLife(10, 10)

	game.grid[1][2] = true
	game.grid[1][3] = true
	game.grid[2][1] = true
	game.grid[2][2] = true
	game.grid[2][3] = true
	game.grid[2][4] = true
	game.grid[3][1] = true
	game.grid[3][2] = true
	game.grid[3][4] = true
	game.grid[4][3] = true
	game.Step()
	game.Step()
	game.Step()
	game.Step()

	if !game.grid[2][3] || !game.grid[2][4] || !game.grid[3][2] || !game.grid[3][3] || !game.grid[3][4] || !game.grid[3][5] || !game.grid[4][2] || !game.grid[4][3] || !game.grid[4][5] || !game.grid[5][4] {
		t.Error("Expected LWSS to move to a new position")
	}
}
*/
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

/*
func TestPentadecathlon(t *testing.T) {
	game := NewGameOfLife(5, 12)

	for i := 1; i <= 10; i++ {
		game.grid[2][i] = true
	}
	game.grid[1][3] = true
	game.grid[1][8] = true
	game.grid[3][3] = true
	game.grid[3][8] = true

	for i := 0; i < 15; i++ {
		game.Step()
	}

	for i := 1; i <= 10; i++ {
		if !game.grid[2][i] {
			t.Error("Expected Pentadecathlon configuration to oscillate back")
		}
	}
	if !game.grid[1][3] || !game.grid[1][8] || !game.grid[3][3] || !game.grid[3][8] {
		t.Error("Expected Pentadecathlon configuration to oscillate back")
	}
}
*/
/*
func TestCross(t *testing.T) {
	game := NewGameOfLife(5, 5)

	game.grid[1][2] = true
	game.grid[2][1] = true
	game.grid[2][2] = true
	game.grid[2][3] = true
	game.grid[3][2] = true
	game.Step()

	if !game.grid[1][2] || !game.grid[2][1] || !game.grid[2][2] || !game.grid[2][3] || !game.grid[3][2] {
		t.Error("Expected cross configuration to oscillate")
	}
	game.Step()
	game.Step()

	if !game.grid[1][2] || !game.grid[2][1] || !game.grid[2][2] || !game.grid[2][3] || !game.grid[3][2] {
		t.Error("Expected cross configuration to oscillate back")
	}
}
*/
/*
func TestMWSS(t *testing.T) {
	game := NewGameOfLife(8, 8)

	game.grid[1][2] = true
	game.grid[1][3] = true
	game.grid[2][1] = true
	game.grid[2][4] = true
	game.grid[3][1] = true
	game.grid[3][4] = true
	game.grid[4][1] = true
	game.grid[4][2] = true
	game.grid[4][3] = true
	game.grid[4][4] = true
	game.grid[5][3] = true

	for i := 0; i < 4; i++ {
		game.Step()
	}

	if !game.grid[2][3] || !game.grid[2][4] || !game.grid[3][2] || !game.grid[3][5] || !game.grid[4][2] || !game.grid[4][5] || !game.grid[5][2] || !game.grid[5][3] || !game.grid[5][4] || !game.grid[5][5] || !game.grid[6][4] {
		t.Error("Expected MWSS to move to a new position")
	}
}
*/
/*
func TestHWSS(t *testing.T) {
	game := NewGameOfLife(10, 10)

	game.grid[1][2] = true
	game.grid[1][3] = true
	game.grid[2][1] = true
	game.grid[2][5] = true
	game.grid[3][1] = true
	game.grid[3][6] = true
	game.grid[4][1] = true
	game.grid[4][6] = true
	game.grid[5][1] = true
	game.grid[5][2] = true
	game.grid[5][3] = true
	game.grid[5][4] = true
	game.grid[5][5] = true
	game.grid[6][4] = true

	for i := 0; i < 4; i++ {
		game.Step()
	}

	if !game.grid[2][3] || !game.grid[2][4] || !game.grid[3][2] || !game.grid[3][6] || !game.grid[4][2] || !game.grid[4][7] || !game.grid[5][2] || !game.grid[5][7] || !game.grid[6][2] || !game.grid[6][3] || !game.grid[6][4] || !game.grid[6][5] || !game.grid[6][6] || !game.grid[7][5] {
		t.Error("Expected HWSS to move to a new position")
	}
}
*/
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

/*
	func TestDiehard(t *testing.T) {
		game := NewGameOfLife(10, 10)

		game.grid[1][7] = true
		game.grid[2][1] = true
		game.grid[2][2] = true
		game.grid[3][2] = true
		game.grid[3][6] = true
		game.grid[3][7] = true
		game.grid[3][8] = true

		for i := 0; i < 130; i++ {
			game.Step()
		}

		if !isDead(game.grid) {
			t.Error("Expected Diehard to die after 130 generations")
		}
	}
*/
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

/*
	func TestGosperGliderGun(t *testing.T) {
		game := NewGameOfLife(40, 40)

		game.grid[5][1] = true
		game.grid[5][2] = true
		game.grid[6][1] = true
		game.grid[6][2] = true
		game.grid[5][11] = true
		game.grid[6][11] = true
		game.grid[7][11] = true
		game.grid[4][12] = true
		game.grid[8][12] = true
		game.grid[3][13] = true
		game.grid[9][13] = true
		game.grid[3][14] = true
		game.grid[9][14] = true
		game.grid[6][15] = true
		game.grid[4][16] = true
		game.grid[8][16] = true
		game.grid[5][17] = true
		game.grid[6][17] = true
		game.grid[7][17] = true
		game.grid[6][18] = true
		game.grid[3][21] = true
		game.grid[4][21] = true
		game.grid[5][21] = true
		game.grid[3][22] = true
		game.grid[4][22] = true
		game.grid[5][22] = true
		game.grid[2][23] = true
		game.grid[6][23] = true
		game.grid[1][25] = true
		game.grid[2][25] = true
		game.grid[6][25] = true
		game.grid[7][25] = true
		game.grid[3][35] = true
		game.grid[4][35] = true
		game.grid[3][36] = true
		game.grid[4][36] = true

		for i := 0; i < 100; i++ {
			game.Step()
		}

		if !hasGliders(game.grid) {
			t.Error("Expected Gosper glider gun to generate gliders")
		}
	}
*/

/*
func TestClock(t *testing.T) {
	game := NewGameOfLife(7, 7)

	game.grid[1][2] = true
	game.grid[2][1] = true
	game.grid[2][3] = true
	game.grid[3][2] = true
	game.grid[4][1] = true
	game.grid[4][3] = true
	game.grid[5][2] = true

	for i := 0; i < 4; i++ {
		game.Step()
	}

	if !isClockOscillating(game.grid) {
		t.Error("Expected Clock configuration to oscillate")
	}
}
*/

func hasGliders(grid [][]bool) bool {
	gliderPatterns := [][][]bool{
		{
			{false, true, false},
			{false, false, true},
			{true, true, true},
		},
	}

	for _, pattern := range gliderPatterns {
		for i := 0; i < len(grid)-len(pattern); i++ {
			for j := 0; j < len(grid[0])-len(pattern[0]); j++ {
				match := true
				for x := 0; x < len(pattern); x++ {
					for y := 0; y < len(pattern[0]); y++ {
						if grid[i+x][j+y] != pattern[x][y] {
							match = false
							break
						}
					}
					if !match {
						break
					}
				}
				if match {
					return true
				}
			}
		}
	}
	return false
}

func matchesPattern(grid [][]bool, x, y int, pattern [][]bool) bool {
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[i]); j++ {
			if grid[x+i][y+j] != pattern[i][j] {
				return false
			}
		}
	}
	return true
}

func isClockOscillating(grid [][]bool) bool {

	originalGrid := make([][]bool, len(grid))
	for i := range grid {
		originalGrid[i] = make([]bool, len(grid[i]))
		copy(originalGrid[i], grid[i])
	}

	game := &GameOfLife{grid: grid}
	for i := 0; i < 4; i++ {
		game.Step()
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != originalGrid[i][j] {
				return false
			}
		}
	}
	return true
}
