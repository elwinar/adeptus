package universe

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// Universe represents a set of configuration, often refered as data or database.
type Universe struct {
	Histories       map[string][]History
	Tarots          []Tarot
	Aptitudes       []Aptitude
	Characteristics []Characteristic
	Gauges          []Gauge
	Skills          []Skill
	Talents         []Talent
	Costs           CostMatrix
}

// ParseUniverse load an universe from a plain JSON file.
// It returns a well-formed universe that describe all the components of a game setting.
func ParseUniverse(file io.Reader) (Universe, error) {

	// open and parse universe
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return Universe{}, fmt.Errorf("unable to read universe: %s", err.Error())
	}

	universe := Universe{}
	err = json.Unmarshal(raw, &universe)
	if err != nil {
		return Universe{}, fmt.Errorf("unable to parse universe: %s", err.Error())
	}

	// Check the aptitudes in Skills, Characteristics and Talents are defined in the universe.
	observed := make(map[Aptitude]struct{})

	// For each aptitude of each characteristic.
	for _, c := range universe.Characteristics {
		for _, a := range c.Aptitudes {

			// Add the aptitude to the slice of observed aptitudes.
			_, f := observed[a]
			if !f {
				observed[a] = struct{}{}
			}
		}
	}

	// For each aptitude of each skill.
	for _, c := range universe.Skills {
		for _, a := range c.Aptitudes {

			// Add the aptitude to the slice of observed aptitudes.
			_, f := observed[a]
			if !f {
				observed[a] = struct{}{}
			}
		}
	}

	// For each aptitude of each talent.
	for _, c := range universe.Talents {
		for _, a := range c.Aptitudes {

			// Add the aptitude to the slice of observed aptitudes.
			_, f := observed[a]
			if !f {
				observed[a] = struct{}{}
			}
		}
	}

	// Check all aptitudes defined by universe are used at least once.
checkDefined:
	for _, a := range universe.Aptitudes {
		for o := range observed {
			if a == o {
				continue checkDefined
			}
		}
		return Universe{}, fmt.Errorf("aptitude %s defined by universe but not used", a)
	}

	// Check all aptitudes defined by universe are used at least once.
checkObserved:
	for o := range observed {
		for _, a := range universe.Aptitudes {
			if a == o {
				continue checkObserved
			}
		}
		return Universe{}, fmt.Errorf("aptitude %s used by universe but not defined", o)
	}

	return universe, nil
}

// FindCharacteristic returns the characteristic correponding to the given label or a zero-value, and a boolean indicating if it was found.
func (u Universe) FindCharacteristic(label string) (Characteristic, bool) {

	for _, characteristic := range u.Characteristics {
		if characteristic.Name == label {
			return characteristic, true
		}
	}

	return Characteristic{}, false
}

// FindSkill returns the skill corresponding to the given label or a zero-value, and a boolean indicating if it was found.
func (u Universe) FindSkill(label string) (Skill, bool) {

	for _, skill := range u.Skills {
		if skill.Name == label {
			return skill, true
		}
	}

	return Skill{}, false
}

// FindTalent returns the talent corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindTalent(label string) (Talent, bool) {

	for _, talent := range u.Talents {
		if talent.Name == label {
			return talent, true
		}
	}

	return Talent{}, false
}

// FindAptitude returns the aptitude corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindAptitude(label string) (Aptitude, bool) {

	for _, aptitude := range u.Aptitudes {
		if string(aptitude) == label {
			return aptitude, true
		}
	}

	return Aptitude(""), false
}

// FindHistory returns the history corresponding to the given label
func (u Universe) FindHistory(typ string, label string) (History, bool, error) {

	histories, found := u.Histories[typ]
	if !found {
		return History{}, false, fmt.Errorf("undefined history type %s in universe", typ)
	}

	for _, history := range histories {
		if history.Name == label {
			return history, true, nil
		}
	}

	return History{}, false, nil
}

// FindTarot returns the tarot corresponding to the given value or a zero value, and a boolean indicating if a tarot exist for this dice.
func (u Universe) FindTarot(dice int) (Tarot, bool) {

	for _, tarot := range u.Tarots {
		if tarot.Min <= dice && dice <= tarot.Max {
			return tarot, true
		}
	}

	return Tarot{}, false
}
