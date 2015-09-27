package universe

// Rule represent a special rule, which are generally home-made additions to the universe.
type Rule struct {
	Name        string
	Description string
}

// Cost implements the Coster interface.
func (r Rule) Cost(matrix CostMatrix, aptitudes []Aptitude) (int, error) {
	return 0, nil
}
