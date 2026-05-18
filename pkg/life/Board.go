// Package life contains the core simulation logic.
package life

import (
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
	Cells         []Cell
	history       [][]Cell
	activeCells   []int
}

// NewGrid creates a new Board with the given dimensions.
func NewGrid(width, height int) *Board {
	board := &Board{
		width:       width,
		height:      height,
		Cells:       make([]Cell, width*height),
		activeCells: []int{},
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := y*width + x
			board.Cells[idx] = Cell{
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

	if alive {
		// Maintain existing rules if character is already present
		u, o, r := 2, 3, 3
		if char := board.Cells[y*board.width+x].Character; char != nil {
			u, o, r = char.GetRules()
		}
		board.SetCharacter(x, y, &LivingCharacter{ID: nextID(), UnderPop: u, OverPop: o, Repro: r})
	} else {
		board.SetCharacter(x, y, nil)
	}
}

// SetCharacter assigns a character to a specific coordinate and updates the active cell list.
func (board *Board) SetCharacter(x, y int, char Character) {
	if x < 0 || x >= board.width || y < 0 || y >= board.height {
		return
	}
	idx := y*board.width + x
	hadCharacter := board.Cells[idx].Character != nil

	board.Cells[idx].Character = char

	if char != nil {
		if !hadCharacter {
			board.activeCells = append(board.activeCells, idx)
		}
	} else if hadCharacter {
		board.removeFromActiveCells(idx)
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
	_, isLiving := cell.Character.(*LivingCharacter)
	return isLiving
}

// countNeighbors returns the number of live neighbors for a cell at (x, y).
func (board *Board) countNeighbors(x, y int) int {
	count := 0
	board.ForEachNeighbor(x, y, 1, func(nx, ny int) {
		if board.GetAlive(nx, ny) {
			count++
		}
	})
	return count
}

// ForEachNeighbor calls fn for each neighbor of (x, y) within grid bounds within the specified radius.
func (board *Board) ForEachNeighbor(x, y, radius int, fn func(nx, ny int)) {
	for dy := -radius; dy <= radius; dy++ {
		for dx := -radius; dx <= radius; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < board.width && ny >= 0 && ny < board.height {
				fn(nx, ny)
			}
		}
	}
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

func (board *Board) collectIntent() map[int][]SpreadEffect {
	effects := make(map[int][]SpreadEffect)
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

func (board *Board) applyActions(effects map[int][]SpreadEffect) []Cell {
	applied := make([]Cell, len(board.Cells))
	for i := range board.Cells {
		char := board.Cells[i].Character
		var newChar Character
		var newCell Cell

		if char != nil {
			newChar, newCell = char.ApplyAction(effects[i], board, i%board.width, i/board.width)
		} else {
			// Handle effects on empty cells (e.g., reproduction by spread or infection)
			newChar = nil
			newCell = board.Cells[i]
			for _, e := range effects[i] {
				if e.TargetX == i%board.width && e.TargetY == i/board.width {
					newChar = e.NewCell.Character
					newCell = e.NewCell
					break
				}
			}
		}

		applied[i] = Cell{
			X:          board.Cells[i].X,
			Y:          board.Cells[i].Y,
			DeathCount: newCell.DeathCount,
			Character:  newChar,
		}
	}
	return applied
}

func (board *Board) calculateNextState(applied []Cell) ([]Cell, []int) {
	next := make([]Cell, len(board.Cells))
	copy(next, applied)

	candidates := make(map[int]struct{})
	for _, idx := range board.activeCells {
		candidates[idx] = struct{}{}
		cell := board.Cells[idx]
		board.ForEachNeighbor(cell.X, cell.Y, 1, func(nx, ny int) {
			candidates[ny*board.width+nx] = struct{}{}
		})
	}

	nextActiveCells := []int{}
	for idx := range candidates {
		x, y := idx%board.width, idx/board.width
		neighbors := board.countNeighbors(x, y)
		char := applied[idx].Character
		var nextChar Character
		var nextCell Cell

		if char != nil {
			nextChar, nextCell = char.NextState(neighbors, board, x, y)
		} else {
			// Reproduction rule: if neighbors is 2 or 3, spawn a new cell
			if neighbors >= 2 && neighbors <= 3 {
				nextChar = &LivingCharacter{ID: nextID(), UnderPop: 2, OverPop: 3, Repro: 3}
			} else {
				nextChar = nil
			}
			nextCell = applied[idx]
		}

		if nextChar != nil {
			// Assign ID if it's a new LivingCharacter
			if lc, ok := nextChar.(*LivingCharacter); ok && lc.ID == -1 {
				lc.ID = nextID()
			}
			nextActiveCells = append(nextActiveCells, idx)
		}

		next[idx] = Cell{
			X:          x,
			Y:          y,
			DeathCount: nextCell.DeathCount,
			Character:  nextChar,
		}
	}
	return next, nextActiveCells
}

// saveSnapshot saves the current board state to history.
func (board *Board) saveSnapshot() {
	snapshot := make([]Cell, len(board.Cells))
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
func (board *Board) GetCell(x, y int) *Cell {
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
			board.Cells[idx] = Cell{
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
		if _, isLiving := cell.Character.(*LivingCharacter); isLiving {
			count++
		}
	}
	return count
}
