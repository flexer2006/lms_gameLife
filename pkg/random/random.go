package random

import (
	"math/rand"
	"time"
)

// GenerateRandomSize generates random grid sizes in the specified range
func GenerateRandomSize(min, max int) (int, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rows := r.Intn(max-min+1) + min
	cols := r.Intn(max-min+1) + min
	return rows, cols
}

// GenerateRandomGrid generates a random grid with the specified dimensions
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
