# Plan: Introduce character Types and Consolidate Rules

## Overview
Refactor the current `[]bool` grid representation to support extensible character types and move game rules (Under/Overpopulation, Reproduction) into the character definition itself. Also, remove all user-facing rule controls.

## Objectives
- [ ] Define extensible `character` struct with rule attributes.
- [ ] Update `pkg/life` to manage a slice of `character` instead of `bool`.
- [ ] Remove UI rule controls in `cmd/go-life/main.go`.
- [ ] Ensure simulation logic uses character-internal rules.

## Tasks
- [ ] Create `pkg/life/character.go` with `character` type containing `UnderPop`, `OverPop`, and `Repro` attributes.
- [ ] Update `Grid` struct in `pkg/life/life.go` to use `characters []character`.
- [ ] Refactor `Set`, `Get`, and `CountAlive` in `life.go`.
- [ ] Update simulation logic in `Step()` to use individual character rules.
- [ ] Remove rule input controls and `drawParameterInput` in `cmd/go-life/main.go`.

### Task Detail
1. **Define `character` struct:**
   ```go
   type character struct {
       Alive bool
       // Rules
       UnderPop int
       OverPop  int
       Repro    int
   }
   ```
2. **Refactor `Grid`**: Remove rule attributes from `Grid`, move them to each `character`.
3. **UI Cleanup**: Remove rule interaction buttons and labels in `main.go`.

## Definition of Done
- [ ] Rules are natively supported by `character` structure.
- [ ] Rule controls are removed from the UI.
- [ ] Simulation behavior remains consistent.
