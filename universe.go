package main

import (
	"encoding/json"
	"fmt"
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
	
	// Sort the universe slices.
	slice.Sort(universe.Aptitudes, func(i, j int) bool {
                return universe.Aptitudes[i] < universe.Aptitudes[j]
        })
        
        slice.Sort(universe.Characteristics, func(i, j int) bool {
                return universe.Characteristics[i] < universe.Characteristics[j]
        })
        
        slice.Sort(universe.Gauges, func(i, j int) bool {
                return universe.Gauges[i] < universe.Gauges[j]
        })
        
        slice.Sort(universe.Skills, func(i, j int) bool {
                return universe.Skills[i] < universe.Skills[j]
        })
        
        slice.Sort(universe.Talents, func(i, j int) bool {
                return universe.Talents[i] < universe.Talents[j]
        })
        
        for k, b := range universe.Backgrounds {
                slice.Sort(universe.Backgrounds[k], func(i, j int) bool {
                        return universe.Background[k][i] < universe.Background[k][j]
                }
        }

	// Check the aptitudes used in the universe.
	used := []Aptitude{}

	// Search for used aptitudes in Characteristics.
	for _, c := range universe.Characteristics {
		for _, a := range c.Aptitudes {
                        used = append(used, a)
		}
	}

	// Search for used aptitudes in Skills.
	for _, c := range universe.Skills {
		for _, a := range c.Aptitudes {
                        used = append(used, a)
		}
	}

	// Search for used aptitudes in Talents.
	for _, c := range universe.Talents {
		for _, a := range c.Aptitudes {
                        used = append(used, a)
		}
	}
        
	// Sort the used slice to compare it with the universe aptitudes slice.
	slice.Sort(used, func(i, j int) bool {
                return used[i] < used[j]
        })
        
        // Remove duplicates on used aptitudes.
        k := 0
        for k < len(used) -1 {
                
                if used[k] == used[k+1] {
                        used = append(used[:k], used[k+1:]...)
                        continue
                }
                k++
        }
        
        // Loop on both aptitudes slices.
        var i, j int
        for i < len(universe.Aptitudes) && j < len(used) {
            
                // The aptitude does not exist in both slices.
                if universe.Aptitudes[i] != universe.Aptitudes[j] {
                        break
                }
                i++
                j++
        }
        
        // Look for diffences in the slices.
        switch {
            // The universe uses more aptitudes than it defines.
            case len(universe.Aptitudes) < len(used):
                return Universe{}, NewError(UndefinedAptitude, used[len(used) -1])
                    
            // The aptitude is used but not defined.
            case universe.Aptitudes[i] > used[j]:
                    return Universe{}, NewError(UndefinedAptitude, used[j])
                
            // The universe defines some unused aptitudes.
            case len(universe.Aptitudes) > len(used):
                return Universe{}, NewError(UnusedAptitude, universe.Aptitudes[len(universe.Aptitudes) -1])
                       
            // The aptitude is defined used but is not used.
            case universe.Aptitudes[i] < used[j]:
                    return Universe{}, NewError(UnusedAptitude, universe.Aptitudes[i])
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

// FindBackground returns the history corresponding to the given label
func (u Universe) FindBackground(typ string, label string) (Background, bool, error) {

	histories, found := u.Backgrounds[typ]
	if !found {
		return Background{}, false, NewError(UndefinedBackgroud, typ)
	}

	for _, history := range histories {
		if history.Name == label {
			return history, true, nil
		}
	}

	return Background{}, false, nil
}
