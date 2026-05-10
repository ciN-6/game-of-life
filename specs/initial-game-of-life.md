# Project: Go-Life (Conway's Game of Life)

A simulation of character life-cycle.

## Game of Life Logic
- **Board Representation:** Use a 2D slice or a flat 1D slice (for performance) of booleans or custom `character` types.
- **Rules:**
  1. Underpopulation: Live character with < 2 live neighbors dies.
  2. Survival: Live character with 2 or 3 live neighbors lives.
  3. Overpopulation: Live character with > 3 live neighbors dies.
  4. Reproduction: Dead character with exactly 3 live neighbors becomes live.
- **Concurrency:** Use Go routines and channels to calculate next-generation updates in parallel if the grid is large.

## UI

create a ui that contains the characters. 
At the start of the game, ask the user to select up to 10 characters.
The selected characters are the seed for this simulation.

UI elements
- Button to start the simulation.
- Button to speed up the simulation.
- button to slow down the simulation.
- button to pause the simulation.
- button to reset the simulation.
- button to clear the simulation.