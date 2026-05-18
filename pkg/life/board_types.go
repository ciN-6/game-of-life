package life

// SpreadEffect represents a change propagated by a character to another cell.
type SpreadEffect struct {
	TargetX, TargetY           int
	SourceX, SourceY           int
	DestinationX, DestinationY int
	NewCell                    Cell
	Weight                     int
	TargetType                 string
}

// Cell acts as a container for a simulation Character and its coordinates.
type Cell struct {
	X, Y       int       // Coordinates on the grid
	DeathCount int       // Persistent property tracking entity deaths
	Character  Character // Behavior implementation
}
