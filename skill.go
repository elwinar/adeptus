package main

import (
	"fmt"
)

// Skill is a character's trait.
type Skill struct {
	Name       string
	Aptitudes  []Aptitude
	Tier       int
	Speciality string
}

// Cost returns the cost of the skill given the character's aptitudes and the current tier.
func (s Skill) Cost(universe Universe, character Character) (int, error) {

	// If the skill isn't defined, set the current tier to 0.
	tier := 0
	if _, found := character.Skills[s.Name]; found {
		tier = character.Characteristics[s.Name].Tier
	}

	// Return the price as determined by the cost matrix.
	return universe.Costs.Price("skill", character.CountMatchingAptitudes(s.Aptitudes), tier+1)
}

// FullName return the name of the skill and it's speciality if defined.
func (s Skill) FullName() string {
	if len(s.Speciality) == 0 {
		return s.Name
	}
	return fmt.Sprintf("%s: %s", s.Name, s.Speciality)
}
