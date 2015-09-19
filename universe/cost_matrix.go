package universe

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// CostMatrix is matrix of upgrade costs
type CostMatrix struct {
	Fares map[string]map[int]map[int]int
}

// costMatrixJSON is the Json representation of the matrix
type costMatrixJSON map[string]map[string]map[string]int

// Price returns the cost in the matrix corresponding to the given type, matches and tier
func (c CostMatrix) Price(tpe string, matches int, tier int) (int, error) {

	var f bool

	// retrieve type matrix
	_, f = c.Fares[tpe]
	if !f {
		return 0, fmt.Errorf("undefined cost: %s", tpe)
	}

	// retrieve matches matrix
	_, f = c.Fares[tpe][matches]
	if !f {
		return 0, fmt.Errorf("undefined cost: %s x %d matches", tpe, matches)
	}

	// retrieve tiers matrix
	cost, f := c.Fares[tpe][matches][tier]
	if !f {
		return 0, fmt.Errorf("undefined cost: type %s x %d matches x %d tier", tpe, matches, tier)
	}

	return cost, nil
}

// MarshalJSON implements the Marshaller interface
func (c *CostMatrix) MarshalJSON() ([]byte, error) {

	jMatrix := costMatrixJSON{}

	// make types
	for tpe, matches := range c.Fares {
		jMatrix[tpe] = make(map[string]map[string]int)

		// make matches
		for match, tiers := range matches {
			m := strconv.Itoa(match)
			jMatrix[tpe][m] = make(map[string]int)

			// make tiers
			for tier, cost := range tiers {
				t := strconv.Itoa(tier)

				// make costs
				jMatrix[tpe][m][t] = cost
			}
		}
	}

	// transform json matrix to raw bytes
	raw, err := json.Marshal(jMatrix)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// UnmarshalJSON implements the Unmarshaller interface
func (c *CostMatrix) UnmarshalJSON(raw []byte) error {

	// parse as json matrix
	jMatrix := costMatrixJSON{}
	err := json.Unmarshal(raw, &jMatrix)
	if err != nil {
		return err
	}

	c.Fares = make(map[string]map[int]map[int]int)

	// make types
	for tpe, matches := range jMatrix {
		c.Fares[tpe] = make(map[int]map[int]int)

		// make matches
		for match, tiers := range matches {
			m, err := strconv.Atoi(match)
			if err != nil {
				return err
			}
			c.Fares[tpe][m] = make(map[int]int)

			// make tiers
			for tier, cost := range tiers {
				t, err := strconv.Atoi(tier)
				if err != nil {
					return err
				}

				// make costs
				c.Fares[tpe][m][t] = cost
			}
		}
	}

	return nil
}
