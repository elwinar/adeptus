package universe

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// Universe represents a set of configuration, often refered as data or database.
type Universe struct {
	Origins         []Origin
	Backgrounds     []Background
	Roles           []Role
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

// FindOrigin returns the origin corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindOrigin(label string) (Origin, bool) {

	for _, origin := range u.Origins {
		if origin.Name == label {
			return origin, true
		}
	}

	return Origin{}, false
}

// FindBackground returns the background corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindBackground(label string) (Background, bool) {

	for _, background := range u.Backgrounds {
		if background.Name == label {
			return background, true
		}
	}

	return Background{}, false
}

// FindRole returns the role corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindRole(label string) (Role, bool) {

	for _, role := range u.Roles {
		if role.Name == label {
			return role, true
		}
	}

	return Role{}, false
}

// FindTarot returns the tarot corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindTarot(label string) (Tarot, bool) {

	for _, tarot := range u.Tarots {
		if tarot.Name == label {
			return tarot, true
		}
	}

	return Tarot{}, false
}

// FindTarotByDice returns the tarot corresponding to the given value or a zero value, and a boolean indicating if a tarot exist for this dice.
func (u Universe) FindTarotByDice(dice int) (Tarot, bool) {

	for _, tarot := range u.Tarots {
		if tarot.Min <= dice && dice <= tarot.Max {
			return tarot, true
		}
	}

	return Tarot{}, false
}
