# GitHub Integration Guidelines

This skill uses the `gh` (GitHub CLI) to interact with your project.

## Prerequisites
- User must be authenticated with `gh auth login`.
- The repository must have GitHub Projects enabled.

## Usage
When the agent needs to create tasks, it will call `scripts/github-task-manager.cjs`.

The script uses:
`gh issue create --title "..." --body "..." --label "..."`

## Automation
- To create a task: Call the script with the path to a completed `task-template.md`.
- To verify DoD: Check that the PR/Issue includes the `dod-checklist.md` as part of the description or checklist.
