package characters

import "game-of-life/pkg/life/types"

type BaseCell struct {
	Alive       bool
	DeathCount  int
	Age         int
	UnderPop    int
	OverPop     int
	Repro       int
	Highlighted bool
	WaitCounter int
}

func (c *BaseCell) IsAlive() bool { return c.Alive }
func (c *BaseCell) IsUndead() bool { return false }
func (c *BaseCell) GetAge() int   { return c.Age }
func (c *BaseCell) GetDeathCount() int { return c.DeathCount }
func (c *BaseCell) IsHighlighted() bool { return c.Highlighted }
func (c *BaseCell) SetHighlighted(h bool) { c.Highlighted = h }
func (c *BaseCell) GetRules() (int, int, int) { return c.UnderPop, c.OverPop, c.Repro }
func (c *BaseCell) SetRules(u, o, r int) { c.UnderPop = u; c.OverPop = o; c.Repro = r }

func (c *BaseCell) Copy() *BaseCell {
	newC := *c
	return &newC
}

type BaseCharacter struct {
	BaseCell
}

func (c *BaseCharacter) IsAlive() bool { return c.BaseCell.IsAlive() }
func (c *BaseCharacter) IsUndead() bool { return c.BaseCell.IsUndead() }
func (c *BaseCharacter) GetAge() int { return c.BaseCell.GetAge() }
func (c *BaseCharacter) GetDeathCount() int { return c.BaseCell.GetDeathCount() }
func (c *BaseCharacter) IsHighlighted() bool { return c.BaseCell.IsHighlighted() }
func (c *BaseCharacter) SetHighlighted(h bool) { c.BaseCell.SetHighlighted(h) }
func (c *BaseCharacter) GetRules() (int, int, int) { return c.BaseCell.GetRules() }
func (c *BaseCharacter) SetRules(u, o, r int) { c.BaseCell.SetRules(u, o, r) }

func (c *BaseCharacter) GetColor() (uint8, uint8, uint8, uint8) {
	return 50, 50, 50, 255 // DarkGray
}

func (c *BaseCharacter) Clone() types.Character {
	return &BaseCharacter{BaseCell: *c.BaseCell.Copy()}
}

func (c *BaseCharacter) ResolveEffects(effects []types.SpreadEffect, grid types.Grid, x, y int) *types.SpreadEffect {
	var bestEffect *types.SpreadEffect
	for i := range effects {
		e := &effects[i]
		if e.TargetType == "Movement" && c.IsAlive() {
			continue // Collision: ignore movement into occupied cell
		}
		if e.TargetType == "Clear" {
			// Movement source: only clear if destination cell is not blocked
			if grid.Get(e.DestinationX, e.DestinationY) {
				continue // Destination blocked: cancel clearing source
			}
		}

		if bestEffect == nil || e.Weight > bestEffect.Weight {
			bestEffect = e
		}
	}
	return bestEffect
}

func (c *BaseCharacter) ApplyEffects(effects []types.SpreadEffect, grid types.Grid, x, y int) types.Character {
	if bestEffect := c.ResolveEffects(effects, grid, x, y); bestEffect != nil {
		return bestEffect.NewCell.Character
	}
	return c
}

func (c *BaseCharacter) NextState(neighbors int, grid types.Grid, x, y int) types.Character {
	if c.Alive {
		// Survival: 2 or 3 neighbors
		isAlive := neighbors >= c.UnderPop && neighbors <= c.OverPop
		if isAlive {
			newBase := *c
			return &LivingCharacter{BaseCharacter: BaseCharacter{BaseCell: newBase.BaseCell}}
		}
		// Death
		newBase := *c
		newBase.Alive = false
		newBase.DeathCount++
		return &BaseCharacter{BaseCell: newBase.BaseCell}
	}
	// Reproduction: 3 neighbors
	if neighbors == c.Repro {
		newBase := *c
		newBase.Alive = true
		return &LivingCharacter{BaseCharacter: BaseCharacter{BaseCell: newBase.BaseCell}}
	}
	return c
}

func (c *BaseCharacter) Spread(grid types.Grid, x, y int) []types.SpreadEffect {
	return []types.SpreadEffect{}
}

