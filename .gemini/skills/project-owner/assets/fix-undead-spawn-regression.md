# Bug Report: Undead Spawn Regression

## Description
The 'Undead spawn' BDD scenario fails even after 10 simulation steps. The LivingCharacter is not transforming into an UndeadCharacter despite exceeding the 5-death threshold.

## QA Findings
- Test: Undead spawn (10 steps)
- Error: expected undead at (4,4), but got type *characters.BaseCharacter (IsUndead: false)

## Impact
Critical regression of the new Undead spawn business rule.

## Acceptance Criteria
- [ ] Fix logic in LivingCharacter.NextState to ensure transformation triggers at 5 deaths.
- [ ] Verify that character death counts accumulate correctly across simulation cycles.
- [ ] Ensure Undead spawn test passes after fix.

Labels: bug
