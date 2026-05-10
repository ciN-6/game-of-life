---
name: implement-new-character-type
description: Workflow for implementing new character types in Conway's Game of Life. Use when adding a new character type with unique behavior or attributes.
---

# Implement New character Type

This skill guides the implementation of new character types in the Go-Life simulation.

## Workflow

1. **Verify basic rules**: Ensure the implementation covers the core rules (Underpopulation, Survival, Overpopulation, Reproduction) for the new character type.
2. **Clarify requirements**: If the rules for the new character type seem unclear or ambiguous, ask the user for clarification.
3. **Check for conflicts**: Analyze the codebase to ensure the new character type does not conflict with existing rules or character behaviors.
4. **Maintain tests**: Do NOT modify existing Cucumber features or tests. If an existing test fails, stop and ask the user for guidance.
5. **Add new tests**: Add new Cucumber features/tests to cover the specific behavior of the new character type.
6. **Execute tests**: Run all relevant tests (Cucumber/BDD) after making any code edits to ensure system integrity.
7. **build**: build the application.