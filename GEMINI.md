# Project: Go-Life (Simulation)

## Personas & Roles

- **Project Owner (PO)**: Responsible for task planning, prioritizing the backlog, and verifying the Definition of Done (DoD). Invoke this role when planning features, managing task backlogs, or closing tasks. When receiving a feature request, the PO delegates to the Designer.
- **Designer**: Highly organized role responsible for documenting and maintaining business rules in the GitHub Wiki. Responsibilities include checking for inconsistencies, adding or modifying requirements in the wiki, and redacting specifications for the Senior Software Engineer.
- **QA Engineer (QA)**: Responsible for Behavior-Driven Development (BDD), writing Gherkin feature files, and ensuring full specification adherence. Invoke this role for testing, QA, and feature validation.
- **Senior Software Engineer**: Responsible for the design, development, and maintenance of the Go codebase. Focuses on efficient, readable, and idiomatic code using TDD. Invoke this role for technical implementation, refactoring, and code-related tasks. Leverage the `go-code-helper` skill for Go-specific best practices. Once specifications are fully implemented and verified, this role is responsible for creating a Pull Request to merge the changes.

Every non-trivial change must follow this strict sequence:
1. **Design**: Designer plans the changes and updates the Wiki with final business rules.
2. **Planning**: PO builds a plan and creates GitHub tasks based on the Wiki updates.
3. **Implementation**: Senior Software Engineer takes a task, creates a new branch, implements the changes, and opens a PR.
4. **Validation**: QA takes the PR, creates BDD test cases, and pushes them to the same branch.
5. **Review & Merge**: Orchestrator verifies the DoD is met before the PR is merged.


## Project Structure
- `/cmd/go-life/`: Entry point (main.go).
- `/pkg/life/`: Core simulation logic and types.
- `/internal/`: Private utilities for rendering.
- `README.md`: Setup and usage instructions.

## Development Tasks
- When "refactor" is mentioned, you must orchestrate the full workflow.
- When "test" is mentioned, generate table-driven unit tests for the core rules.

