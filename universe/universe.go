package universe

import (
	"fmt"
	"io"
	"io/ioutil"
)

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
}

// ParseUniverse load an universe from a plain JSON file. It returns a well-formed universe that describe all the components of a game setting.
func ParseUniverse(file io.Reader) (Universe, error) {
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return Universe{}, fmt.Errorf("unable to read universe: %s", err.Error())
	}

	universe := Universe{}
	err := json.Unmarshall(file, &universe)
	if err != nil {
		return Universe{}, fmt.Errorf("unable to parse universe: %s", err.Error())
	}

	return universe, nil
}
