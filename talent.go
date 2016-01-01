package main

import "fmt"

// Talent is a character's trait.
type Talent struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Aptitudes    []Aptitude    `json:"aptitudes"`
	Tier         int           `json:"tier"`
	Requirements []Requirement `json:"requirements"`
	Speciality   string        `json:"-"`
	Value        int           `json:"-"`
	Stackable    bool          `json:"stackable"`
}

// Cost returns the cost of the talent given the character's aptitudes and the current tier.
func (t Talent) Cost(universe Universe, character Character) (int, error) {

	// Return the price as determined by the cost matrix.
	return universe.Costs.Price("talent", character.Intersect(t.Aptitudes), t.Tier)
}

// FullName return the name of the talent and it's speciality if defined.
func (t Talent) FullName() string {
	if len(t.Speciality) == 0 {
		return t.Name
	}
	return fmt.Sprintf("%s: %s", t.Name, t.Speciality)
}

// Apply applys the upgrade on the character:
// * affect the talent tier
// * affect the talent value if stackable
// * does not affect the character's XP
func (t Talent) Apply(character *Character, upgrade Upgrade) error {

	// Get the talent from the character.
	tmp, found := character.Talents[t.FullName()]
	if found {
		t = tmp
	}

	switch upgrade.Mark {
	case MarkSpecial:
		return NewError(ForbidenUpgradeMark, upgrade.Line)
	case MarkRevert:
		t.Value--
	case MarkApply:
		t.Value++
	}

	// Remove the talent if it is negative.
	if t.Value <= 0 {
		if !found {
			return NewError(ForbidenUpgradeLoss, upgrade.Line)
		}
		delete(character.Talents, t.FullName())
		return nil
	}

	// Check the talent is stackable.
	if !t.Stackable && t.Value > 1 {
		return NewError(DuplicateUpgrade, upgrade.Line)
	}

	// Put it back on the map.
	character.Talents[t.FullName()] = t

	return nil
}

// DefaultName returns the default upgrade name.
func (t Talent) DefaultName() string {
	return t.FullName()
}
