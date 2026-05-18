# Design: Merging Grid Interface into Board

## Objective
Simplify the simulation architecture by removing the `types.Grid` interface and merging its functionality directly into the `Board` struct.

## Technical Challenge: Circular Dependency
Currently, the `Character` interface (in `pkg/life/types`) depends on the `Grid` interface (also in `pkg/life/types`). The `Board` (in `pkg/life`) implements `Grid` and uses `Character`.
If we remove `Grid` and have `Character` depend directly on `Board`, we create a circular dependency:
`pkg/life/types` (Character) -> `pkg/life` (Board)
`pkg/life` (Board) -> `pkg/life/types` (Character)

## Proposed Solution: Package Consolidation
To resolve this while fulfilling the requirement to merge `Grid` into `Board`, we will consolidate the core simulation types and implementations into a single package: `pkg/life`.

### Structural Changes
1. **Consolidate `pkg/life/types`**:
    - Move `types.Character` interface to `pkg/life/character.go`.
    - Move `types.Cell` and `types.SpreadEffect` to `pkg/life/board_types.go` (or similar).
    - Remove `pkg/life/types/Grid.go` and its interface.
2. **Consolidate `pkg/life/characters`**:
    - Move `LivingCharacter` and `UndeadCharacter` into `pkg/life/`.
    - This eliminates the need for `pkg/life` to import `pkg/life/characters`.
3. **Merge Grid into Board**:
    - Implement `ForEachNeighbor` directly as a method of `*Board`.
    - Update `Character` methods to accept `*Board`.

## Impact Analysis
- All files currently in `pkg/life/types` and `pkg/life/characters` will be moved to `pkg/life`.
- Package declarations will be updated to `package life`.
- Import paths in `cmd/go-life/main.go` and tests will be updated.

## Transition Plan
1. **Planning**: PO creates tasks for consolidation and refactoring.
2. **Execution**:
    - Move files and update package names.
    - Remove `Grid` interface and update `Character` interface.
    - Update all method signatures and call sites.
    - Fix imports across the project.
3. **Validation**: QA ensures all simulation rules (Living/Undead) still hold via BDD tests.
