package life

import "github.com/flexer2006/lms_gameLife/pkg/random"

type GameOfLife struct {
	grid    [][]bool
	newGrid [][]bool
}

// NewGameOfLife creates a new game with the specified grid size
func NewGameOfLife(rows, cols int) *GameOfLife {
	grid := make([][]bool, rows)
	newGrid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
		newGrid[i] = make([]bool, cols)
	}
	return &GameOfLife{grid: grid, newGrid: newGrid}
}

// SetGrid sets the initial state of the grid
func (g *GameOfLife) SetGrid(pattern [][]bool) {
	rows := len(pattern)
	cols := len(pattern[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			g.grid[i][j] = pattern[i][j]
		}
	}
}

// Step performs step of the game
func (g *GameOfLife) Step() {
	rows := len(g.grid)
	cols := len(g.grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			liveNeighbors := g.countLiveNeighbors(i, j)
			if g.grid[i][j] {
				g.newGrid[i][j] = liveNeighbors == 2 || liveNeighbors == 3
			} else {
				g.newGrid[i][j] = liveNeighbors == 3
			}
		}
	}

	// Swap grids
	g.grid, g.newGrid = g.newGrid, g.grid
}

// countLiveNeighbors counts the number of live neighbors for a cell
func (g *GameOfLife) countLiveNeighbors(row, col int) int {
	rows := len(g.grid)
	cols := len(g.grid[0])
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			r := row + i
			c := col + j
			if r >= 0 && r < rows && c >= 0 && c < cols && g.grid[r][c] {
				count++
			}
		}
	}

	return count
}

// GetGrid returns the current state of the grid
func (g *GameOfLife) GetGrid() [][]bool {
	return g.grid
}

// SetRandomState sets a random initial state of the grid
func (g *GameOfLife) SetRandomState() {
	rows := len(g.grid)
	cols := len(g.grid[0])
	randomGrid := random.GenerateRandomGrid(rows, cols)
	g.SetGrid(randomGrid)
}
