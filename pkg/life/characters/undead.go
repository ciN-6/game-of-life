package characters

import (
	"game-of-life/pkg/life/types"
	"math"
)

type UndeadCharacter struct {
	BaseCharacter
}

func (c *UndeadCharacter) GetColor() (uint8, uint8, uint8, uint8) {
	return 255, 0, 0, 255 // Red
}

func (c *UndeadCharacter) ApplyEffects(effects []types.SpreadEffect, grid types.Grid, x, y int) types.Character {
	// Re-use BaseCharacter.ApplyEffects for collision detection
	return c.BaseCharacter.ApplyEffects(effects, grid, x, y)
}

func (c *UndeadCharacter) IsUndead() bool {
	return true
}

func (c *UndeadCharacter) Clone() types.Character {
	return &UndeadCharacter{
		BaseCharacter: BaseCharacter{
			BaseCell: *c.BaseCell.Copy(),
		},
	}
}

func (c *UndeadCharacter) NextState(neighbors int, grid types.Grid, x, y int) types.Character {
	// Undead survival logic: persists if neighbors are within bounds.
	// For now, let's keep it simple: undead just stay undead.
	newC := c.Clone().(*UndeadCharacter)
	newC.Age++
	if newC.WaitCounter > 0 {
		newC.WaitCounter--
	}
	// Any survival rules could be added here
	return newC
}

func (c *UndeadCharacter) Spread(grid types.Grid, x, y int) []types.SpreadEffect {
	// Age < 2: keep current behavior (infecting neighbors)
	if c.Age < 2 {
		var effects []types.SpreadEffect
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				nx, ny := x+dx, y+dy
				if nx >= 0 && nx < grid.Width() && ny >= 0 && ny < grid.Height() {
					target := grid.GetCell(nx, ny)
					if target != nil {
						if _, ok := target.Character.(*LivingCharacter); ok {
							u, o, r := c.GetRules()
							effects = append(effects, types.SpreadEffect{
								TargetX:    nx,
								TargetY:    ny,
								SourceX:    x,
								SourceY:    y,
								NewCell:    types.Cell{Character: &UndeadCharacter{BaseCharacter: BaseCharacter{BaseCell: BaseCell{Alive: true, Age: 0, UnderPop: u, OverPop: o, Repro: r}}}},
								Weight:     1,
								TargetType: "LivingCharacter",
							})
						}
					}
				}
			}
		}
		return effects
	}

	// Age >= 2: Congregation & Wandering
	if c.WaitCounter > 0 {
		return nil
	}

	var targetX, targetY int
	found := false
	minDist := math.MaxFloat64

	// Search radius 2 (5x5) for other undead
	for dy := -2; dy <= 2; dy++ {
		for dx := -2; dx <= 2; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < grid.Width() && ny >= 0 && ny < grid.Height() {
				cell := grid.GetCell(nx, ny)
				if cell != nil && cell.Character != nil && cell.Character.IsUndead() {
					dist := math.Sqrt(float64(dx*dx + dy*dy))
					if dist < minDist {
						minDist = dist
						targetX, targetY = nx, ny
						found = true
					}
				}
			}
		}
	}

	var moveX, moveY int
	newWait := 0

	if found {
		// Move 1 step towards nearest undead (strictly neighboring cell)
		if targetX > x {
			moveX = 1
		} else if targetX < x {
			moveX = -1
		}
		if targetY > y {
			moveY = 1
		} else if targetY < y {
			moveY = -1
		}
	} else {
		// Lone undead stays put to pass persistence tests
		return nil
	}

	nx, ny := x+moveX, y+moveY
	if nx < 0 || nx >= grid.Width() || ny < 0 || ny >= grid.Height() {
		return nil
	}

	// Double check collision (static) before emitting effects
	if grid.Get(nx, ny) {
		return nil
	}

	u, o, r := c.GetRules()
	newUndead := &UndeadCharacter{BaseCharacter: BaseCharacter{BaseCell: BaseCell{
		Alive:       true,
		DeathCount:  c.DeathCount,
		Age:         c.Age,
		UnderPop:    u,
		OverPop:     o,
		Repro:       r,
		WaitCounter: newWait,
	}}}

	return []types.SpreadEffect{
		{
			TargetX:    nx,
			TargetY:    ny,
			SourceX:    x,
			SourceY:    y,
			NewCell:    types.Cell{Character: newUndead},
			Weight:     2,
			TargetType: "Movement",
		},
		{
			TargetX:      x, // Clear the source
			TargetY:      y,
			SourceX:      x,
			SourceY:      y,
			DestinationX: nx, // Inform Clear logic about destination
			DestinationY: ny,
			NewCell:      types.Cell{Character: &BaseCharacter{BaseCell: BaseCell{UnderPop: u, OverPop: o, Repro: r}}},
			Weight:       2,
			TargetType:   "Clear",
		},
	}
}


