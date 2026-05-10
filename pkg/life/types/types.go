// Package types defines shared simulation interfaces and structures.
package types

// Grid defines the interface for interacting with the simulation board.
type Grid interface {
	// Get returns the state (alive/dead) of a character at (x, y).
	Get(x, y int) bool
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

// Cell acts as a container for a simulation Character.
type Cell struct {
	Character Character
}

// NextState delegates the transition logic to the underlying character.
func (c *Cell) NextState(neighbors int, grid Grid, x, y int) *Cell {
	return &Cell{Character: c.Character.NextState(neighbors, grid, x, y)}
}

// ApplyEffects delegates the resolution of spread effects to the underlying character.
func (c *Cell) ApplyEffects(effects []SpreadEffect, grid Grid, x, y int) *Cell {
	return &Cell{Character: c.Character.ApplyEffects(effects, grid, x, y)}
}
