package universe

// Talent is a character's trait
type Talent struct {
	Name         string
	Aptitudes    []Aptitude
	Tier         int
	Requirements []Requirement
}

// Cost returns the cost of the talent given the character's aptitudes and the current tier
func (t Talent) Cost(matrix CostMatrix, aptitudes []Aptitude) (int, error) {

	// count matches
	var m int
	for _, a := range aptitudes {
		for _, ta := range t.Aptitudes {
			if a == ta {
				m++
			}
		}
	}

	// price
	return matrix.Price("talent", m, t.Tier)
}
