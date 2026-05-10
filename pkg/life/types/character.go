// Package types defines shared simulation interfaces and structures.
package types

// Character defines the simulation behavior for entities on the board.
type Character interface {
	// NextState calculates the character's state for the next generation based on neighbors.
	NextState(neighbors int, grid Grid, x, y int) Character
	// ApplyEffects resolves potential state changes from surrounding characters.
	ApplyEffects(effects []SpreadEffect, grid Grid, x, y int) Character
	// Spread generates effects to propagate to adjacent characters.
	Spread(grid Grid, x, y int) []SpreadEffect
	// IsAlive returns whether the character is currently alive.
	IsAlive() bool
	// IsUndead returns whether the character is undead.
	IsUndead() bool
	// GetAge returns the character's age.
	GetAge() int
	// GetRules returns the character's rule set (UnderPop, OverPop, Repro).
	GetRules() (int, int, int)
	// SetRules updates the character's rules.
	SetRules(u, o, r int)
	// GetDeathCount returns the character's accumulated death count.
	GetDeathCount() int
	// IsHighlighted returns the highlight status for rendering.
	IsHighlighted() bool
	// SetHighlighted updates the highlight status.
	SetHighlighted(bool)
	// GetColor returns the RGBA color components for rendering.
	GetColor() (r, g, b, a uint8)
	// Clone returns a deep copy of the character, preserving its specific type and state.
	Clone() Character
}
