---
name: project-planner
description: Automates project planning by creating a standardized markdown file in ./specs/ for a given project name. Use this skill when a user asks for a plan for a new task.
---

# Project Planner

This skill helps you create a standardized plan file in the `./specs/` directory.

## Usage

When a user asks for a plan, follow these steps:

1. Identify the project name from the user.
2. Create a file at `specs/<name-of-plan>.md` inside the project.
3. Populate the file with the template defined in `assets/template.md`.
4. Make sure the plan includes unit testing for all edits. Unit testing should cover at least 90% of the code edits.
5. Add cucumber tests for feature testing if applicable.


## Resources

- **Template**: The template is located at `assets/template.md`. Use this for all plans.
