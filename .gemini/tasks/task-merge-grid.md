Title: Merge Grid interface into Board
Priority: High
Description: Remove the Grid interface and update the Character interface and all simulation logic to use *Board directly.
Acceptance Criteria:
- [ ] Remove Grid interface definition
- [ ] Move ForEachNeighbor to Board.go as a method of *Board
- [ ] Update Character interface methods (PrepareAction, ApplyAction, NextState) to accept *Board
- [ ] Update LivingCharacter and UndeadCharacter implementation methods to accept *Board
- [ ] Update all call sites of ForEachNeighbor and Character methods
Labels: refactor, architecture
