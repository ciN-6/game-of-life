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
	activeCells   []int
}

// NewGrid creates a new Board with the given dimensions.
func NewGrid(width, height int) *Board {
	board := &Board{
		width:       width,
		height:      height,
		Cells:       make([]types.Cell, width*height),
		activeCells: []int{},
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := y*width + x
			board.Cells[idx] = types.Cell{
				X:          x,
				Y:          y,
				DeathCount: 0,
				Character:  nil,
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

	// Default rules if no character exists
	u, o, r := 2, 3, 3
	if board.Cells[idx].Character != nil {
		u, o, r = board.Cells[idx].Character.GetRules()
	}

	hadCharacter := board.Cells[idx].Character != nil

	if alive {
		board.Cells[idx] = types.Cell{
			X:          x,
			Y:          y,
			DeathCount: board.Cells[idx].DeathCount,
			Character:  &characters.LivingCharacter{ID: nextID(), UnderPop: u, OverPop: o, Repro: r},
		}
		if !hadCharacter {
			board.activeCells = append(board.activeCells, idx)
		}
	} else {
		board.Cells[idx] = types.Cell{
			X:          x,
			Y:          y,
			DeathCount: board.Cells[idx].DeathCount,
			Character:  nil,
		}
		if hadCharacter {
			board.removeFromActiveCells(idx)
		}
	}
}

func (board *Board) removeFromActiveCells(idx int) {
	for i, v := range board.activeCells {
		if v == idx {
			board.activeCells = append(board.activeCells[:i], board.activeCells[i+1:]...)
			break
		}
	}
}

func (board *Board) updateActiveCells() {
	board.activeCells = []int{}
	for i, cell := range board.Cells {
		if cell.Character != nil {
			board.activeCells = append(board.activeCells, i)
		}
	}
}

// GetAlive returns the alive status of a cell at the given coordinates.
func (board *Board) GetAlive(x, y int) bool {
	if x < 0 || x >= board.width || y < 0 || y >= board.height {
		return false
	}
	cell := board.Cells[y*board.width+x]
	if cell.Character == nil {
		return false
	}
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

	intent := board.collectIntent()
	applied := board.applyActions(intent)
	nextCells, nextActiveCells := board.calculateNextState(applied)
	board.Cells = nextCells
	board.activeCells = nextActiveCells
}

func (board *Board) collectIntent() map[int][]types.SpreadEffect {
	effects := make(map[int][]types.SpreadEffect)
	for _, idx := range board.activeCells {
		cell := board.Cells[idx]
		// Character should not be nil if activeCells is maintained correctly
		if cell.Character == nil {
			continue
		}
		cellEffects := cell.Character.PrepareAction(board, cell.X, cell.Y)
		for _, e := range cellEffects {
			if e.TargetX >= 0 && e.TargetX < board.width && e.TargetY >= 0 && e.TargetY < board.height {
				targetIdx := e.TargetY*board.width + e.TargetX
				effects[targetIdx] = append(effects[targetIdx], e)
			}
		}
	}
	return effects
}

func (board *Board) applyActions(effects map[int][]types.SpreadEffect) []types.Cell {
	applied := make([]types.Cell, len(board.Cells))
	for i := range board.Cells {
		char := board.Cells[i].Character
		var newChar types.Character
		var newCell types.Cell
		if char == nil {
			// Apply effects to empty cell
			if len(effects[i]) > 0 {
				newChar = effects[i][0].NewCell.Character
				newCell = board.Cells[i]
			} else {
				newChar = nil
				newCell = board.Cells[i]
			}
		} else {
			newChar, newCell = char.ApplyAction(effects[i], board, i%board.width, i/board.width)
		}
		applied[i] = types.Cell{
			X:          board.Cells[i].X,
			Y:          board.Cells[i].Y,
			DeathCount: newCell.DeathCount,
			Character:  newChar,
		}
	}
	return applied
}

func (board *Board) calculateNextState(applied []types.Cell) ([]types.Cell, []int) {
	next := make([]types.Cell, len(board.Cells))
	nextActiveCells := []int{}
	for y := 0; y < board.height; y++ {
		for x := 0; x < board.width; x++ {
			idx := y*board.width + x
			neighbors := board.countNeighbors(x, y)
			char := applied[idx].Character
			var nextChar types.Character
			var nextCell types.Cell

			if char == nil {
				// Reproduction rule: if 2-3 neighbors, spawn a new cell
				if neighbors >= 2 && neighbors <= 3 {
					// Use a new character for the new life form
					nextChar = &characters.LivingCharacter{ID: nextID(), UnderPop: 2, OverPop: 3, Repro: 3}
				} else {
					nextChar = nil
				}
				nextCell = applied[idx]
			} else {
				// Character exists: handle state and death transition
				nextChar, nextCell = char.NextState(neighbors, board, x, y)
			}

			next[idx] = types.Cell{
				X:          x,
				Y:          y,
				DeathCount: nextCell.DeathCount,
				Character:  nextChar,
			}
			if nextChar != nil {
				nextActiveCells = append(nextActiveCells, idx)
			}
		}
	}
	return next, nextActiveCells
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
	board.updateActiveCells()
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
			board.Cells[idx] = types.Cell{
				X:          x,
				Y:          y,
				DeathCount: 0,
				Character:  nil,
			}
		}
	}
	board.activeCells = []int{}
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
