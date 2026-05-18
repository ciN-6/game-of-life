package life

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

func (c *UndeadCharacter) Clone() Character {
	return &UndeadCharacter{
		ID:       c.ID,
		UnderPop: c.UnderPop,
		OverPop:  c.OverPop,
		Repro:    c.Repro,
		Age:      c.Age,
	}
}

func (c *UndeadCharacter) ApplyAction(effects []SpreadEffect, grid Grid, x, y int) (Character, Cell) {
	for _, e := range effects {
		if e.TargetX == x && e.TargetY == y {
			return e.NewCell.Character, e.NewCell
		}
	}
	return c, *grid.GetCell(x, y)
}

func (c *UndeadCharacter) NextState(neighbors int, grid Grid, x, y int) (Character, Cell) {
	cell := *grid.GetCell(x, y)

	// Undead characters are persistent and do not change state
	// or age based on standard neighbor rules.
	return c, cell
}

func (c *UndeadCharacter) PrepareAction(grid Grid, x, y int) []SpreadEffect {
	// Infect neighbors logic
	var effects []SpreadEffect
	// Note: ForEachNeighbor is still a global function in Grid.go for now
	ForEachNeighbor(grid, x, y, 1, func(nx, ny int) {
		target := grid.GetCell(nx, ny)
		if target != nil {
			if _, ok := target.Character.(*LivingCharacter); ok {
				victimID := target.Character.GetID()

				newCell := Cell{
					X:          nx,
					Y:          ny,
					DeathCount: 0,
					Character:  &UndeadCharacter{ID: victimID, UnderPop: c.UnderPop, OverPop: c.OverPop, Repro: c.Repro},
				}
				effects = append(effects, SpreadEffect{
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
	})
	return effects
}
