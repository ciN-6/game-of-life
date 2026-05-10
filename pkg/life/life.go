// Package life contains the core simulation logic.
package life

import (
	"game-of-life/pkg/life/characters"
	"game-of-life/pkg/life/types"
)

// Board represents the Game of Life board.
type Board struct {
	width, height int
	Cells         []types.Cell
}

// NewGrid creates a new Board with the given dimensions.
func NewGrid(width, height int) *Board {
	board := &Board{
		width:  width,
		height: height,
		Cells:  make([]types.Cell, width*height),
	}
	for i := range board.Cells {
		board.Cells[i] = types.Cell{Character: &characters.BaseCharacter{BaseCell: characters.BaseCell{UnderPop: 2, OverPop: 3, Repro: 3}}}
	}
	return board
}

// Set sets the state of a cell at the given coordinates.
func (g *Board) Set(x, y int, alive bool) {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return
	}
	idx := y*g.width + x
	u, o, r := g.Cells[idx].Character.GetRules()
	if alive {
		g.Cells[idx] = types.Cell{Character: &characters.LivingCharacter{BaseCharacter: characters.BaseCharacter{BaseCell: characters.BaseCell{Alive: true, UnderPop: u, OverPop: o, Repro: r}}}}
	} else {
		g.Cells[idx] = types.Cell{Character: &characters.LivingCharacter{BaseCharacter: characters.BaseCharacter{BaseCell: characters.BaseCell{Alive: false, UnderPop: u, OverPop: o, Repro: r}}}}
	}
}

// Get returns the alive status of a cell at the given coordinates.
func (g *Board) Get(x, y int) bool {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return false
	}
	return g.Cells[y*g.width+x].Character.IsAlive()
}

// countNeighbors returns the number of live neighbors for a cell at (x, y).
func (g *Board) countNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if g.Get(x+dx, y+dy) {
				count++
			}
		}
	}
	return count
}

// Step advances the simulation by one generation.
func (g *Board) Step() {
	// Reset highlight status
	for i := range g.Cells {
		g.Cells[i].Character.SetHighlighted(false)
	}

	// 1. Collect Spread Effects
	effects := make(map[int][]types.SpreadEffect)
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			cell := g.Cells[y*g.width+x]
			cellEffects := cell.Character.Spread(g, x, y)
			for _, e := range cellEffects {
				if e.TargetX >= 0 && e.TargetX < g.width && e.TargetY >= 0 && e.TargetY < g.height {
					idx := e.TargetY*g.width + e.TargetX
					effects[idx] = append(effects[idx], e)
				}
			}
		}
	}

	// 2. Apply Effects
	applied := make([]types.Cell, len(g.Cells))
	for i := range g.Cells {
		applied[i] = *g.Cells[i].ApplyEffects(effects[i], g, i%g.width, i/g.width)
	}

	// 3. Calculate Next State
	next := make([]types.Cell, len(g.Cells))
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			idx := y*g.width + x
			neighbors := g.countNeighbors(x, y)
			nextChar := applied[idx].Character.NextState(neighbors, g, x, y)
			next[idx] = types.Cell{Character: nextChar}
		}
	}
	g.Cells = next
}

// GetCell returns the cell structure at the given coordinates.
func (g *Board) GetCell(x, y int) *types.Cell {
	if x < 0 || x >= g.width || y < 0 || y >= g.height {
		return nil
	}
	return &g.Cells[y*g.width+x]
}

// Width returns the grid width.
func (g *Board) Width() int { return g.width }

// Height returns the grid height.
func (g *Board) Height() int { return g.height }

// Clear resets the board to a default state.
func (g *Board) Clear() {
	for i := range g.Cells {
		u, o, r := g.Cells[i].Character.GetRules()
		g.Cells[i] = types.Cell{Character: &characters.BaseCharacter{BaseCell: characters.BaseCell{UnderPop: u, OverPop: o, Repro: r}}}
	}
}

// CountAlive returns the total count of live cells on the board.
func (g *Board) CountAlive() int {
	count := 0
	for _, cell := range g.Cells {
		if cell.Character.IsAlive() {
			count++
		}
	}
	return count
}
