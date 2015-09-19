package universe

// Coster is implemented by an upgrade capable of being priced
type Coster interface {
	// Cost returns the price of the upgrade to purchase
	Cost(CostMatrix, []Aptitude) (int, error)
}
