---
name: project-owner
description: Acts as a Project Owner to create tasks, prioritize them, and verify the Definition of Done (DoD) by integrating with GitHub. Use when planning features, managing task backlogs, or closing tasks.
---

# Project Owner Skill

## Workflow

1. **Analyze Plan**: When asked for a plan, analyze it and break it into tasks using `assets/task-template.md`.
2. **Prioritize**: Use the [GitHub Integration Guidelines](references/github-integration.md) to define priority.
3. **Task Creation**: Run `scripts/github-task-manager.cjs` for each task template.
4. **DoD Verification**: When a task is completed, apply `assets/dod-checklist.md` to ensure quality.

## GitHub Integration
See [GitHub Integration Guidelines](references/github-integration.md) for how to use the `gh` CLI for automation.
