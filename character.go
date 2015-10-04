package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/bradfitz/slice"
)

// Character is the type representing a role playing character
type Character struct {
	Name            string
	Backgrounds     map[string][]Background
	Aptitudes       []Aptitude
	Characteristics map[*Characteristic]int
	Skills          map[*Skill]int
	Talents         map[*Talent]int
	Gauges          map[*Gauge]int
	Rules           []Rule
	Experience      int
	Spent           int
}

// NewCharacter creates a new character given a sheet
func NewCharacter(u Universe, s Sheet) (*Character, error) {

	// Create the character.
	c := &Character{}

	// Alias Header.
	h := s.Header

	// Retrieve character's name.
	c.Name = h.Name

	// Apply the initial characteristics from the sheet.
	c.Characteristics = make(map[*Characteristic]int)
	for _, characteristic := range s.Characteristics {

		// Identify name and value.
		name, value, _, err := IdentifyCharacteristic(characteristic)
		if err != nil {
			return nil, err
		}

		// Retrieve characteristic from universe given it's name.
		char, found := u.FindCharacteristic(name)
		if !found {
			return nil, NewError(UndefinedCharacteristic, characteristic.Line)
		}

		// Check the characteristic is not set twice
		_, found = c.Characteristics[&char]
		if found {
			return nil, NewError(DuplicateCharacteristic, characteristic.Line)
		}

		// Associate the characteristic and it' value to the characteristics map
		c.Characteristics[&char] = value
	}

	// Check all characteristics are defined for the character
checkCharacteristics:
	for _, u := range u.Characteristics {
		for c := range c.Characteristics {
			if c.Name == u.Name {
				continue checkCharacteristics
			}
		}
		return nil, NewError(MissingCharacteristic, u.Name)
	}

	// Make the character's gauges, skills and talents maps
	c.Skills = make(map[*Skill]int)
	c.Talents = make(map[*Talent]int)
	c.Gauges = make(map[*Gauge]int)

	// Apply each Meta.
	c.Backgrounds = make(map[string][]Background)
	for typ, metas := range h.Metas {

		if len(metas) == 0 {
			panic(fmt.Sprintf("empty metas for type %s", typ))
		}
		line := metas[0].Line

		bagrounds, found := u.Backgrounds[typ]

		// Check the background type exists in universe.
		if !found {
			return nil, NewError(UndefinedBackgroundType, line, typ)
		}

		c.Backgrounds[typ] = []Background{}

	metasLoop:
		for _, meta := range metas {

			// Search the background corresponding to the provided meta.
			for _, b := range bagrounds {
				if meta.Label != b.Name {
					continue
				}

				// Apply the background.
				c.Backgrounds[typ] = append(c.Backgrounds[typ], b)
				err := c.ApplyBackground(b, u)
				if err != nil {
					return nil, err
				}

				continue metasLoop
			}
			return nil, NewError(UndefinedBackgroundValue, line, meta.Label, typ)
		}
	}

	// Apply the sessions.
	for _, s := range s.Sessions {

		// Add experience value.
		if s.Reward != nil {
			c.Experience += *s.Reward
		}

		// For each upgrade.
		for _, up := range s.Upgrades {

			// Apply the upgrade.
			err := c.ApplyUpgrade(up, u)
			if err != nil {
				return nil, err
			}
		}
	}

	return c, nil
}

// ApplyBackground changes the character's trait according to the history values
func (c *Character) ApplyBackground(h Background, u Universe) error {

	// For each upgrade associated to the history, apply each option.
	for _, upgrades := range h.Upgrades {
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
func (c *Character) ApplyUpgrade(up Upgrade, un Universe) error {

	// Defer payment of upgrade
	var payed bool
	var coster Coster

	// Pay the upgrade.
	switch {

	// Upgrade is free.
	case up.Mark == "-":
		if up.Cost != nil {
			return NewError(MismatchMarkCost, up.Line)
		}
		payed = true

	// Cost is hard defined.
	case up.Cost != nil:
		c.Spent += *up.Cost
		payed = true
	}

	// Identify characteristic.
	name, value, sign, err := IdentifyCharacteristic(up)
	if err == nil {

		// Check the characteristic exists.
		characteristic, found := un.FindCharacteristic(name)
		if !found {
			return NewError(UndefinedCharacteristic, up.Line)
		}

		// Apply the modification to the character's characteristic.
		// The characteristic must have this characteristic
		for char, v := range c.Characteristics {
			if char.Name == characteristic.Name {

				// Increment the characteristic Tier.
				if up.Mark == "*" {
					char.Tier++
				}

				// Apply the characteristic.
				c.Characteristics[char] = ApplyCharacteristicUpgrade(v, sign, value)
				coster = char

				// Pay for it
				if !payed && coster != nil && err == nil {
					var cost int
					// Transmit the error value to the parent func
					cost, err = coster.Cost(un.Costs, c.Aptitudes)
					if err != nil {
						return err
					}
					c.Spent += cost
				}
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

	// Sort aptitudes by name and remove duplicates.
	slice.Sort(c.Aptitudes, func(i, j int) bool {
		return c.Aptitudes[i] < c.Aptitudes[j]
	})
	var aptitudes []Aptitude
	for _, a := range c.Aptitudes {
		if len(aptitudes) != 0 && aptitudes[len(aptitudes)-1] == a {
			continue
		}
		aptitudes = append(aptitudes, a)
	}
	c.Aptitudes = aptitudes

	// Identify skill or talent.
	name, speciality, err := up.Split()
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

		// Look for the skill with the given name in the
		// character's skill, and apply the modification
		for s := range c.Skills {
			if s.Name == skill.Name {

				// Increment the skill tier
				if up.Mark == "*" {
					s.Tier++
				}

				// Change the value of the skill of index s
				c.Skills[s] += 10
				coster = s

				// Pay for it
				if !payed && coster != nil && err == nil {
					var cost int
					// Transmit the error value to the parent func
					cost, err = coster.Cost(un.Costs, c.Aptitudes)
					if err != nil {
						return err
					}
					c.Spent += cost
				}
				return nil
			}
		}

		// Increment the skill tier
		if up.Mark == "*" {
			skill.Tier++
		}

		// Create the skill of index *skill
		c.Skills[&skill] = 0
		coster = &skill

		// Pay for it
		if !payed && coster != nil && err == nil {
			var cost int
			// Transmit the error value to the parent func
			cost, err = coster.Cost(un.Costs, c.Aptitudes)
			if err != nil {
				return err
			}
			c.Spent += cost
		}
		return nil
	}

	// Talent identified.
	talent, isTalent := un.FindTalent(name)
	if isTalent {

		// The talent has a speciality
		if len(speciality) != 0 {
			talent.Name = fmt.Sprintf("%s: %s", name, speciality)
		}

		// Look for the talent with the given name in
		// the character's talents, and apply the modification
		for s := range c.Talents {
			if s.Name == talent.Name {

				// Change the value of the talent of index s
				c.Talents[s]++
				coster = s

				// Pay for it
				if !payed && coster != nil && err == nil {
					var cost int
					// Transmit the error value to the parent func
					cost, err = coster.Cost(un.Costs, c.Aptitudes)
					if err != nil {
						return err
					}
					c.Spent += cost
				}
				return nil
			}
		}

		// Create the talent of index *talent
		c.Talents[&talent] = 1
		coster = &talent

		// Pay for it
		if !payed && coster != nil && err == nil {
			var cost int
			// Transmit the error value to the parent func
			cost, err = coster.Cost(un.Costs, c.Aptitudes)
			if err != nil {
				return err
			}
			c.Spent += cost
		}
		return nil
	}

	// It's a special rule
	rule := Rule{
		Name:        name,
		Description: speciality,
	}
	c.Rules = append(c.Rules, rule)

	return nil
}

// Print the character sheet on the screen
func (c Character) Print() {
	// Print the name
	fmt.Printf("%s\t%s\n", theme.Title("Name"), c.Name)

	// Print the histories
	histories := []History{}
	for _, histories_ := range c.Histories {
		histories = append(histories, histories_...)
	}

	slice.Sort(histories, func(i, j int) bool {
		if histories[i].Type != histories[j].Type {
			return histories[i].Type < histories[j].Type
		}

		return histories[i].Name < histories[j].Name
	})

	for _, history := range histories {
		fmt.Printf("%s\t%s\n", theme.Title(strings.Title(history.Type)), strings.Title(history.Name))
	}

	// Print the experience
	fmt.Printf("\n%s\t%d/%d\n", theme.Title("Experience"), c.Spent, c.Experience)

	// Print the characteristics
	fmt.Printf("\n%s\n", theme.Title("Characteristics"))

	characteristics := []*Characteristic{}
	for characteristic, _ := range c.Characteristics {
		characteristics = append(characteristics, characteristic)
	}

	slice.Sort(characteristics, func(i, j int) bool {
		return characteristics[i].Name < characteristics[j].Name
	})

	for _, characteristic := range characteristics {
		fmt.Printf("%s\t%s\n", characteristic.Name, theme.Value(c.Characteristics[characteristic]))
	}

	// Print the talents
	fmt.Printf("\n%s\n", theme.Title("Talents"))

	talents := []*Talent{}
	for talent, _ := range c.Talents {
		talents = append(talents, talent)
	}

	slice.Sort(talents, func(i, j int) bool {
		return talents[i].Name < talents[j].Name
	})

	for _, talent := range talents {
		if c.Talents[talent] != 1 {
			fmt.Printf("%s (%s)\n", strings.Title(talent.Name), theme.Value(c.Talents[talent]))
		} else {
			fmt.Printf("%s\n", strings.Title(talent.Name))
		}
	}

	// Print the skills using a tabwriter
	fmt.Printf("\n%s\n", theme.Title("Skills"))

	skills := []*Skill{}
	for skill, _ := range c.Skills {
		skills = append(skills, skill)
	}

	slice.Sort(skills, func(i, j int) bool {
		return skills[i].Name < skills[j].Name
	})

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 2, ' ', 0)
	for _, skill := range skills {
		fmt.Fprintf(w, "%s\t%s\n", strings.Title(skill.Name), theme.Value(c.Skills[skill]))
	}
	w.Flush()

	// Print the special rules
	fmt.Printf("\n%s\n", theme.Title("Rules"))

	rules := c.Rules
	slice.Sort(rules, func(i, j int) bool {
		return rules[i].Name < rules[j].Name
	})

	for _, rule := range rules {
		fmt.Printf("%s\t%s\n", strings.Title(rule.Name), rule.Description)
	}
}
