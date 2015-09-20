package main

import (
	"fmt"
	"strconv"

	"github.com/elwinar/adeptus/parser"
	"github.com/elwinar/adeptus/universe"
)

// Character is the type representing a role playing character
type Character struct {
	Name            string
	Aptitudes       []universe.Aptitude
	Origin          universe.Origin
	Background      universe.Background
	Role            universe.Role
	Tarot           universe.Tarot
	Characteristics map[*universe.Characteristic]int
	Skills			map[*universe.Skill]int
	Talents			map[*universe.Talent]int
	Gauges			map[*universe.Gauge]int
	Rules			[]universe.Rule
}

// NewCharacter creates a new character given a sheet
func NewCharacter(u universe.Universe, s parser.Sheet) (*Character, error) {
	
	// Create the character.
	c := &Character{}

	// Apply the initial characteristics from the sheet.
	c.Characteristics = make(map[*universe.Characteristic]int)
	for _, characteristic := range s.Characteristics {

		// Identify name and value.
		name, value, _, err := IdentifyCharacteristic(characteristic.Name)
		if err != nil {
			return nil, err
		}

		// Retrieve characteristic from universe given it's name.
		char, found := u.FindCharacteristic(name)
		if !found {
			return nil, fmt.Errorf("undefined characteristic %s", name)
		}

		// Check the characteristic is not set twice
		_, found = c.Characteristics[&char]
		if found {
			return nil, fmt.Errorf("characteristic %s previously defined in character sheet", name)
		}

		// Associate the characteristic and it' value to the characteristics map
		c.Characteristics[&char] = value
	}
	
	// Check all characteristics from universe are defined for the character
	checkCharacteristics:
	for _, u := range u.Characteristics {
		for c, _ := range c.Characteristics {
			if c.Name == u.Name {
				continue checkCharacteristics
			}
		}
		return nil, fmt.Errorf("charactersitic %s of universe not defined for character", u.Name)
	}
	
	// Make the character's gauges, skills and talents maps
	c.Skills = make(map[*universe.Skill]int)
	c.Talents = make(map[*universe.Talent]int)
	c.Gauges = make(map[*universe.Gauge]int)

	h := s.Header

	// Retrieve the name from the sheet header.
	if len(h.Name) == 0 {
		return nil, fmt.Errorf("empty name")
	}
	c.Name = h.Name
	
	// Apply each Meta
	var err error

	// Retrieve the origin from the universe.
	if h.Origin == nil {
		return nil, fmt.Errorf("unspecified origin")
	}
	origin, found := u.FindOrigin(h.Origin.Label)
	if !found {
		return nil, fmt.Errorf("origin %s not found", h.Origin.Label)
	}
	err = c.ApplyHistory(origin, u)
	if err != nil {
		return nil, err
	}

	// Retrieve the background from the universe.
	if h.Background == nil {
		return nil, fmt.Errorf("unspecified background")
	}
	background, found := u.FindBackground(h.Background.Label)
	if !found {
		return nil, fmt.Errorf("background %s not found", h.Background.Label)
	}
	err = c.ApplyHistory(background, u)
	if err != nil {
		return nil, err
	}

	// Retrieve the role from the universe.
	if h.Role == nil {
		return nil, fmt.Errorf("unspecified role")
	}
	role, found := u.FindRole(h.Role.Label)
	if !found {
		return nil, fmt.Errorf("role %s not found", h.Role.Label)
	}
	err = c.ApplyHistory(role, u)
	if err != nil {
		return nil, err
	}

	// Retrieve the tarot from the universe.
	var tarot universe.History
	if h.Tarot == nil {
		return nil, fmt.Errorf("unspecified tarot")
	}

	dice, err := strconv.Atoi(h.Tarot.Label)
	if err == nil {
		tarot, found = u.FindTarotByDice(dice)
	} else {
		tarot, found = u.FindTarot(h.Tarot.Label)
	}
	if !found {
		return nil, fmt.Errorf("tarot %s not found", h.Tarot.Label)
	}
	err = c.ApplyHistory(tarot, u)
	if err != nil {
		return nil, err
	}

	// Apply the sessions.

	return c, nil
}

// Debug prints the current values of the character
func (c Character) Debug() {
	fmt.Printf("Name		%s\n", c.Name)
	fmt.Printf("Origin		%s\n", c.Origin.Name)
	fmt.Printf("Background	%s\n", c.Background.Name)
	fmt.Printf("Role		%s\n", c.Role.Name)
	fmt.Printf("Tarot		%s\n", c.Tarot.Name)
	fmt.Printf("\nCharacteristics\n")
	for c, value := range c.Characteristics {
		fmt.Printf("%s		%d\n", c.Name, value)
	}
	fmt.Printf("\nAptitudes\n")
	for _, a := range c.Aptitudes {
		fmt.Printf("		%s\n", a)
	}
	fmt.Printf("\nTalents\n")
	for t, v := range c.Talents {
		fmt.Printf("%s		%d\n", t.Name, v)
	}
	fmt.Printf("\nSkills\n")
	for s, v := range c.Skills {
		fmt.Printf("%s		%d\n", s.Name, v)
	}
	fmt.Printf("\nRules\n")
	for _, r := range c.Rules {
		fmt.Printf("%s		%s\n", r.Name, r.Description)
	}
}

// ApplyHistory changes the character's trait according to the meta values
func (c *Character) ApplyHistory(h universe.History, u universe.Universe) error {
	
	// Attach the history to the proper character's meta.
	switch t := h.(type) {
		case universe.Origin:
			c.Origin = t
		case universe.Background:
			c.Background = t
		case universe.Role:
			c.Role = t
		case universe.Tarot:
			c.Tarot = t
	}
	
	// For each upgrade associated to the meta, apply each option.
	for _, upgrades := range h.GetUpgrades() {
		for _, option := range upgrades {
			err := c.ApplyUpgrade(option, u)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ApplyUpgrade changes the character's trait according to the given upgrade
func (c *Character) ApplyUpgrade(up parser.Upgrade, un universe.Universe) error {
			
	// Identify characteristic.
	name, value, sign, err := IdentifyCharacteristic(up.Name)
	if err == nil {
		
		// Check the characteristic exists.
		characteristic, found := un.FindCharacteristic(name)
		if !found {
			return fmt.Errorf(`universe provides upgrade "%s" but does not define the charactersitic "%s"`, up.Name, name)
		}
		
		// Apply the modification to the character's characteristic.
		// The characteristic must have this characteristic
		for char, v := range c.Characteristics {
			if char.Name == characteristic.Name {
				c.Characteristics[char] = ApplyCharacteristicUpgrade(v, sign, value)
				return nil
			}
		}
		panic(fmt.Sprintf("unidefined characteristic %s for character", characteristic.Name))
	}
	
	// Identify aptitude.
	aptitude, found := un.FindAptitude(up.Name)
	if found {
		c.Aptitudes = append(c.Aptitudes, aptitude)
		return nil
	}
	
	// Identify skill or talent.
	name, speciality, err := SplitUpgrade(up.Name)
	if err != nil {
		return err
	}
		
	// Skill identified.
	skill, isSkill := un.FindSkill(name)
	if isSkill {
			
		// The skill has a speciality
		if len(speciality) != 0 {
			skill.Name = fmt.Sprintf("%s: %s", name, speciality)
		}
	
		// Look for the skill with the given name in the character's skill,
		// and apply the modification
		for s := range c.Skills {
			if s.Name == skill.Name {
				
				// Change the value of the skill of index s
				c.Skills[s] += 10
				return nil
			}
		}
		// Create the skill of index *skill
		c.Skills[&skill] = 0
		return nil
	}
	
	// Talent identified.
	talent, isTalent := un.FindTalent(name)
	if isTalent {
			
		// The talent has a speciality
		if len(speciality) != 0 {
			talent.Name = fmt.Sprintf("%s: %s", name, speciality)
		}
	
		// Look for the talent with the given name in the character's talents,
		// and apply the modification
		for s := range c.Talents {
			if s.Name == talent.Name {
				
				// Change the value of the talent of index s
				c.Talents[s]++
				return nil
			}
		}
		// Create the talent of index *talent
		c.Talents[&talent] = 1
		return nil
	}
	
	// It's a special rule
	rule := universe.Rule{
		Name: name,
		Description: speciality,
	}
	c.Rules = append(c.Rules, rule)
	
	return nil
}