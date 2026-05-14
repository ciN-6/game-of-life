package characters

import (
	"game-of-life/pkg/life/types"
)

// UndeadCharacter represents an undead entity that persists and spreads across generations.
type UndeadCharacter struct {
	BaseCharacter
	WaitCounter int
}

func (c *UndeadCharacter) GetColor() (uint8, uint8, uint8, uint8) {
	return 255, 0, 0, 255 // Red
}

func (c *UndeadCharacter) ApplyAction(effects []types.SpreadEffect, grid types.Grid, x, y int) (types.Character, types.Cell) {
	if bestEffect := c.ResolveEffects(effects, grid, x, y); bestEffect != nil {
		return bestEffect.NewCell.Character, bestEffect.NewCell
	}
	cell := *grid.GetCell(x, y)
	return c, cell
}

func (c *UndeadCharacter) IsUndead() bool {
	return true
}

func (c *UndeadCharacter) Clone() types.Character {
	return &UndeadCharacter{
		BaseCharacter: BaseCharacter{
			ID:       c.ID,
			UnderPop: c.UnderPop,
			OverPop:  c.OverPop,
			Repro:    c.Repro,
		},
		WaitCounter: c.WaitCounter,
	}
}

func (c *UndeadCharacter) NextState(neighbors int, grid types.Grid, x, y int) (types.Character, types.Cell) {
	cell := *grid.GetCell(x, y)
	
	newUndead := &UndeadCharacter{
		BaseCharacter: BaseCharacter{
			ID:       c.ID,
			UnderPop: c.UnderPop,
			OverPop:  c.OverPop,
			Repro:    c.Repro,
		},
		WaitCounter: c.WaitCounter - 1,
	}
	if newUndead.WaitCounter < 0 {
		newUndead.WaitCounter = 0
	}

	return newUndead, cell
}

func (c *UndeadCharacter) PrepareAction(grid types.Grid, x, y int) []types.SpreadEffect {
	// Infect neighbors logic
	var effects []types.SpreadEffect
	types.ForEachNeighbor(grid, x, y, 1, func(nx, ny int) bool {
		target := grid.GetCell(nx, ny)
		if target != nil {
			if _, ok := target.Character.(*LivingCharacter); ok {
				victimID := target.Character.GetID()
				u, o, r := c.GetRules()
				newCell := types.Cell{
					X:          nx,
					Y:          ny,
					DeathCount: 0,
					Character:  &UndeadCharacter{BaseCharacter: BaseCharacter{ID: victimID, UnderPop: u, OverPop: o, Repro: r}},
				}
				effects = append(effects, types.SpreadEffect{
					TargetX:    nx,
					TargetY:    ny,
					SourceX:    x,
					SourceY:    y,
					NewCell:    newCell,
					Weight:     1,
					TargetType: "LivingCharacter",
				})
			}
		}
		return true
	})
	return effects
}
