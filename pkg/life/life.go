package life

import (
	"bytes"
	"log"
)

type GameOfLife struct {
	grid [][]bool
}

// NewGameOfLife creates a new game with the specified grid size
func NewGameOfLife(rows, cols int) *GameOfLife {
	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
	}
	return &GameOfLife{grid: grid}
}

// SetGrid sets the initial grid with a pattern
func (g *GameOfLife) SetGrid(pattern [][]bool) {
	rows := len(pattern)
	cols := len(pattern[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			g.grid[i][j] = pattern[i][j]
		}
	}
}

// Step performs one step of the game
func (g *GameOfLife) Step() {
	rows := len(g.grid)
	cols := len(g.grid[0])
	newGrid := make([][]bool, rows)

	for i := range newGrid {
		newGrid[i] = make([]bool, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			liveNeighbors := g.countLiveNeighbors(i, j)
			if g.grid[i][j] {
				newGrid[i][j] = liveNeighbors == 2 || liveNeighbors == 3
			} else {
				newGrid[i][j] = liveNeighbors == 3
			}
		}
	}

	g.grid = newGrid
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

// String returns a string representation of the grid
func (g *GameOfLife) String() string {
	var buffer bytes.Buffer
	for _, row := range g.grid {
		for _, cell := range row {
			if cell {
				buffer.WriteString("1 ")
			} else {
				buffer.WriteString("0 ")
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}

// Reset resets the game to the initial state
func (g *GameOfLife) Reset() {
	for i := range g.grid {
		for j := range g.grid[i] {
			g.grid[i][j] = false
		}
	}
}

// Run performs several steps of the game
func (g *GameOfLife) Run(steps int) {
	for i := 0; i < steps; i++ {
		log.Printf("Step %d\n%s", i+1, g)
		g.Step()
	}
}
