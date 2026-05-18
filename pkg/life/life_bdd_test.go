package life_test

import (
	"errors"
	"fmt"
	"testing"

	"game-of-life/pkg/life"

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
	sc.Step(`^the cell at (\d+), (\d+) should have (\d+) neighbors$`, ctx.theCellAtShouldHaveNeighbors)
}

func (c *simulationContext) theCellAtShouldHaveNeighbors(x, y, expected int) error {
	// We need to use unexported countNeighbors or just ForEachNeighbor
	// Since countNeighbors is unexported, we'll re-implement it or use a trick
	// Wait, countNeighbors IS in the same package (life_test) but unexported?
	// No, life_bdd_test.go is package life_test.
	// But Board.go is package life.
	// So unexported methods are not visible.
	
	count := 0
	c.grid.ForEachNeighbor(x, y, 1, func(nx, ny int) {
		if c.grid.GetAlive(nx, ny) {
			count++
		}
	})
	
	if count != expected {
		return fmt.Errorf("expected %d neighbors at (%d,%d), but got %d", expected, x, y, count)
	}
	return nil
}

func (c *simulationContext) aGridWithAnUndeadcharacterAt(x, y int) error {
	if c.grid == nil {
		c.grid = life.NewGrid(10, 10)
	}
	undead := &life.UndeadCharacter{ID: 1, UnderPop: 2, OverPop: 3, Repro: 3, Age: 0}
	c.grid.SetCharacter(x, y, undead)
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
