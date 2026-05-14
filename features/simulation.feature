Feature: Game of Life Simulation
  Scenario: A single character dies
    Given a grid with a single living character at 0, 0
    When the simulation steps once
    Then the character at 0, 0 should be dead

  Scenario: Blinker rotation
    Given a grid with living characters at:
      | x | y |
      | 4 | 4 |
      | 4 | 5 |
      | 4 | 6 |
When the simulation steps once
    Then the following characters should be alive:
      | x | y |
      | 3 | 5 |
      | 4 | 5 |
      | 5 | 5 |

  Scenario: Reproduction
    Given a grid with living characters at:
      | x | y |
      | 4 | 4 |
      | 4 | 5 |
      | 5 | 5 |  
    When the simulation steps once
    Then the following characters should be alive:
      | x | y |
      | 4 | 4 |
      | 4 | 5 |
      | 5 | 5 |  
      | 5 | 5 |
  
  @stable
  Scenario: Stable population
    Given a grid with living characters at: 
      | x | y |
      | 4 | 4 |
      | 4 | 5 |
      | 5 | 4 |
      | 5 | 5 |
    When the simulation steps once
    Then the following characters should be alive:
      | x | y |
      | 4 | 4 |
      | 4 | 5 |
      | 5 | 4 |
      | 5 | 5 |

  Scenario: Overpopulated
    Given a grid with living characters at: 
      | x | y |
      | 4 | 3 |
      | 4 | 4 |
      | 4 | 5 |
      | 5 | 4 |
      | 3 | 4 |
    When the simulation steps once
    Then the following characters should be dead:
      | x | y |
      | 4 | 4 |

  Scenario: Undead character persists
    Given a grid with an undead character at 4, 4
    When the simulation steps 1 times
    Then the character at 4, 4 should be undead

  Scenario: Undead spawn
    Given a grid with living characters at:
      | x | y |
      | 4 | 4 |
      | 4 | 5 |
      | 4 | 6 |
    When the simulation steps 10 times
    Then the character at 4, 4 should be undead
  
  Scenario: dead character dont become undead
    Given a grid with an undead character at 4, 4
    When the simulation steps 1 times
    Then the character at 4, 3 should be dead