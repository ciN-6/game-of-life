# Refactoring Plan: character-Oriented Logic

## Objective
Refactor the character state logic from `pkg/life/life.go` into individual `character` types to improve encapsulation and adhere to the user's request for "character responsibility."

## Proposed Design (Go Idiomatic Polymorphism)
Since Go lacks traditional classes, I will use:
- **`character` Interface**: Defines behavior for all character types (`NextState(neighbors int, grid *Grid, x, y int) character`).
- **Structs**:
    - `Basecharacter`: Holds common state (age, death count).
    - `Livingcharacter`: Implements standard Conways rules + transition to Undead.
    - `Undeadcharacter`: Implements longevity and infectious behavior.
- **Grid Refactoring**: `Grid.characters` will become a slice of the `character` interface.

## Implementation Steps
1. **Define the Interface**: Create the `character` interface in `pkg/life/character.go`.
2. **Implement Structs**: Define `Livingcharacter` and `Undeadcharacter` structs with their specific logic.
3. **Refactor `Grid.Step()`**: Update `Grid.Step()` to delegate to the `character`'s `NextState` method.
4. **Update Initialization**: Ensure `NewGrid` populates the grid with initial `Livingcharacter` instances.
5. **Testing**: Run existing BDD and unit tests to ensure behavior remains consistent.

## Verification
- Validate that `Step()` still results in the same game outcomes.
- Ensure the BDD scenarios (including the new Undead test) pass.
- Verify memory management (avoid unnecessary allocations in the `Step` loop).

## Alternatives Considered
- **Direct Refactoring**: Keep the current struct approach but add methods to it.
    - *Cons*: Does not satisfy the user's requirement for "specialization" via class hierarchies.
- **Factory Pattern**: Create a `characterFactory`.
    - *Cons*: Potentially overkill for this project's scale.
