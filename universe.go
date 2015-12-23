package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// Universe represents a set of configuration, often refered as data or database.
type Universe struct {
	Backgrounds     map[string][]Background `json:"backgrounds"`
	Aptitudes       []Aptitude              `json:"aptitudes"`
	Characteristics []Characteristic        `json:"characteristics"`
	Gauges          []Gauge                 `json:"gauges"`
	Skills          []Skill                 `json:"skills"`
	Talents         []Talent                `json:"talents"`
	Costs           CostMatrix              `json:"costs"`
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

	// Lowercase the types of background.
	backgrounds := make(map[string][]Background)
	for typ, b := range universe.Backgrounds {
		backgrounds[strings.ToLower(typ)] = b
	}
	universe.Backgrounds = backgrounds

	// Add it's type to each defined background.
	for typ, backgrounds := range universe.Backgrounds {
		for i := range backgrounds {
			universe.Backgrounds[typ][i].Type = typ
		}
	}

	return universe, nil
}

// FindCoster returns the coster associated to the label,
// and false if none is.
func (u Universe) FindCoster(upgrade Upgrade) (Coster, bool) {
	characteristic, found := u.FindCharacteristic(upgrade)
	if found {
		return characteristic, true
	}

	skill, found := u.FindSkill(upgrade)
	if found {
		return skill, true
	}

	talent, found := u.FindTalent(upgrade)
	if found {
		return talent, true
	}

	aptitude, found := u.FindAptitude(upgrade)
	if found {
		return aptitude, true
	}

	gauge, found := u.FindGauge(upgrade)
	if found {
		return gauge, true
	}

	return nil, false
}

// FindCharacteristic returns the characteristic correponding to the given label or a zero-value, and a boolean indicating if it was found.
func (u Universe) FindCharacteristic(upgrade Upgrade) (Characteristic, bool) {

	// Characteristics upgrades are defined by a name and a value, separated by a space, so we need to look for the first
	// part of the label.
	// Examples: STR +5, FEL -1, TOU 40.
	fields := split(upgrade.Name, ' ')
	name := fields[0]

	for _, characteristic := range u.Characteristics {
		if characteristic.Name == name {
			return characteristic, true
		}
	}

	return Characteristic{}, false
}

// FindSkill returns the skill corresponding to the given label or a zero-value, and a boolean indicating if it was found.
func (u Universe) FindSkill(upgrade Upgrade) (Skill, bool) {

	// Skills upgrades are defined by a name and eventually a speciality, separated by a colon.
	// Examples: Common Lore: Dark Gods
	fields := split(upgrade.Name, ':')
	name := fields[0]

	for _, skill := range u.Skills {
		if strings.EqualFold(skill.Name, name) {

			if len(fields) == 2 {
				skill.Speciality = fields[1]
			}

			return skill, true
		}
	}

	return Skill{}, false
}

// FindTalent returns the talent corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindTalent(upgrade Upgrade) (Talent, bool) {

	// Talents upgrades are defined by a name and eventually a speciality, separated by a colon.
	// Examples: Psychic Resistance: Fear
	fields := split(upgrade.Name, ':')
	name := fields[0]

	for _, talent := range u.Talents {
		if strings.EqualFold(talent.Name, name) {

			if len(fields) == 2 {
				talent.Speciality = fields[1]
			}

			return talent, true
		}
	}

	return Talent{}, false
}

// FindAptitude returns the aptitude corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindAptitude(upgrade Upgrade) (Aptitude, bool) {

	for _, aptitude := range u.Aptitudes {
		if strings.EqualFold(string(aptitude), upgrade.Name) {
			return aptitude, true
		}
	}

	return Aptitude(""), false
}

// FindGauge returns the gauge corresponding to the given label or a zero value, and a boolean indicating if it was found.
func (u Universe) FindGauge(upgrade Upgrade) (Gauge, bool) {

	// Gauges upgrades are defined by a name and a value, separated by a space.
	fields := split(upgrade.Name, ' ')

	for _, gauge := range u.Gauges {
		if strings.EqualFold(gauge.Name, fields[0]) {
			val, err := strconv.Atoi(fields[1])
			if err != nil {
				panic(err)
			}
			gauge.Value = val
			return gauge, true
		}
	}

	return Gauge{}, false
}

// FindBackground returns the background corresponding to the given label, and a boolean indicating if it was found.
func (u Universe) FindBackground(typ string, label string) (Background, bool) {

	backgrounds, found := u.Backgrounds[strings.ToLower(typ)]
	if !found {
		return Background{}, false
	}

	for _, background := range backgrounds {
		if strings.EqualFold(background.Name, label) {
			return background, true
		}
	}

	return Background{}, false
}
