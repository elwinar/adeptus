package universe

// Skill is a character's trait
type Skill struct {
	Name      string
	Aptitudes []Aptitude
	Tier      int
}

// Cost returns the cost of the skill given the character's aptitudes and the current tier
func (s Skill) Cost(matrix CostMatrix, aptitudes []Aptitude) (int, error) {

	// count matches
	var m int
	for _, a := range aptitudes {
		for _, sa := range s.Aptitudes {
			if a == sa {
				m++
			}
		}
	}

	// price
	return matrix.Price("skill", m, s.Tier)
}
