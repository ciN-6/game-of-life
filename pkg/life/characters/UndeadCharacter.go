package characters

import (
	"game-of-life/pkg/life/types"
)

// UndeadCharacter represents an undead entity that persists and spreads across generations.
type UndeadCharacter struct {
	ID       int
	UnderPop int
	OverPop  int
	Repro    int
	Age      int
}

func (c *UndeadCharacter) GetID() int                             { return c.ID }
func (c *UndeadCharacter) IsUndead() bool                         { return true }
func (c *UndeadCharacter) GetRules() (int, int, int)              { return c.UnderPop, c.OverPop, c.Repro }
func (c *UndeadCharacter) SetRules(u, o, r int)                   { c.UnderPop = u; c.OverPop = o; c.Repro = r }
func (c *UndeadCharacter) GetColor() (uint8, uint8, uint8, uint8) { return 255, 0, 0, 255 }

func (c *UndeadCharacter) Clone() types.Character {
	return &UndeadCharacter{
		ID:       c.ID,
		UnderPop: c.UnderPop,
		OverPop:  c.OverPop,
		Repro:    c.Repro,
		Age:      c.Age,
	}
}

func (c *UndeadCharacter) ApplyAction(effects []types.SpreadEffect, grid types.Grid, x, y int) (types.Character, types.Cell) {
	for _, e := range effects {
		if e.TargetX == x && e.TargetY == y {
			return e.NewCell.Character, e.NewCell
		}
	}
	return c, *grid.GetCell(x, y)
}

func (c *UndeadCharacter) NextState(neighbors int, grid types.Grid, x, y int) (types.Character, types.Cell) {
	cell := *grid.GetCell(x, y)

	newWait := c.Age - 1
	if newWait < 0 {
		newWait = 0
	}

	return &UndeadCharacter{
		ID:       c.ID,
		UnderPop: c.UnderPop,
		OverPop:  c.OverPop,
		Repro:    c.Repro,
		Age:      newWait,
	}, cell
}

func (c *UndeadCharacter) PrepareAction(grid types.Grid, x, y int) []types.SpreadEffect {
	// Infect neighbors logic
	var effects []types.SpreadEffect
	types.ForEachNeighbor(grid, x, y, 1, func(nx, ny int) bool {
		target := grid.GetCell(nx, ny)
		if target != nil {
			if _, ok := target.Character.(*LivingCharacter); ok {
				victimID := target.Character.GetID()

				newCell := types.Cell{
					X:          nx,
					Y:          ny,
					DeathCount: 0,
					Character:  &UndeadCharacter{ID: victimID, UnderPop: c.UnderPop, OverPop: c.OverPop, Repro: c.Repro},
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
