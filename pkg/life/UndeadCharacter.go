package life

// UndeadCharacter represents an undead entity that persists and spreads across generations.
type UndeadCharacter struct {
	ID       int
	UnderPop int
	OverPop  int
	Repro    int
	Age      int
}

func (c *UndeadCharacter) GetID() int                { return c.ID }
func (c *UndeadCharacter) IsUndead() bool            { return true }
func (c *UndeadCharacter) GetRules() (int, int, int) { return c.UnderPop, c.OverPop, c.Repro }
func (c *UndeadCharacter) SetRules(u, o, r int)      { c.UnderPop = u; c.OverPop = o; c.Repro = r }

func (c *UndeadCharacter) GetColor() Color {
	startColor := Color{100, 0, 0, 255} // dark-red
	targetColor := Color{255, 0, 0, 255} // bright-red
	maxAge := 50.0
	factor := float64(c.Age) / maxAge
	if factor > 1.0 {
		factor = 1.0
	}
	return startColor.Interpolate(targetColor, factor)
}

func (c *UndeadCharacter) Clone() Character {
	return &UndeadCharacter{
		ID:       c.ID,
		UnderPop: c.UnderPop,
		OverPop:  c.OverPop,
		Repro:    c.Repro,
		Age:      c.Age,
	}
}

func (c *UndeadCharacter) ApplyAction(effects []SpreadEffect, board *Board, x, y int) (Character, Cell) {
	for _, e := range effects {
		if e.TargetX == x && e.TargetY == y {
			return e.NewCell.Character, e.NewCell
		}
	}
	return c, *board.GetCell(x, y)
}

func (c *UndeadCharacter) NextState(neighbors int, board *Board, x, y int) (Character, Cell) {
	cell := *board.GetCell(x, y)
	c.Age++
	// Undead characters are persistent and do not change state based on neighbors,
	// but they do age.
	return c, cell
}

func (c *UndeadCharacter) PrepareAction(board *Board, x, y int) []SpreadEffect {
	// Infect neighbors logic
	var effects []SpreadEffect
	board.ForEachNeighbor(x, y, 1, func(nx, ny int) {
		target := board.GetCell(nx, ny)
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
