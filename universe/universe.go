package universe

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// Universe represents a set of configuration, often refered as data or database
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

// ParseUniverse load an universe from a plain JSON file. It returns a well-formed universe that describe all the components of a game setting.
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

// FindCharacteristic returns the characteristic related to the label and it's value, plus an error if none is found
func (u Universe) FindCharacteristic(label string) (Characteristic, string, error) {

	// harmonize charactersitic format
	splits := strings.Fields(label)

	// characteristic improperly formated
	if len(splits) != 2 {
		return Characteristic{}, "", fmt.Errorf("incorrect characteristic format %s", label)
	}

	// check the value is clean
	_, err := strconv.Atoi(splits[1])
	if err != nil {
		return Characteristic{}, "", fmt.Errorf("%s is not a correct characteristic value", splits[1])
	}

	// retrieve the characteristic
	l := strings.ToLower(splits[0])
	for _, c := range u.Characteristics {
		n := strings.ToLower(c.Name)

		// characteristic found, return it and it's stringy value
		if n == l {
			return c, splits[1], nil
		}
	}
	return Characteristic{}, "", fmt.Errorf("undefined characteristic %s", label)
}

// FindSkill returns the skill related to the label and it's speciality, plus an error if none is found
func (u Universe) FindSkill(label string) (Skill, string, error) {

	// retrieve label
	splits := strings.Split(label, ":")

	if len(splits) > 2 {
		return Skill{}, "", fmt.Errorf("incorrect skill format %s", label)
	}

	// retrieve the skill
	l := strings.ToLower(splits[0])
	for _, s := range u.Skills {
		n := strings.ToLower(s.Name)

		// skill not found
		if n != l {
			continue
		}

		// skill found, return it and it's stringy value
		var spec string
		if len(splits) == 2 {
			spec = splits[1]
		}

		return s, spec, nil
	}

	return Skill{}, "", fmt.Errorf("undefined skill %s", label)
}

// FindTalent returns the talent related to the label and it's speciality, plus an error if none is found
func (u Universe) FindTalent(label string) (Talent, string, error) {

	// retrieve label
	splits := strings.Split(label, ":")

	if len(splits) > 2 {
		return Talent{}, "", fmt.Errorf("incorrect talent format %s", label)
	}

	// retrieve the talent
	l := strings.ToLower(splits[0])
	for _, t := range u.Talents {
		n := strings.ToLower(t.Name)

		// talent not found
		if n != l {
			continue
		}

		// talent found, return it and it's stringy value
		var spec string
		if len(splits) == 2 {
			spec = splits[1]
		}

		return t, spec, nil
	}

	return Talent{}, "", fmt.Errorf("undefined talent %s", label)
}
