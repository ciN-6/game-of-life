Title: Consolidate packages into pkg/life
Priority: High
Description: Move all files from pkg/life/types and pkg/life/characters into pkg/life to prepare for merging the Grid interface into the Board struct. This resolves potential circular dependencies.
Acceptance Criteria:
- [ ] Move LivingCharacter.go and UndeadCharacter.go to pkg/life/
- [ ] Move character.go, Grid.go, and board-related types to pkg/life/
- [ ] Update all package declarations to `package life`
- [ ] Fix imports in Board.go and other moved files
- [ ] Ensure project compiles (basic check)
Labels: refactor, architecture
