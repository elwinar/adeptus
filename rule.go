package main

// Rule represent a special rule, which are generally home-made additions to the
type Rule struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Cost returns 0, a rule has no calculated cost.
func (r Rule) Cost(u Universe, character Character) (int, error) {

	return 0, nil
}

// Apply applys the upgrade on the character:
// * gives the rule to the character
// * does not affect the character's XP
func (r Rule) Apply(character *Character, upgrade Upgrade) error {

	switch upgrade.Mark {
	case MarkSpecial:
		return NewError(ForbidenUpgradeMark, upgrade.Line)
	case MarkRevert:
		_, found := character.Rules[r.Name]
		if !found {
			return NewError(ForbidenUpgradeLoss, upgrade.Line)
		}
		delete(character.Rules, r.Name)
	case MarkApply:
		character.Rules[r.Name] = r
	}

	return nil
}

// DefaultName returns the default upgrade name.
func (r Rule) DefaultName() string {
	return r.Name
}
