package main

// Aptitude represents an aptitude, required to purchase upgrades.
type Aptitude string

// Cost return the cost of the aptitude. Implements Coster.
func (a Aptitude) Cost(u Universe, c Character) (int, error) {
	return 0, nil
}

// Apply applys the upgrade on the character:
// * give the aptitute to the character
// * does not affect the character's XP
func (a Aptitude) Apply(character *Character, upgrade Upgrade) error {

	switch upgrade.Mark {
	case MarkSpecial:
		return NewError(ForbidenUpgradeMark, upgrade.Line)

	case MarkApply:
		character.Aptitudes[string(a)] = a

	case MarkRevert:
		_, found := character.Aptitudes[string(a)]
		if !found {
			return NewError(ForbidenUpgradeLoss, upgrade.Line)
		}
		delete(character.Aptitudes, string(a))
	}
	return nil
}

// DefaultName returns the default upgrade name.
func (a Aptitude) DefaultName() string {
	return string(a)
}
