package main

// Rule represent a special rule, which are generally home-made additions to the
type Rule struct {
	Name        string
	Description string
}

// Cost returns 0, a rule has no calculated cost.
func (r Rule) Cost(u Universe, character Character) (int, error) {

	return 0, nil
}
