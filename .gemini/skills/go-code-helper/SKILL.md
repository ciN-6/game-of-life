---
name: go-code-helper
description: Go code generation and TDD assistant. Use when writing new Go logic or updating existing features to ensure idiomatic code and comprehensive unit testing.
---

# Go Code Helper

This skill guides Gemini CLI in producing high-quality, idiomatic Go code and following Test-Driven Development (TDD) principles.

## Workflow

1. **Analyze Requirements**: Understand the requested code change.
2. **Draft Implementation**: Write the Go code, adhering strictly to:
   - [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
   - Project-specific standards (found in `GEMINI.md`)
   - Explicit error handling (no `panic`)
   - Descriptive variable names
   - Do not use deprecated functions.
   - Adding Godoc comments
3. **Generate Tests**: Create or update unit tests using Go's `testing` package.
   - Use table-driven tests for core logic.
   - Ensure test coverage is sufficient.
4. **Format**: Run `go fmt`.
5. **Verify**: Run `go test` and fix any issues.

## Go Idioms
- Always favor readability and maintainability.
- Keep functions small and focused.
- Use explicit error handling for all operations.


## Testing Standards
- Always create a test file alongside the implementation.
- Use table-driven tests for logic with multiple inputs/outputs.
- Run `go test ./...` after changes.
