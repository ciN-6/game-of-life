// Package types defines shared simulation interfaces and structures.
package life

// Character defines the simulation behavior for entities on the board.
type Character interface {
	// GetID returns the unique identifier of the character.
	GetID() int
	// NextState calculates the next generation based on neighbor count and cell info.
	NextState(neighbors int, board *Board, x, y int) (Character, Cell)
	// ApplyAction resolves potential state changes from surrounding characters' effects.
	ApplyAction(effects []SpreadEffect, board *Board, x, y int) (Character, Cell)
	// PrepareAction generates effects to propagate to adjacent cells.
	PrepareAction(board *Board, x, y int) []SpreadEffect
	// IsUndead returns whether this character type is undead (persists across generations).
	IsUndead() bool
	// GetRules returns the character's rule set (UnderPop, OverPop, Repro).
	GetRules() (int, int, int)
	// SetRules updates the character's rules.
	SetRules(u, o, r int)
	// GetColor returns the color of the character.
	GetColor() Color
	// Clone returns a deep copy of the character, preserving its specific type and rules.
	Clone() Character
}
