# Go-Life: Conway’s Game of Life

A performant, extensible implementation of Conway’s Game of Life in Go, featuring a graphical UI (Ebiten), custom character types, and robust test coverage.

## Features

- Pure Go simulation engine (`pkg/life`)
- Graphical UI with Ebiten (`cmd/go-life`)
- Customizable rules and character types (Living, Undead, etc.)
- BDD and unit tests (Godog, table-driven)
- Extensible architecture for new behaviors

## Project Structure

- `cmd/go-life/` — Main entry/UI
- `pkg/life/` — Core simulation logic
- `features/` — BDD scenarios
- `GEMINI.md` — Dev/architecture notes

## Getting Started

### Prerequisites

- Go 1.26+
- Windows, macOS, or Linux

### Build & Run

```powershell
go build -o go-life.exe ./cmd/go-life
./go-life.exe
```

### Run Tests

```powershell
go test ./... -v
```

## Usage

- Select grid size, seed cells, and run/pause/step the simulation.
- Supports custom rules and character types.

## Testing

- Table-driven unit tests: `pkg/life/life_test.go`
- BDD scenarios: `features/simulation.feature`, `pkg/life/life_bdd_test.go`

## Contributing

- Follow Go best practices and code review comments.
- Add godocs for all exported functions.
- Prefer table-driven and BDD tests for new features.
