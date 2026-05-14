# Bug Report: Undead Spawn Test Failure

## Description
The BDD scenario `Undead spawn` in `simulation.feature` failed.
The test expects an `UndeadCharacter` after 10 steps, but it remained a `LivingCharacter`.
This is likely because the recent change requires 5 deaths for an undead to spawn, and the test environment is not triggering enough deaths to meet this new threshold within the 10 steps.

## Steps to Reproduce
1. Run BDD tests: `go test -v -tags=cucumber ./...`
2. Observe `Scenario: Undead spawn` failure.

## Impact
Feature specification is failing due to the recent logic change. The test expectation needs to be adjusted to reflect the new business rule (5 deaths).

## Recommendation
Update the `simulation.feature` scenario or the test setup to ensure sufficient death accumulation to trigger the spawn, or update the test expectations to align with the new threshold.
