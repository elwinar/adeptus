package universe

// Characteristic is is a character's trait which holds a value.
type Characteristic struct {
	Name      string
	Aptitudes []Aptitude
	Tier      int
}

// Cost returns the cost of a standard characteristic upgrade given the character's aptitudes and the characteristic current tier.
func (c Characteristic) Cost(matrix CostMatrix, aptitudes []Aptitude) (int, error) {

	// Count the number of matching aptitudes.
	var m int
	for _, a := range aptitudes {
		for _, ca := range c.Aptitudes {
			if a == ca {
				m++
			}
		}
	}

	// Return the price as determined by the cost matrix.
	return matrix.Price("characteristic", m, c.Tier)
}
