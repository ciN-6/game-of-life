# Plan: Add User Input for Grid Size

## Overview
Implement an interactive selection menu during game runtime where the user can choose from 3 predefined grid size templates. Update the rendering logic to redraw the window to accommodate the selected grid size.

## Objectives
- [ ] Add interactive selection menu for grid size templates.
- [ ] Implement window redrawing logic to adjust to the selected grid size.
- [ ] Update simulation initialization to use the chosen template dimensions.

## Tasks
- [ ] Define the 3 grid size templates (e.g., Small, Medium, Large).
- [ ] Implement the selection menu in `cmd/go-life/main.go`.
- [ ] Ensure the renderer/window handles size changes (redraw functionality).
- [ ] Update `pkg/life/life.go` to support configurable grid sizes.
- [ ] Add unit tests for template selection and grid resizing.

### Task detail
- Templates: Define constants or a map for the 3 grid sizes.
- Selection: Implement a combobox-style selection element for the user to pick from the templates.
- Redraw: Research terminal/console library used for rendering to trigger a window/buffer refresh.
- Update `pkg/life`: Ensure `NewGrid` accepts dimensions.

## Definition of Done
- [ ] Grid size is selected via an interactive menu.
- [ ] Rendering window updates (redraws) correctly to fit the chosen size.
- [ ] Unit tests pass for template selection and grid initialization.
