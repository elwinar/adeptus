package main

import (	
	"gopkg.in/yaml.v2"
)

// Spell castable.
type Spell struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	XP          int    `yaml:"xp"`
	Attributes  map[string]interface{}
}

// Cost returns the default cost value of the spell.
func (s Spell) Cost(u Universe, character Character) (int, error) {

	return s.XP, nil
}

// Apply applys the upgrade on the character:
// * does not affect the character's XP
func (s Spell) Apply(character *Character, upgrade Upgrade) error {

	// Get the gauge from the character.
	_, found := character.Spells[s.Name]
	if found {
		return NewError(DuplicateUpgrade, upgrade.Line)
	}

	// Set the spell to the map.
	character.Spells[s.Name] = s

	return nil
}

// UnmarshalYAML implements the Unmarshaler interface
func (s *Spell) UnmarshalYAML(raw []byte) error {
	err := yaml.Unmarshal(raw, &s.Attributes)
	if err != nil {
		return err
	}

	xp, ok := s.Attributes["xp"]
	if ok {
		xp, ok := xp.(float64)
		if ok {
			s.XP = int(xp)
			delete(s.Attributes, "xp")
		}
	}

	s.Name, ok = s.Attributes["name"].(string)
	if ok {
		delete(s.Attributes, "name")
	}

	s.Description, ok = s.Attributes["description"].(string)
	if ok {
		delete(s.Attributes, "description")
	}

	return nil
}

// DefaultName returns the default upgrade name.
func (s Spell) DefaultName() string {
	return s.Name
}
