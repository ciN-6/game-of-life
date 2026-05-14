package characters

import (
	"game-of-life/pkg/life/types"
)

// LivingCharacter represents an entity that exhibits fear of undead and flees.
type LivingCharacter struct {
	ID       int
	UnderPop int
	OverPop  int
	Repro    int
}

func (c *LivingCharacter) GetColor() (uint8, uint8, uint8, uint8) {
	return 0, 255, 255, 255 // Cyan when alive
}

func (c *LivingCharacter) PrepareAction(grid types.Grid, x, y int) []types.SpreadEffect {
	// Simple implementation
	return nil
}
func (c *LivingCharacter) GetID() int                { return c.ID }
func (c *LivingCharacter) IsUndead() bool            { return false }
func (c *LivingCharacter) GetRules() (int, int, int) { return c.UnderPop, c.OverPop, c.Repro }
func (c *LivingCharacter) SetRules(u, o, r int)      { c.UnderPop = u; c.OverPop = o; c.Repro = r }

func (c *LivingCharacter) ApplyAction(effects []types.SpreadEffect, grid types.Grid, x, y int) (types.Character, types.Cell) {
	// If fleeing, apply the effect
	for _, e := range effects {
		if e.TargetX == x && e.TargetY == y {
			return e.NewCell.Character, e.NewCell
		}
	}
	return c, *grid.GetCell(x, y)
}

func (c *LivingCharacter) Clone() types.Character {
	return &LivingCharacter{
		ID:       c.ID,
		UnderPop: c.UnderPop,
		OverPop:  c.OverPop,
		Repro:    c.Repro,
	}
}

func (c *LivingCharacter) NextState(neighbors int, grid types.Grid, x, y int) (types.Character, types.Cell) {
	cell := *grid.GetCell(x, y)

	// Survival logic: 2-3 neighbors
	if neighbors >= 2 && neighbors <= 3 {
		cell.DeathCount = 0
		return c, cell
	}

	// Death logic (Underpopulation: <=1, Overpopulation: >=4)
	cell.DeathCount++
	if cell.DeathCount >= 5 {
		return &UndeadCharacter{
			ID:       c.ID,
			UnderPop: 2,
			OverPop:  3,
			Repro:    3,
			Age:      0,
		}, cell
	}

	// Return nil for "dead" cell
	cell.Character = nil
	return nil, cell
}
