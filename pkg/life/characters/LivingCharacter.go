package characters

import (
	"game-of-life/pkg/life/types"
)

// LivingCharacter represents an entity that exhibits fear of undead and flees.
type LivingCharacter struct {
	BaseCharacter
}

func (c *LivingCharacter) GetColor() (uint8, uint8, uint8, uint8) {
	return 0, 255, 255, 255 // Cyan when alive
}

func (c *LivingCharacter) ApplyAction(effects []types.SpreadEffect, grid types.Grid, x, y int) (types.Character, types.Cell) {
	if bestEffect := c.ResolveEffects(effects, grid, x, y); bestEffect != nil {
		return bestEffect.NewCell.Character, bestEffect.NewCell
	}
	cell := *grid.GetCell(x, y)
	return c, cell
}

func (c *LivingCharacter) PrepareAction(grid types.Grid, x, y int) []types.SpreadEffect {
	// Look for Undead neighbors
	hasUndeadNeighbor := false
	types.ForEachNeighbor(grid, x, y, 1, func(nx, ny int) bool {
		if target := grid.GetCell(nx, ny); target != nil && target.Character != nil && target.Character.IsUndead() {
			hasUndeadNeighbor = true
			return false // break
		}
		return true
	})

	if !hasUndeadNeighbor {
		return nil
	}

	// Try to find an empty neighbor to flee to
	var fleeEffect []types.SpreadEffect
	types.ForEachNeighbor(grid, x, y, 1, func(nx, ny int) bool {
		target := grid.GetCell(nx, ny)
		if target != nil {
			// Ensure the destination cell is not adjacent to any undead
			isSafe := true
			types.ForEachNeighbor(grid, nx, ny, 1, func(tx, ty int) bool {
				if nCell := grid.GetCell(tx, ty); nCell != nil && nCell.Character != nil && nCell.Character.IsUndead() {
					isSafe = false
					return false // break
				}
				return true
			})

			if !isSafe {
				return true // Continue to next neighbor
			}

			// Found an empty cell to move to.
			newDeathCount := target.DeathCount

			newCell := types.Cell{
				X:          nx,
				Y:          ny,
				DeathCount: newDeathCount,
				Character:  c,
			}

			fleeEffect = []types.SpreadEffect{
				{
					TargetX:    nx,
					TargetY:    ny,
					NewCell:    newCell,
					Weight:     10,
					TargetType: "EmptyCell",
				},
				{
					TargetX:    x,
					TargetY:    y,
					NewCell:    types.Cell{X: x, Y: y, DeathCount: target.DeathCount, Character: &BaseCharacter{ID: c.ID, UnderPop: c.UnderPop, OverPop: c.OverPop, Repro: c.Repro}},
					Weight:     10,
					TargetType: "LivingCharacter",
				},
			}
			return false // Break iteration
		}
		return true
	})

	if len(fleeEffect) > 0 {
		return fleeEffect
	}
	return nil
}

func (c *LivingCharacter) Clone() types.Character {
	return &LivingCharacter{
		BaseCharacter: BaseCharacter{
			ID:       c.ID,
			UnderPop: c.UnderPop,
			OverPop:  c.OverPop,
			Repro:    c.Repro,
		},
	}
}

func (c *LivingCharacter) NextState(neighbors int, grid types.Grid, x, y int) (types.Character, types.Cell) {
	cell := *grid.GetCell(x, y)

	// Survival logic
	isAlive := neighbors >= c.UnderPop && neighbors <= c.OverPop
	if isAlive {
		cell.DeathCount = 0 
		return c, cell
	}

	// Death logic
	cell.DeathCount++
	if cell.DeathCount >= 5 {
		return &UndeadCharacter{
			BaseCharacter: BaseCharacter{
				ID:       c.ID, 
				UnderPop: c.UnderPop, 
				OverPop:  c.OverPop, 
				Repro:    c.Repro,
			},
			WaitCounter: 0,
		}, cell
	}

	// Become BaseCharacter upon death
	return &BaseCharacter{
		ID:       c.ID, 
		UnderPop: c.UnderPop, 
		OverPop:  c.OverPop, 
		Repro:    c.Repro,
	}, cell
}
