# Copilot Instructions for game-of-life

## Build, Test, and Lint Commands

### Running Tests
```powershell
# Run all tests
go test ./...

# Run tests with verbose output (includes BDD scenario names)
go test ./... -v

# Run tests for a specific package
go test ./pkg/life -v

# Run a specific test
go test ./pkg/life -run TestRules -v
```

The project uses:
- **Unit tests** (table-driven): `pkg/life/life_test.go` - tests Conway's Game of Life rules
- **BDD tests**: `pkg/life/life_bdd_test.go` - Gherkin-based scenario testing using godog framework
- **Feature files**: `features/simulation.feature` - human-readable scenario definitions

### Building the Executable
```powershell
# Build the game
go build -o go-life.exe ./cmd/go-life

# Run the game after building
go run ./cmd/go-life/main.go
```

### Code Formatting
```powershell
# Format all Go files in the project
go fmt ./...

# Format a specific file
go fmt path/to/file.go
```

### Linting
The project follows [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments). Use Go's built-in tooling:
```powershell
# Check for common issues
go vet ./...
```

---

## High-Level Architecture

### Core Simulation Logic vs. UI
The architecture cleanly separates **game logic** from **rendering**:

- **`pkg/life/`** - Pure simulation engine
  - `life.go`: Board management (grid state, neighbor counting, step/iteration)
  - `types/types.go`: Core interfaces (`Grid`, `Character`, `Cell`)
  - `characters/`: Character implementations (BaseCharacter, LivingCharacter, UndeadCharacter)
  
- **`cmd/go-life/main.go`** - Game UI and graphics
  - Uses Ebiten (2D game engine) for rendering
  - Manages game state machine (StateGridSizeSelection → StateSeedSelection → StateRunning/StatePaused)
  - Handles user input and button interactions

This separation makes the core logic easily testable without UI dependencies.

### Character System (Architecture Issue - Needs Refactoring)
The simulation **currently mixes concerns** and needs separation:

**Current (Problematic) Structure:**
- `BaseCell` - holds state (Alive, Age, DeathCount) + rules (UnderPop, OverPop, Repro)
- `BaseCharacter` - wrapper around BaseCell with methods that delegate to it
- `Character` interface - includes state accessors (`GetAge()`, `GetRules()`, `IsAlive()`) mixed with behavior (`NextState()`, `ApplyEffects()`)

**Intended (Needs Implementation):**
- **Cell** - just a container (coordinates + Character reference)
- **Character** - behavior only (NextState, ApplyEffects, Spread, GetColor)
- **State** - separate struct for cell properties (Alive, Age, DeathCount, rules)

**Types Present:**
- `BaseCharacter` - standard Conway's Game of Life behavior
- `LivingCharacter` - extends BaseCharacter, represents alive state
- `UndeadCharacter` - persists across generations, can spawn from certain conditions

**Why This Matters:**
- Character type should not conflate state properties with behavior
- Makes it hard to add new character types without duplicating state logic
- Violates separation of concerns: Grid/Board manages state, Character manages behavior

### Game Flow (Step Function)
The `Board.Step()` method executes one generation:

1. **Count neighbors** for all cells
2. **Apply rules** - each cell delegates to its `Character.NextState()`
3. **Apply effects** - resolve secondary effects like undead spreading
4. Update board state atomically

---

## Key Conventions

### Naming
- Use **whole words** for variable names (e.g., `grid`, not `g`; `neighbors`, not `n`)
- Package-level exports must be documented with GoDoc comments
- Use verb-noun patterns for methods: `CountAlive()`, `ApplyEffects()`, etc.

### Error Handling
- No `panic()` - handle errors explicitly
- Boundary checks (grid bounds) return safe defaults (e.g., `Get()` returns `false` for out-of-bounds)

### Testing Strategy
- **Table-driven tests** for rules validation (see `TestRules` in `life_test.go`)
- **BDD scenarios** for end-to-end behavior (see `features/simulation.feature`)
- Focus on testing the character lifecycle: underpopulation → survival → overpopulation → reproduction

### Board Representation
- **1D array** (`Cells []types.Cell`) storing grid row-by-row
- Index calculation: `idx = y * width + x`
- Cells are indexed as `Cells[y*width+x]`; coordinates are `(x, y)` where x=column, y=row

### Ebiten UI Patterns
- Game state machine drives UI rendering
- Button clicks handled before grid interactions
- Screen coordinates are pixel-based; grid coordinates are cell-based
- Rendering loops through `Board.Cells` and calls `Character.GetColor()`

### Architectural Issue: Character vs. Cell Separation
**Current Problem:** `BaseCell` conflates state with character behavior. The design mixes:
- State properties (Alive, Age, DeathCount, rules) in `BaseCell`
- Behavior methods in `Character` interface that duplicate state accessors
- Character wrapper that just delegates to BaseCell

**What It Should Be:**
1. `Cell` = just a container with coordinates and a Character reference
2. `Character` = behavior interface (NextState, ApplyEffects, Spread, GetColor only)
3. `State` = separate concern (Alive, Age, DeathCount stored at the board or character level, not both)

**Impact:** Makes extending character types awkward and violates separation of concerns. Refactoring needed if adding complex character interactions.

---

## Performance Notes

When optimizing neighbor counting:
- Current implementation loops all 8 neighbors with bounds checking
- Consider spatial patterns (e.g., pre-computing edges) only if profiling shows it's the bottleneck
- Keep the interface stable to avoid breaking character implementations
