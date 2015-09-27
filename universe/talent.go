package universe

// Talent is a character's trait.
type Talent struct {
	Name         string
	Aptitudes    []Aptitude
	Tier         int
	Requirements []Requirement
}

// Cost returns the cost of the talent given the character's aptitudes and the current tier.
func (t Talent) Cost(matrix CostMatrix, aptitudes []Aptitude) (int, error) {

	// Retrieve the number of matching aptitudes between the character's aptitudes and the talent's aptitudes
	matching := countMatches(aptitudes, t.Aptitudes)

	// Return the price of the upgrade as determined by the cost matrix.
	return matrix.Price("talent", matching, t.Tier)
}
