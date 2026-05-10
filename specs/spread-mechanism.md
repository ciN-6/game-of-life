# Plan: Refactor to Spread Mechanism

## Overview
Replace hardcoded infection logic in `Grid.Step()` with a pluggable `Spread` mechanism. characters will define their influence via a `Spread` method, and a `SpreadResolver` will orchestrate the application of these effects.

## Objectives
- [ ] Define `SpreadEffect` and update `character` interface in `pkg/life/types/types.go`.
- [ ] Implement `Spread` for `Undeadcharacter` in `pkg/life/characters/undead_character.go`.
- [ ] Implement `Spread` for `Livingcharacter` (no-op) in `pkg/life/characters/living_character.go`.
- [ ] Refactor `Grid.Step()` in `pkg/life/life.go` to collect and apply spread effects.
- [ ] Ensure backward compatibility and test system behavior.

## Tasks
- [ ] Update `types.go`: Add `Spread` method and `SpreadEffect` struct.
- [ ] Update `base_character.go`: Implement default `Spread` (returning empty slice).
- [ ] Update `undead_character.go`: Implement infection logic inside `Spread`.
- [ ] Update `life.go`: Update `Step()` to resolve spread effects after state transitions.

### Task Detail
1. **Interface Update**: Add `Spread(grid Grid, x, y int) []SpreadEffect`.
2. **Resolution Phase**: In `Step()`, after `NextState` is called, iterate over all characters to collect `Spread` effects, then apply them.
3. **Infection Logic**: Move the current infection loop into `Undeadcharacter.Spread`.

## Definition of Done
- [ ] Hardcoded infection logic is removed from `Grid.Step()`.
- [ ] Spread effects correctly propagate from `Undeadcharacter`.
- [ ] System passes all existing tests.
