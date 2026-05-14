package characters

import (
	"game-of-life/pkg/life/types"
)

// BaseCharacter represents the standard simulation behavior.
type BaseCharacter struct {
	ID       int
	UnderPop int
	OverPop  int
	Repro    int
}

func (c *BaseCharacter) GetID() int {
	return c.ID
}

func (c *BaseCharacter) IsUndead() bool {
	return false
}

func (c *BaseCharacter) GetRules() (int, int, int) {
	return c.UnderPop, c.OverPop, c.Repro
}

func (c *BaseCharacter) SetRules(u, o, r int) {
	c.UnderPop = u
	c.OverPop = o
	c.Repro = r
}

func (c *BaseCharacter) GetColor() (uint8, uint8, uint8, uint8) {
	return 20, 20, 20, 255 // Darker gray for dead
}

func (c *BaseCharacter) Clone() types.Character {
	return &BaseCharacter{
		ID:       c.ID,
		UnderPop: c.UnderPop,
		OverPop:  c.OverPop,
		Repro:    c.Repro,
	}
}

func (c *BaseCharacter) ResolveEffects(effects []types.SpreadEffect, grid types.Grid, x, y int) *types.SpreadEffect {
	var bestEffect *types.SpreadEffect
	for i := range effects {
		e := &effects[i]
		
		// Note: The logic here relies on GetAlive, which we will need to refactor. 
		// For now, I will keep the skeleton.
		
		if bestEffect == nil || e.Weight > bestEffect.Weight {
			bestEffect = e
		}
	}
	return bestEffect
}

func (c *BaseCharacter) ApplyAction(effects []types.SpreadEffect, grid types.Grid, x, y int) (types.Character, types.Cell) {
	if bestEffect := c.ResolveEffects(effects, grid, x, y); bestEffect != nil {
		return bestEffect.NewCell.Character, bestEffect.NewCell
	}
	cell := *grid.GetCell(x, y)
	return c, cell
}

func (c *BaseCharacter) NextState(neighbors int, grid types.Grid, x, y int) (types.Character, types.Cell) {
	cell := *grid.GetCell(x, y)
	
	// Reproduction logic: if exactly 3 neighbors, become alive
	if neighbors == c.Repro {
		return &LivingCharacter{BaseCharacter: *c}, cell
	}

	// Persistent death state handling
	cell.DeathCount++
	if cell.DeathCount >= 5 {
		return &UndeadCharacter{BaseCharacter: *c, WaitCounter: 0}, cell
	}
	
	return c, cell
}

func (c *BaseCharacter) PrepareAction(grid types.Grid, x, y int) []types.SpreadEffect {
	return []types.SpreadEffect{}
}
