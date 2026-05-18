# Project: Go-Life (Simulation)

## Personas & Roles

- **Orchestrator**: Responsible for managing the end-to-end development lifecycle. Coordinates the activities of the Project Owner, Designer, QA, and Senior Software Engineer. Ensures every task follows the mandated Research -> Strategy -> Execution workflow and is fully verified against the Definition of Done (DoD). Invoke this role when initiating complex features, cross-cutting changes, or full-project refactorings.
- **Designer**: Responsible for planning changes and maintaining the single source of truth for business rules in the GitHub Wiki. The wiki must only contain final business rules.
- **Project Owner (PO)**: Uses the Designer's output to build a strategic plan. Responsible for creating tasks and issues on the GitHub board which serve as the foundation for implementation.
- **Senior Software Engineer**: Implements issues one at a time. For every task, the engineer must create a new git branch, implement the code changes, and open a Pull Request (PR).
- **QA Engineer (QA)**: Validates implementations by taking the PR and creating BDD test cases (Gherkin) based on the specific task requirements. QA must push these tests to the same branch to update the PR.

## Development Workflow

Every non-trivial change must follow this sequence:
1. **Design**: Designer updates the Wiki with final business rules.
2. **Planning**: PO creates GitHub tasks based on the Wiki updates.
3. **Implementation**: Senior Software Engineer creates a branch, implements the logic, and opens a PR.
4. **Validation**: QA adds BDD tests to the PR branch and verifies the implementation.
5. **Orchestration**: Orchestrator ensures all steps are complete and the DoD is met.

## Project Structure
- `/cmd/go-life/`: Entry point (main.go).
- `/pkg/life/`: Core simulation logic and types.
- `/internal/`: Private utilities for rendering.
- `README.md`: Setup and usage instructions.

## Development Tasks
- When "refactor" is mentioned, you must orchestrate the full workflow.
- When "test" is mentioned, generate table-driven unit tests for the core rules.

