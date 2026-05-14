package life_test

import (
	"errors"
	"fmt"
	"testing"

	"game-of-life/pkg/life"
	"game-of-life/pkg/life/characters"
	"game-of-life/pkg/life/types"

	"github.com/cucumber/godog"
)

type simulationContext struct {
	grid *life.Board
}

func (c *simulationContext) aGridWithASingleLivingcharacterAt(x, y int) error {
	c.grid = life.NewGrid(10, 10)
	c.grid.SetAlive(x, y, true)
	return nil
}

func (c *simulationContext) theSimulationStepsTimes(steps int) error {
	for i := 0; i < steps; i++ {
		c.grid.Step()
	}
	return nil
}

func (c *simulationContext) theSimulationStepsOnce() error {
	return c.theSimulationStepsTimes(1)
}

func (c *simulationContext) thecharacterAtShouldBeDead(x, y int) error {
	if c.grid.GetAlive(x, y) {
		return errors.New("expected character to be dead, but it is alive")
	}
	return nil
}

func (c *simulationContext) aGridWithLivingcharactersAt(table *godog.Table) error {
	c.grid = life.NewGrid(10, 10)
	for _, row := range table.Rows[1:] {
		x := 0
		y := 0
		fmt.Sscanf(row.Cells[0].Value, "%d", &x)
		fmt.Sscanf(row.Cells[1].Value, "%d", &y)
		c.grid.SetAlive(x, y, true)
	}
	return nil
}

func (c *simulationContext) theFollowingcharactersShouldBeAlive(table *godog.Table) error {
	for _, row := range table.Rows[1:] {
		x := 0
		y := 0
		fmt.Sscanf(row.Cells[0].Value, "%d", &x)
		fmt.Sscanf(row.Cells[1].Value, "%d", &y)
		if !c.grid.GetAlive(x, y) {
			return fmt.Errorf("expected character (%d,%d) to be alive", x, y)
		}
	}
	return nil
}

func (c *simulationContext) theFollowingcharactersShouldBeDead(table *godog.Table) error {
	for _, row := range table.Rows[1:] {
		x := 0
		y := 0
		fmt.Sscanf(row.Cells[0].Value, "%d", &x)
		fmt.Sscanf(row.Cells[1].Value, "%d", &y)
		if c.grid.GetAlive(x, y) {
			return fmt.Errorf("expected character (%d,%d) to be dead", x, y)
		}
	}
	return nil
}

func InitializeScenario(sc *godog.ScenarioContext) {
	ctx := &simulationContext{}

	sc.Step(`^a grid with a single living character at (\d+), (\d+)$`, ctx.aGridWithASingleLivingcharacterAt)
	sc.Step(`^a grid with an undead character at (\d+), (\d+)$`, ctx.aGridWithAnUndeadcharacterAt)
	sc.Step(`^the simulation steps (\d+) times$`, ctx.theSimulationStepsTimes)
	sc.Step(`^the simulation steps once$`, ctx.theSimulationStepsOnce)
	sc.Step(`^the character at (\d+), (\d+) should be dead$`, ctx.thecharacterAtShouldBeDead)
	sc.Step(`^the character at (\d+), (\d+) should be undead$`, ctx.thecharacterAtShouldBeUndead)
	sc.Step(`^a grid with living characters at:$`, ctx.aGridWithLivingcharactersAt)
	sc.Step(`^the following characters should be alive:$`, ctx.theFollowingcharactersShouldBeAlive)
	sc.Step(`^the following characters should be dead:$`, ctx.theFollowingcharactersShouldBeDead)
}

func (c *simulationContext) aGridWithAnUndeadcharacterAt(x, y int) error {
	c.grid = life.NewGrid(10, 10)
	undead := &characters.UndeadCharacter{ID: 1, UnderPop: 2, OverPop: 3, Repro: 3, WaitCounter: 0}
	c.grid.Cells[y*10+x] = types.Cell{X: x, Y: y, DeathCount: 0, Character: undead}
	return nil
}

func (c *simulationContext) thecharacterAtShouldBeUndead(x, y int) error {
	cell := c.grid.Cells[y*10+x]
	if cell.Character == nil {
		return fmt.Errorf("no character at (%d,%d)", x, y)
	}
	char := cell.Character
	if !char.IsUndead() {
		return fmt.Errorf("expected undead at (%d,%d), but got type %T (IsUndead: %v)", x, y, char, char.IsUndead())
	}
	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("failed to run feature tests")
	}
}
