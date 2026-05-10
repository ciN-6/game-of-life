package characters

import "game-of-life/pkg/life/types"

type LivingCharacter struct {
	BaseCharacter
}

func (c *LivingCharacter) GetColor() (uint8, uint8, uint8, uint8) {
	return 0, 255, 255, 255 // Cyan
}

func (c *LivingCharacter) GetDeathCount() int { return c.DeathCount }

func (c *LivingCharacter) ApplyEffects(effects []types.SpreadEffect, grid types.Grid, x, y int) types.Character {
	if bestEffect := c.ResolveEffects(effects, grid, x, y); bestEffect != nil {
		return bestEffect.NewCell.Character
	}

	neighbors := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if grid.Get(x+dx, y+dy) {
				neighbors++
			}
		}
	}
	return c.NextState(neighbors, grid, x, y)
}

func (c *LivingCharacter) Spread(grid types.Grid, x, y int) []types.SpreadEffect {
	// Look for Undead neighbors
	hasUndeadNeighbor := false
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if target := grid.GetCell(nx, ny); target != nil && target.Character.IsUndead() {
				hasUndeadNeighbor = true
				break
			}
		}
	}

	if !hasUndeadNeighbor {
		return nil
	}

	// Try to find an empty neighbor to flee to
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < grid.Width() && ny >= 0 && ny < grid.Height() {
				target := grid.GetCell(nx, ny)
				if target != nil && !target.Character.IsAlive() {
					// Found an empty cell to move to.
					newDeathCount := c.DeathCount
					if target.Character.GetDeathCount() > newDeathCount {
						newDeathCount = target.Character.GetDeathCount()
					}

					newChar := &LivingCharacter{BaseCharacter: BaseCharacter{BaseCell: *c.BaseCell.Copy()}}
					newChar.DeathCount = newDeathCount
					newChar.Highlighted = true

					return []types.SpreadEffect{
						{
							TargetX:    nx,
							TargetY:    ny,
							NewCell:    types.Cell{Character: newChar},
							Weight:     10,
							TargetType: "EmptyCell",
						},
						{
							TargetX:    x,
							TargetY:    y,
							NewCell:    types.Cell{Character: &BaseCharacter{BaseCell: BaseCell{UnderPop: c.UnderPop, OverPop: c.OverPop, Repro: c.Repro}}},
							Weight:     10,
							TargetType: "LivingCharacter",
						},
					}
				}
			}
		}
	}
	return nil
}

func (c *LivingCharacter) Clone() types.Character {
	return &LivingCharacter{
		BaseCharacter: BaseCharacter{
			BaseCell: *c.BaseCell.Copy(),
		},
	}
}

func (c *LivingCharacter) NextState(neighbors int, grid types.Grid, x, y int) types.Character {
	if c.Alive {
		// Survival
		isAlive := neighbors >= c.UnderPop && neighbors <= c.OverPop
		if isAlive {
			newChar := c.Clone().(*LivingCharacter)
			newChar.Alive = true
			return newChar
		}
		// Death
		deathCount := c.DeathCount + 1
		if deathCount >= 5 {
			return &UndeadCharacter{BaseCharacter: BaseCharacter{BaseCell: BaseCell{Alive: true, Age: 0, UnderPop: c.UnderPop, OverPop: c.OverPop, Repro: c.Repro}}}
		}
		// Return BaseCharacter as a dead character
		return &BaseCharacter{BaseCell: BaseCell{Alive: false, DeathCount: deathCount, UnderPop: c.UnderPop, OverPop: c.OverPop, Repro: c.Repro}}
	}
	// Reproduction
	if neighbors == c.Repro {
		newChar := c.Clone().(*LivingCharacter)
		newChar.Alive = true
		newChar.Age = 0
		newChar.DeathCount = 0
		return newChar
	}
	// Return BaseCharacter as a dead character
	return &BaseCharacter{BaseCell: BaseCell{Alive: false, DeathCount: c.DeathCount, UnderPop: c.UnderPop, OverPop: c.OverPop, Repro: c.Repro}}
}
