package random

import (
	"math/rand"
	"time"
)

// GenerateRandomGrid generates a random grid with the given dimensions
func GenerateRandomGrid(rows, cols int) [][]bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
		for j := range grid[i] {
			grid[i][j] = r.Intn(2) == 1
		}
	}
	return grid
}
