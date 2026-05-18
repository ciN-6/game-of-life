Feature: Game of Life Simulation
  Scenario: A single character dies
    Given a grid with a single living character at 0, 0
    When the simulation steps once
    Then the character at 0, 0 should be dead

  Scenario: Blinker expansion
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
      | 3 | 4 |
      | 3 | 6 |
      | 5 | 4 |
      | 5 | 6 |

  Scenario: Reproduction with 2 neighbors
    Given a grid with living characters at:
      | x | y |
      | 4 | 4 |
      | 4 | 5 |
    When the simulation steps once
    Then the following characters should be alive:
      | x | y |
      | 3 | 4 |
      | 5 | 4 |
      | 3 | 5 |
      | 5 | 5 |
    And the following characters should be dead:
      | x | y |
      | 4 | 4 |
      | 4 | 5 |

  Scenario: Reproduction with 3 neighbors
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
      | 5 | 4 |
      | 4 | 6 |
  
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

  Scenario: dead character dont become undead
    Given a grid with an undead character at 4, 4
    When the simulation steps 1 times
    Then the character at 4, 3 should be dead

  Scenario: Undead infection
    Given a grid with living characters at:
      | x | y |
      | 4 | 4 |
    And a grid with an undead character at 4, 5
    When the simulation steps once
    Then the character at 4, 4 should be undead

  Scenario: Neighbor counting
    Given a grid with living characters at:
      | x | y |
      | 4 | 4 |
      | 4 | 5 |
      | 5 | 4 |
    Then the cell at 5, 5 should have 3 neighbors