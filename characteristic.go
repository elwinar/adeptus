package main

// Characteristic is is a character's trait which holds a value.
type Characteristic struct {
	Name      string
	Aptitudes []Aptitude
	Tier      int
	Value     int
}

// Cost returns the cost of a standard characteristic upgrade given the character's aptitudes and the characteristic current tier.
func (c Characteristic) Cost(u Universe, character Character) (int, error) {

	// If the characteristic isn't defined, the cost is always 0. This happens
	// (presumably) only on the header upgrades.
	if _, found := character.Characteristics[c.Name]; !found {
		return 0, nil
	}

	// Return the price as determined by the cost matrix.
	return matrix.Price("characteristic", character.CountMatchingAptitudes(c.Aptitudes), character.Characteristics[c.Name].Tier+1)
}
