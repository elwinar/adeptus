package universe

// Characteristic is is a character's trait which holds a value
type Characteristic struct {
	Name      string
	Aptitudes []Aptitude
	Tier      int
}

// Cost returns the cost of the characteristic given the character's aptitudes
func (c Characteristic) Cost(matrix CostMatrix, aptitudes []Aptitude) (int, error) {

	// count matches
	var m int
	for _, a := range aptitudes {
		for _, ca := range c.Aptitudes {
			if a == ca {
				m++
			}
		}
	}

	// price
	return matrix.Price("characteristic", m, c.Tier)
}
