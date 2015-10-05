package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Universe represents a set of configuration, often refered as data or database.
type Universe struct {
	Backgrounds     map[string][]Background
	Aptitudes       []Aptitude
	Characteristics []Characteristic
	Gauges          []Gauge
	Skills          []Skill
	Talents         []Talent
	Costs           CostMatrix
}

// ParseUniverse load an from a plain JSON file.
// It returns a well-formed universe that describe all the components of a game setting.
func ParseUniverse(file io.Reader) (Universe, error) {

	// Open and parse universe.
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return Universe{}, err
	}
	universe := Universe{}
	err = json.Unmarshal(raw, &universe)
	if err != nil {
		return Universe{}, err
	}

	// Add the type value to each history defined.
	for typ, backgrounds := range universe.Backgrounds {
		for i, _ := range backgrounds {
			universe.Backgrounds[typ][i].Type = typ
		}
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

// FindBackground returns the background corresponding to the given label
func (u Universe) FindBackground(typ string, label string) (Background, error) {

	histories, found := u.Backgrounds[typ]
	if !found {
		return Background{}, NewError(UndefinedBackgroundType, typ)
	}

	for _, background := range histories {
		if background.Name == label {
			return background, nil
		}
	}

	return Background{}, NewError(UndefinedBackgroundValue, typ, label)
}
