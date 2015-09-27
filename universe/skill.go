package universe

// Skill is a character's trait.
type Skill struct {
	Name      string
	Aptitudes []Aptitude
	Tier      int
}

// Cost returns the cost of the skill given the character's aptitudes and the current tier.
func (s Skill) Cost(matrix CostMatrix, aptitudes []Aptitude) (int, error) {

	// Retrieve the number of matching aptitudes between the character's aptitudes and the skill's aptitudes
	matching := countMatches(aptitudes, s.Aptitudes)

	// Return the price of the upgrade as determined by the cost matrix.
	return matrix.Price("skill", matching, s.Tier)
}
