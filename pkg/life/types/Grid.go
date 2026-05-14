// Package types defines shared simulation interfaces and structures.
package types

// Grid defines the interface for interacting with the simulation board.
type Grid interface {
	// GetCell returns the cell structure at (x, y).
	GetCell(x, y int) *Cell
	// Width returns the grid width.
	Width() int
	// Height returns the grid height.
	Height() int
}

// SpreadEffect represents a change propagated by a character to another cell.
type SpreadEffect struct {
	TargetX, TargetY           int
	SourceX, SourceY           int
	DestinationX, DestinationY int
	NewCell                    Cell
	Weight                     int
	TargetType                 string
}

// Cell acts as a container for a simulation Character and its coordinates.
type Cell struct {
	X, Y       int       // Coordinates on the grid
	DeathCount int       // Persistent property tracking entity deaths
	Character  Character // Behavior implementation
}

// ForEachNeighbor calls fn for each neighbor of (x, y) within grid bounds within the specified radius.
// If fn returns false, iteration stops early.
func ForEachNeighbor(grid Grid, x, y, radius int, fn func(nx, ny int) bool) {
	for dy := -radius; dy <= radius; dy++ {
		for dx := -radius; dx <= radius; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < grid.Width() && ny >= 0 && ny < grid.Height() {
				if !fn(nx, ny) {
					return
				}
			}
		}
	}
}
