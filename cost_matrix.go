package main

import (
	"strconv"
	
	"gopkg.in/yaml.v2"
)

// CostMatrix is matrix of upgrade costs.
type CostMatrix map[string]map[int]map[int]int

// costMatrixYAML is the Json representation of the matrix.
type costMatrixYAML map[string]map[string]map[string]int

// Price returns the cost in the matrix corresponding to the given type, matches and tier.
func (c CostMatrix) Price(typ string, matches int, tier int) (int, error) {

	var f bool

	// Check if the upgrade type is listed in the matrix.
	_, f = c[typ]
	if !f {
		return 0, NewError(UndefinedTypeCost, typ)
	}

	// Check if the number of matching aptitudes is in the matrix.
	_, f = c[typ][matches]
	if !f {
		return 0, NewError(UndefinedMatchCost, typ, matches)
	}

	// Get the cost corresponding to the upgrade's tier.
	cost, f := c[typ][matches][tier]
	if !f {
		return 0, NewError(UndefinedTierCost, typ, matches, tier)
	}

	return cost, nil
}

// MarshalYAML return the YAML representation of the cost matrix.
// Implements the Marshaller interface.
func (c *CostMatrix) MarshalYAML() ([]byte, error) {

	jMatrix := costMatrixYAML{}

	// For each upgrade type of the matrix.
	for typ, matches := range *c {

		// Instanciate the correponding map of matching aptitudes.
		jMatrix[typ] = make(map[string]map[string]int)

		// For each number of matching aptitudes.
		for match, tiers := range matches {

			// Instanciate the map of tiers costs.
			m := strconv.Itoa(match)
			jMatrix[typ][m] = make(map[string]int)

			// For each cost.
			for tier, cost := range tiers {

				// Add the corresponding cost to the transient matrix.
				t := strconv.Itoa(tier)
				jMatrix[typ][m][t] = cost
			}
		}
	}

	// Marshal the transient structure to the returned YAML object.
	raw, err := yaml.Marshal(jMatrix)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// UnmarshalYAML parse the YAML representation of a cost matrix.
// Implements the Unmarshaller interface.
func (c *CostMatrix) UnmarshalYAML(raw []byte) error {

	// Unmarshal to the transient structure.
	jMatrix := costMatrixYAML{}
	err := yaml.Unmarshal(raw, &jMatrix)
	if err != nil {
		return err
	}

	// Initialize the matrix map.
	tmp := make(map[string]map[int]map[int]int)

	// For each upgrade type of the matrix.
	for typ, matches := range jMatrix {

		// Initialize the map of matching aptitudes.
		tmp[typ] = make(map[int]map[int]int)

		// For each number of matching aptitudes.
		for match, tiers := range matches {

			// Initialize the map of tiers costs.
			m, err := strconv.Atoi(match)
			if err != nil {
				return err
			}
			tmp[typ][m] = make(map[int]int)

			// For each cost.
			for tier, cost := range tiers {

				// Add the corresponding cost to the matrix.
				t, err := strconv.Atoi(tier)
				if err != nil {
					return err
				}
				tmp[typ][m][t] = cost
			}
		}
	}

	*c = CostMatrix(tmp)

	return nil
}
