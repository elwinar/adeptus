package main

import (
	"strconv"
	"strings"
)

// Gauge represent a trait of the character.
// Typically used to track fate points, magic rating, corruption, insanity,
// and the type of attributes that generally aren't bought by spending
// experience points.
type Gauge struct {
	Name  string `json:"name"`
	Value int    `json:"-"`
	XP    int    `json:"xp"`
}

// Cost returns 0, a gauge has no calculated cost.
func (g Gauge) Cost(u Universe, character Character) (int, error) {

	return g.XP * g.Value, nil
}

// Apply applys the upgrade on the character:
// * affect the gauge value
// * does not affect the character's XP
func (g Gauge) Apply(character *Character, upgrade Upgrade) error {

	// Get the gauge from the character.
	old, found := character.Gauges[g.Name]
	if !found {
		old = g
		old.Value = 0
	}

	// Parse the gauge's upgrade value.
	raw := strings.TrimSpace(strings.TrimLeft(upgrade.Name, g.Name))
	value, err := strconv.Atoi(raw)
	if err != nil {
		return NewError(InvalidUpgradeValue, upgrade.Line)
	}

	// Update the gauge value.
	if !(strings.HasPrefix(raw, "+") || strings.HasPrefix(raw, "-")) {
		return NewError(ForbidenUpgradeValue, upgrade.Line)
	}
	old.Value += value

	// Set the gauge back on the map.
	character.Gauges[g.Name] = old

	return nil
}
