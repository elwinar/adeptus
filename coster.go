package main

// Coster is the interface implemented by an upgrade capable of being priced.
type Coster interface {
	Cost(Universe, Character) (int, error)
}
