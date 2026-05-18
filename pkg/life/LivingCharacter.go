package life

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

func (c *LivingCharacter) PrepareAction(board *Board, x, y int) []SpreadEffect {
	// Simple implementation
	return nil
}
func (c *LivingCharacter) GetID() int                { return c.ID }
func (c *LivingCharacter) IsUndead() bool            { return false }
func (c *LivingCharacter) GetRules() (int, int, int) { return c.UnderPop, c.OverPop, c.Repro }
func (c *LivingCharacter) SetRules(u, o, r int)      { c.UnderPop = u; c.OverPop = o; c.Repro = r }

func (c *LivingCharacter) ApplyAction(effects []SpreadEffect, board *Board, x, y int) (Character, Cell) {
	// If fleeing, apply the effect
	for _, e := range effects {
		if e.TargetX == x && e.TargetY == y {
			return e.NewCell.Character, e.NewCell
		}
	}
	return c, *board.GetCell(x, y)
}

func (c *LivingCharacter) Clone() Character {
	return &LivingCharacter{
		ID:       c.ID,
		UnderPop: c.UnderPop,
		OverPop:  c.OverPop,
		Repro:    c.Repro,
	}
}

func (c *LivingCharacter) NextState(neighbors int, board *Board, x, y int) (Character, Cell) {
	cell := *board.GetCell(x, y)

	// Survival logic: Lives if neighbors are within the UnderPop and OverPop bounds.
	if neighbors >= c.UnderPop && neighbors <= c.OverPop {
		return c, cell
	}

	// Death logic (Underpopulation: <=1, Overpopulation: >=4)
	cell.DeathCount++
	if cell.DeathCount >= 5 {
		return &UndeadCharacter{
			ID:       c.ID,
			UnderPop: c.UnderPop,
			OverPop:  c.OverPop,
			Repro:    c.Repro,
			Age:      5,
		}, cell
	}

	// Return nil for "dead" cell
	return nil, cell
}
