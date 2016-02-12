package main

// Background represents an element providing traits to a character.
type Background struct {
	Type     string   `yaml:"type"`
	Name     string   `yaml:"name"`
	Upgrades []string `yaml:"upgrades"`
}

// Apply changes the character's trait according to the history values
func (b Background) Apply(character *Character, universe Universe) error {

	// For each upgrade associated to the history, apply each option.
	for _, raw := range b.Upgrades {

		upgrade := Upgrade{
			Mark: MarkApply,
			Name: raw,
			Cost: IntP(0),
		}
		_, found := universe.FindCharacteristic(upgrade)
		if found {
			upgrade.Mark = MarkSpecial
		}

		err := character.ApplyUpgrade(upgrade, universe)
		if err != nil {
			return err
		}
	}

	// Add the background to the character's backgrounds
	character.Backgrounds[b.Name] = b

	return nil
}
