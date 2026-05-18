package life

import (
	"testing"
)

func TestRules(t *testing.T) {
	tests := []struct {
		name     string
		initial  [][]bool
		expected [][]bool
	}{
		{
			name: "Underpopulation",
			initial: [][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},
			},
			expected: [][]bool{
				{false, false, false},
				{false, false, false},
				{false, false, false},
			},
		},
		{
			name: "Survival (2 neighbors)",
			initial: [][]bool{
				{true, true, false},
				{true, false, false},
				{false, false, false},
			},
			expected: [][]bool{
				{true, true, false},
				{true, true, false},
				{false, false, false},
			},
		},
		{
			name: "Survival (3 neighbors)",
			initial: [][]bool{
				{true, true, false},
				{true, true, false},
				{false, false, false},
			},
			expected: [][]bool{
				{true, true, true},
				{true, true, true},
				{true, true, false},
			},
		},
		{
			name: "Overpopulation",
			initial: [][]bool{
				{true, true, true},
				{true, true, false},
				{false, false, false},
			},
			expected: [][]bool{
				{true, false, true},
				{true, false, true},
				{true, true, false},
			},
		},
		{
			name: "Reproduction",
			initial: [][]bool{
				{true, true, false},
				{true, false, false},
				{false, false, false},
			},
			// (1,1) has 3 neighbors, should become alive
			expected: [][]bool{
				{true, true, false},
				{true, true, false},
				{false, false, false},
			},
		},
		{
			name: "Blinker Oscillator",
			initial: [][]bool{
				{false, true, false},
				{false, true, false},
				{false, true, false},
			},
			expected: [][]bool{
				{true, false, true},
				{true, true, true},
				{true, false, true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Increase grid size to 5x5 to avoid edge issues
			w, h := 5, 5
			g := NewGrid(w, h)
			// Offset the pattern to the center
			offsetX, offsetY := 1, 1
			for y := 0; y < len(tt.initial); y++ {
				for x := 0; x < len(tt.initial[0]); x++ {
					g.SetAlive(x+offsetX, y+offsetY, tt.initial[y][x])
				}
			}

			g.Step()

			for y := 0; y < len(tt.initial); y++ {
				for x := 0; x < len(tt.initial[0]); x++ {
					expected := tt.expected[y][x]
					actual := g.GetAlive(x+offsetX, y+offsetY)
					if actual != expected {
						t.Errorf("%s: at (%d, %d) expected %v, got %v", tt.name, x, y, expected, actual)
					}
				}
			}
		})
	}
}

func TestUndeadEvolution(t *testing.T) {
	g := NewGrid(5, 5)
	// Place a cell that will die immediately and stay dead
	g.SetAlive(2, 2, true)
	// (2,2) has 0 neighbors, will die
	
	for i := 0; i < 5; i++ {
		g.Step()
	}
	
	cell := g.GetCell(2, 2)
	if cell.Character == nil || !cell.Character.IsUndead() {
		t.Errorf("expected undead at (2,2) after 5 steps, got %v", cell.Character)
	}
}

func TestCustomRules(t *testing.T) {
	// Test reproduction with 2 neighbors
	g := NewGrid(3, 3)
	// Initialize character as living so it exists
	g.SetAlive(1, 1, true)
	// (1,1) has 2 neighbors (0,0) and (1,0), should become alive
	g.SetAlive(0, 0, true)
	g.SetAlive(1, 0, true)
	
	// Ensure the rule is 2-3 survival/reproduction
	g.Cells[1*3+1].Character.SetRules(2, 3, 3)
	
	g.Step()
	if !g.GetAlive(1, 1) {
		t.Errorf("expected (1,1) to be alive with 2 neighbors")
	}

	// Test survival with 1 neighbor (should die as UnderPop=2)
	g = NewGrid(3, 3)
	g.SetAlive(1, 1, true)
	g.SetAlive(0, 0, true)
	// (1,1) has 1 neighbor (0,0), should die as UnderPop = 2
	g.Step()
	if g.GetAlive(1, 1) {
		t.Errorf("expected (1,1) to die with 1 neighbor")
	}

	// Test overpopulation with 4 neighbors (should die as OverPop=3)
	g = NewGrid(3, 3)
	g.SetAlive(1, 1, true)
	g.SetAlive(0, 0, true)
	g.SetAlive(1, 0, true)
	g.SetAlive(2, 0, true)
	g.SetAlive(0, 1, true)
	// (1,1) has 4 neighbors, should die as OverPop = 3
	g.Step()
	if g.GetAlive(1, 1) {
		t.Errorf("expected (1,1) to die with 4 neighbors")
	}
}

func TestUndo(t *testing.T) {
	tests := []struct {
		name string
		fn   func(*Board)
	}{
		{
			name: "Undo single step",
			fn: func(g *Board) {
				// Set up initial state
				g.SetAlive(0, 0, true)
				g.SetAlive(1, 0, true)
				g.SetAlive(2, 0, true)

				// Verify initial state
				if !g.GetAlive(0, 0) || !g.GetAlive(1, 0) || !g.GetAlive(2, 0) {
					t.Errorf("initial state not set correctly")
				}

				// Step the simulation
				g.Step()

				// Undo
				g.Undo()

				// Verify state is restored
				if !g.GetAlive(0, 0) || !g.GetAlive(1, 0) || !g.GetAlive(2, 0) {
					t.Errorf("undo failed to restore state")
				}
				if g.CountAlive() != 3 {
					t.Errorf("expected 3 alive cells after undo, got %d", g.CountAlive())
				}
			},
		},
		{
			name: "Undo multiple steps",
			fn: func(g *Board) {
				// Set up initial state
				g.SetAlive(0, 0, true)
				g.SetAlive(1, 0, true)
				g.SetAlive(2, 0, true)

				initialAlive := g.CountAlive()

				// Take 3 steps
				g.Step()
				g.Step()
				g.Step()

				// Undo 3 times
				g.Undo()
				g.Undo()
				g.Undo()

				// Verify state is restored to initial
				if g.CountAlive() != initialAlive {
					t.Errorf("expected %d alive cells after 3 undos, got %d", initialAlive, g.CountAlive())
				}
				if !g.GetAlive(0, 0) || !g.GetAlive(1, 0) || !g.GetAlive(2, 0) {
					t.Errorf("undo multiple steps failed to restore state")
				}
			},
		},
		{
			name: "Undo with no history",
			fn: func(g *Board) {
				// Set up initial state
				g.SetAlive(0, 0, true)

				initialAlive := g.CountAlive()

				// Undo without any steps (should do nothing)
				g.Undo()

				// State should be unchanged
				if g.CountAlive() != initialAlive {
					t.Errorf("undo with no history modified state")
				}
			},
		},
		{
			name: "History respects max size",
			fn: func(g *Board) {
				// Set up initial state
				g.SetAlive(0, 0, true)

				// Take MaxHistory + 5 steps
				for i := 0; i < MaxHistory+5; i++ {
					g.Step()
				}

				// Should only be able to undo MaxHistory steps
				for i := 0; i < MaxHistory; i++ {
					g.Undo()
				}

				// One more undo should do nothing (history exhausted)
				aliveBeforeLastUndo := g.CountAlive()
				g.Undo()
				if g.CountAlive() != aliveBeforeLastUndo {
					t.Errorf("history exceeded MaxHistory limit")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGrid(10, 10)
			tt.fn(g)
		})
	}
}
