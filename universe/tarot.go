package universe

// Tarot is a character's trait providing him alterations.
type Tarot struct {
	history

	// Min is the lowest value of the tarot range.
	Min int

	// Max is the highest value of the tarot range.
	Max int
}
