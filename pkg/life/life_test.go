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
				{true, true, false},
				{true, true, false},
				{false, false, false},
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
				{false, false, false},
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
				{false, false, false},
				{true, true, true},
				{false, false, false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := len(tt.initial)
			w := len(tt.initial[0])
			g := NewGrid(w, h)
			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					g.SetAlive(x, y, tt.initial[y][x])
				}
			}

			g.Step()

			for y := 0; y < h; y++ {
				for x := 0; x < w; x++ {
					if g.GetAlive(x, y) != tt.expected[y][x] {
						t.Errorf("%s: at (%d, %d) expected %v, got %v", tt.name, x, y, tt.expected[y][x], g.GetAlive(x, y))
					}
				}
			}
		})
	}
}

func TestCustomRules(t *testing.T) {
	// Test reproduction with 2 neighbors instead of 3
	g := NewGrid(3, 3)
	// Apply custom rule to the character at (1,1)
	g.Cells[1*3+1].Character.SetRules(2, 3, 2)
	g.SetAlive(0, 0, true)
	g.SetAlive(1, 0, true)
	g.SetAlive(1, 1, false) // Ensure it's not alive
	// (1,1) has 2 neighbors, should become alive because Repro = 2
	g.Step()
	if !g.GetAlive(1, 1) {
		t.Errorf("expected (1,1) to be alive with 2 neighbors when Repro=2")
	}

	// Test survival with 1 neighbor
	g = NewGrid(3, 3)
	// Apply custom rule to the character at (1,1)
	g.Cells[1*3+1].Character.SetRules(1, 3, 3)
	g.SetAlive(0, 0, true)
	g.SetAlive(1, 1, true)
	// (1,1) has 1 neighbor (0,0), should survive because UnderPop = 1
	g.Step()
	if !g.GetAlive(1, 1) {
		t.Errorf("expected (1,1) to survive with 1 neighbor when UnderPop=1")
	}

	// Test overpopulation with 4 neighbors (survival)
	g = NewGrid(3, 3)
	// Apply custom rule to the character at (1,1)
	g.Cells[1*3+1].Character.SetRules(2, 4, 3)
	g.SetAlive(1, 1, true)
	g.SetAlive(0, 0, true)
	g.SetAlive(1, 0, true)
	g.SetAlive(2, 0, true)
	g.SetAlive(0, 1, true)
	// (1,1) has 4 neighbors, should survive because OverPop = 4
	g.Step()
	if !g.GetAlive(1, 1) {
		t.Errorf("expected (1,1) to survive with 4 neighbors when OverPop=4")
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
