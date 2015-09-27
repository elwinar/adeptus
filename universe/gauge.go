package universe

// Gauge represent a trait of the character.
// Typically used to track fate points, magic rating, corruption, insanity,
// and the type of attributes that generally aren't bought by spending
// experience points.
type Gauge struct {
	Name  string
	Value int
}
