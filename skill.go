package main

import (
	"fmt"
)

// Skill is a character's trait.
type Skill struct {
	Name       string     `json:"name"`
	Aptitudes  []Aptitude `json:"aptitudes"`
	Tier       int        `json:"tier"`
	Speciality string     `json:"-"`
}

// Cost returns the cost of the skill given the character's aptitudes and the current tier.
func (s Skill) Cost(universe Universe, character Character) (int, error) {

	// If the skill isn't defined, set the current upgrade to 0.
	tier := 0
	if _, found := character.Skills[s.FullName()]; found {
		tier = character.Skills[s.FullName()].Tier
	}

	// Return the price as determined by the cost matrix.
	return universe.Costs.Price("skill", character.Intersect(s.Aptitudes), tier+1)
}

// FullName return the name of the skill and it's speciality if defined.
func (s Skill) FullName() string {
	if len(s.Speciality) == 0 {
		return s.Name
	}
	return fmt.Sprintf("%s: %s", s.Name, s.Speciality)
}

// Apply applys the upgrade on the character:
// * affect the skill tier
// * does not affect the character's XP
func (s Skill) Apply(character *Character, upgrade Upgrade) error {

	// Get the skill from the character's skill map.
	tmp, found := character.Skills[s.FullName()]
	if found {
		s = tmp
	}

	switch upgrade.Mark {
	case MarkSpecial:
		return NewError(ForbidenUpgradeMark, upgrade.Line)
	case MarkRevert:
		s.Tier--
	case MarkApply:
		s.Tier++
	}

	// Remove the skill if it is negative.
	if s.Tier <= 0 {
		if !found {
			return NewError(ForbidenUpgradeLoss, upgrade.Line)
		}
		delete(character.Skills, s.FullName())
		return nil
	}

	// Put the skill back on the map.
	character.Skills[s.FullName()] = s

	return nil
}

// DefaultName returns the default upgrade name.
func (s Skill) DefaultName() string {
	return s.FullName()
}
