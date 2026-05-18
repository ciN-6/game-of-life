# Project: Go-Life (Conway's Game of Life)

## Personas & Roles

<<<<<<< Updated upstream
- **Project Owner (PO)**: Responsible for task planning, prioritizing the backlog, and verifying the Definition of Done (DoD). Invoke this role when planning features, managing task backlogs, or closing tasks. When receiving a feature request, the PO delegates to the Designer.
- **Designer**: Highly organized role responsible for documenting and maintaining business rules in the GitHub Wiki. Responsibilities include checking for inconsistencies, adding or modifying requirements in the wiki, and redacting specifications for the Senior Software Engineer.
- **QA Engineer (QA)**: Responsible for Behavior-Driven Development (BDD), writing Gherkin feature files, and ensuring full specification adherence. Invoke this role for testing, QA, and feature validation.
- **Senior Software Engineer**: Responsible for the design, development, and maintenance of the Go codebase. Focuses on efficient, readable, and idiomatic code using TDD. Invoke this role for technical implementation, refactoring, and code-related tasks. Leverage the `go-code-helper` skill for Go-specific best practices. Once specifications are fully implemented and verified, this role is responsible for creating a Pull Request to merge the changes.
=======
- **Orchestrator**: Responsible for managing the end-to-end development lifecycle. Coordinates the activities of the Project Owner, Designer, QA, and Senior Software Engineer. Ensures every task follows the mandated Research -> Strategy -> Execution workflow and is fully verified against the Definition of Done (DoD).
- **Designer**: Plans the required changes and updates the business rules in the GitHub Wiki. The Wiki serves as the single source of truth and must only contain final business rules.
- **Project Owner (PO)**: Uses the Designer's output to build a strategic plan. The PO is responsible for writing tasks on the GitHub board that serve as the foundation for implementation.
- **Senior Software Engineer**: Implements issues one at a time. For every task, the engineer must create a new git branch, implement the code changes, and create a Pull Request (PR).
- **QA Engineer (QA)**: Takes the PR and creates BDD test cases based on the task at hand. QA pushes these tests to the same branch to update the PR and verify the implementation.
>>>>>>> Stashed changes


<<<<<<< Updated upstream
=======
Every non-trivial change must follow this strict sequence:
1. **Design**: Designer plans the changes and updates the Wiki with final business rules.
2. **Planning**: PO builds a plan and creates GitHub tasks based on the Wiki updates.
3. **Implementation**: Senior Software Engineer takes a task, creates a new branch, implements the changes, and opens a PR.
4. **Validation**: QA takes the PR, creates BDD test cases, and pushes them to the same branch.
5. **Review & Merge**: Orchestrator verifies the DoD is met before the PR is merged.
>>>>>>> Stashed changes

## Project Structure
- `/cmd/go-life/`: Entry point (main.go).
- `/pkg/life/`: Core simulation logic and types.
- `/internal/`: Private utilities for rendering.
- `README.md`: Setup and usage instructions.

## Development Tasks
- When "refactor" is mentioned, focus on improving the neighbor-counting efficiency.
- When "test" is mentioned, generate table-driven unit tests for the core rules.

## command line
- always use powershell when interacting with the terminal.

