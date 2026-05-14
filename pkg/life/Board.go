// Package life contains the core simulation logic.
package life

import (
	"game-of-life/pkg/life/characters"
	"game-of-life/pkg/life/types"
	"sync/atomic"
)

var idCounter int64

func nextID() int {
	return int(atomic.AddInt64(&idCounter, 1))
}

const MaxHistory = 10

// Board represents the simulation grid.
type Board struct {
	width, height int
	Cells         []types.Cell
	history       [][]types.Cell
}

// NewGrid creates a new Board with the given dimensions.
func NewGrid(width, height int) *Board {
	board := &Board{
		width:  width,
		height: height,
		Cells:  make([]types.Cell, width*height),
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := y*width + x
			board.Cells[idx] = types.Cell{
				X:          x,
				Y:          y,
				DeathCount: 0,
				Character:  &characters.BaseCharacter{ID: nextID(), UnderPop: 2, OverPop: 3, Repro: 3},
			}
		}
	}
	return board
}

// SetAlive sets the state of a cell at the given coordinates.
func (board *Board) SetAlive(x, y int, alive bool) {
	if x < 0 || x >= board.width || y < 0 || y >= board.height {
		return
	}
	idx := y*board.width + x
	u, o, r := board.Cells[idx].Character.GetRules()
	if alive {
		board.Cells[idx] = types.Cell{
			X:          x,
			Y:          y,
			DeathCount: board.Cells[idx].DeathCount,
			Character:  &characters.LivingCharacter{BaseCharacter: characters.BaseCharacter{ID: nextID(), UnderPop: u, OverPop: o, Repro: r}},
		}
	} else {
		board.Cells[idx] = types.Cell{
			X:          x,
			Y:          y,
			DeathCount: board.Cells[idx].DeathCount,
			Character:  &characters.BaseCharacter{ID: nextID(), UnderPop: u, OverPop: o, Repro: r},
		}
	}
}

// GetAlive returns the alive status of a cell at the given coordinates.
func (board *Board) GetAlive(x, y int) bool {
	if x < 0 || x >= board.width || y < 0 || y >= board.height {
		return false
	}
	cell := board.Cells[y*board.width+x]
	_, isLiving := cell.Character.(*characters.LivingCharacter)
	return isLiving
}

// countNeighbors returns the number of live neighbors for a cell at (x, y).
func (board *Board) countNeighbors(x, y int) int {
	count := 0
	types.ForEachNeighbor(board, x, y, 1, func(nx, ny int) bool {
		if board.GetAlive(nx, ny) {
			count++
		}
		return true
	})
	return count
}

// Step advances the simulation by one generation.
func (board *Board) Step() {
	// Save current state to history before modifying cells
	board.saveSnapshot()

	// 1. Collect PrepareAction Effects
	effects := make(map[int][]types.SpreadEffect)
	for y := 0; y < board.height; y++ {
		for x := 0; x < board.width; x++ {
			cell := board.Cells[y*board.width+x]
			cellEffects := cell.Character.PrepareAction(board, x, y)
			for _, e := range cellEffects {
				if e.TargetX >= 0 && e.TargetX < board.width && e.TargetY >= 0 && e.TargetY < board.height {
					idx := e.TargetY*board.width + e.TargetX
					effects[idx] = append(effects[idx], e)
				}
			}
		}
	}

	// 2. Apply Action
	applied := make([]types.Cell, len(board.Cells))
	for i := range board.Cells {
		char, cell := board.Cells[i].Character.ApplyAction(effects[i], board, i%board.width, i/board.width)
		applied[i] = types.Cell{
			X:          board.Cells[i].X,
			Y:          board.Cells[i].Y,
			DeathCount: cell.DeathCount,
			Character:  char,
		}
	}

	// 3. Calculate Next State
	next := make([]types.Cell, len(board.Cells))
	for y := 0; y < board.height; y++ {
		for x := 0; x < board.width; x++ {
			idx := y*board.width + x
			neighbors := board.countNeighbors(x, y)
			char, cell := applied[idx].Character.NextState(neighbors, board, x, y)
			next[idx] = types.Cell{
				X:          x,
				Y:          y,
				DeathCount: cell.DeathCount,
				Character:  char,
			}
		}
	}
	board.Cells = next
}

// saveSnapshot saves the current board state to history.
func (board *Board) saveSnapshot() {
	snapshot := make([]types.Cell, len(board.Cells))
	copy(snapshot, board.Cells)
	board.history = append(board.history, snapshot)

	// Maintain history size limit
	if len(board.history) > MaxHistory {
		board.history = board.history[1:]
	}
}

// Undo reverts the board to the previous state.
func (board *Board) Undo() {
	if len(board.history) == 0 {
		return
	}
	// Pop the last snapshot and restore it
	lastIdx := len(board.history) - 1
	board.Cells = board.history[lastIdx]
	board.history = board.history[:lastIdx]
}

// GetCell returns the cell structure at the given coordinates.
func (board *Board) GetCell(x, y int) *types.Cell {
	if x < 0 || x >= board.width || y < 0 || y >= board.height {
		return nil
	}
	return &board.Cells[y*board.width+x]
}

// Width returns the grid width.
func (board *Board) Width() int { return board.width }

// Height returns the grid height.
func (board *Board) Height() int { return board.height }

// Clear resets the board to a default state.
func (board *Board) Clear() {
	for y := 0; y < board.height; y++ {
		for x := 0; x < board.width; x++ {
			idx := y*board.width + x
			u, o, r := board.Cells[idx].Character.GetRules()
			board.Cells[idx] = types.Cell{
				X:          x,
				Y:          y,
				DeathCount: 0,
				Character:  &characters.BaseCharacter{ID: nextID(), UnderPop: u, OverPop: o, Repro: r},
			}
		}
	}
}

// CountAlive returns the total count of live cells on the board.
func (board *Board) CountAlive() int {
	count := 0
	for _, cell := range board.Cells {
		if _, isLiving := cell.Character.(*characters.LivingCharacter); isLiving {
			count++
		}
	}
	return count
}
