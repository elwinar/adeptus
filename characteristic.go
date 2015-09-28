package main

// Characteristic is is a character's trait which holds a value.
type Characteristic struct {
	Name      string
	Aptitudes []Aptitude
	Tier      int
}

// Cost returns the cost of a standard characteristic upgrade given the character's aptitudes and the characteristic current tier.
func (c Characteristic) Cost(matrix CostMatrix, aptitudes []Aptitude) (int, error) {

	// Retrieve the number of matching aptitudes between the character's aptitudes and the characteristic's aptitudes
	matching := countMatches(aptitudes, c.Aptitudes)

	// Return the price as determined by the cost matrix.
	return matrix.Price("characteristic", matching, c.Tier)
}
