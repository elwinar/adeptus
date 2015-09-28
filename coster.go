package main

// Coster is the interface implemented by an upgrade capable of being priced.
type Coster interface {
	Cost(CostMatrix, []Aptitude) (int, error)
}
