package main

import (
	"strconv"
	"strings"
)

// Characteristic is is a character's trait which holds a value.
type Characteristic struct {
	Name      string     `json:"name"`
	Aptitudes []Aptitude `json:"aptitudes"`
	Tier      int        `json:"tier"`
	Value     int        `json:"-"`
}

// Cost returns the cost of a standard characteristic upgrade given the character's aptitudes and the characteristic current tier.
func (c Characteristic) Cost(universe Universe, character Character) (int, error) {

	// If the characteristic isn't defined, the cost is always 0. This happens
	// (presumably) only on the header upgrades.
	if _, found := character.Characteristics[c.Name]; !found {
		return 0, nil
	}

	// Return the price as determined by the cost matrix.
	return universe.Costs.Price("characteristic", character.Intersect(c.Aptitudes), character.Characteristics[c.Name].Tier+1)
}

// Level returns a string representing the tier of the characteristic.
func (c Characteristic) Level() string {
	var out string
	var i int
	for i < c.Tier {
		out += "*"
		i++
	}
	return out
}

// Apply applys the upgrade on the character:
// * affect the characteristics tier
// * affect the characteristic value
// * does not affect the character's XP
func (c Characteristic) Apply(character *Character, upgrade Upgrade) error {

	// Get the attribute from the character's characteristic map.
	tmp, found := character.Characteristics[c.Name]
	if found {
		c = tmp
	}

	// Affect the characteristic's tier
	switch upgrade.Mark {
	case MarkApply:
		c.Tier++
	case MarkRevert:
		c.Tier--
	}

	// Check the tier is not negative.
	if c.Tier < 0 {
		return NewError(ForbidenUpgradeLoss, upgrade.Line, c.Name)
	}

	// Parse the characteristic's upgrade value.
	raw := strings.TrimSpace(strings.TrimLeft(upgrade.Name, c.Name))
	value, err := strconv.Atoi(raw)
	if err != nil {
		return NewError(InvalidUpgradeValue, upgrade.Line)
	}

	// Deny absolute values for non special marks.
	if upgrade.Mark != MarkSpecial && !strings.HasPrefix(raw, "+") {
		return NewError(ForbidenUpgradeValue, upgrade.Line)
	}

	// Update the characteristic's value.
	if strings.HasPrefix(raw, "+") || strings.HasPrefix(raw, "-") {
		c.Value += value
	} else {
		c.Value = value
	}

	character.Characteristics[c.Name] = c

	return nil
}
