# Task: Fix Undead Spawn Threshold

Title: Fix Undead Spawn Logic
Priority: High
Description: Current implementation spawns an UndeadCharacter immediately upon the first death of a LivingCharacter. It should only spawn an UndeadCharacter when the cell's death count reaches 5.
Acceptance Criteria:
- [ ] LivingCharacter death increments the cell's death count.
- [ ] UndeadCharacter spawns only if cell.DeathCount >= 5.
- [ ] Existing logic is updated in LivingCharacter.NextState or appropriate location.

Labels: bug, game-logic
