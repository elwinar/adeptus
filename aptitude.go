package main

// Aptitude represents an aptitude, required to purchase upgrades.
type Aptitude string

// Cost return the cost of the aptitude. Implements Coster.
func (a Aptitude) Cost(u Universe, c Character) (int, error) {
	return 0, nil
}
