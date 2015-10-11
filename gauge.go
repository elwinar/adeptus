package main

// Gauge represent a trait of the character.
// Typically used to track fate points, magic rating, corruption, insanity,
// and the type of attributes that generally aren't bought by spending
// experience points.
type Gauge struct {
	Name  string `json:"name"`
	Value int    `json:"-"`
}

// Cost returns 0, a gauge has no calculated cost.
func (g Gauge) Cost(u Universe, character Character) (int, error) {

	return 0, nil
}
