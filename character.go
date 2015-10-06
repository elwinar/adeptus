package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"strconv"

	"github.com/bradfitz/slice"
)

// Character is the type representing a role playing character
type Character struct {
	Name            string
	Backgrounds     map[string]Background
	Aptitudes       map[string]Aptitude
	Characteristics map[string]Characteristic
	Skills          map[string]Skill
	Talents         map[string]Talent
	Gauges          map[string]Gauge
	Rules           map[string]Rule
	Experience      int
	Spent           int
}

// CountMatchingAptitudes return the number of aptitudes of the given slice
// that are in the character's aptitudes.
func (c Character) CountMatchingAptitudes(aptitudes []Aptitude) int {

	count := 0
	for _, aptitude := range aptitudes {
		if _, found := c.Aptitudes[string(aptitude)]; found {
			count++
		}
	}
	return count
}

// ApplyBackground changes the character's trait according to the history values
func (c *Character) ApplyBackground(b Background, u Universe) error {

	// For each upgrade associated to the history, apply each option.
	for _, upgrades := range b.Upgrades {
		for _, option := range upgrades {
			err := c.ApplyUpgrade(option, u)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// ApplyUpgrade changes the character's trait according to the given upgrade.
func (c *Character) ApplyUpgrade(upgrade Upgrade, universe Universe) error {
	
	var err error

	// Find the attribute corresponding to the upgrade, and initialize a new
	// rule if there isn't any.
	coster, found := universe.FindCoster(upgrade.Name)
	if !found {
		coster = Rule{
			Name: upgrade.Name,
		}
	}

	// If no cost is defined, compute it on the fly.
	if upgrade.Cost == nil {
		cost, err := coster.Cost(universe, *c)
		if err != nil {
			return err
		}
		
		upgrade.Cost = &cost
	}

	// Update the spent experience.
	c.Spent += *upgrade.Cost

	// Apply the upgrade depending on the target attribute.
	switch attribute := coster.(type) {
	case Characteristic:
		err = c.ApplyCharacteristicUpgrade(attribute, upgrade)
	case Skill:
		err = c.ApplySkillUpgrade(attribute, upgrade)
	case Talent:
		err = c.ApplyTalentUpgrade(attribute, upgrade)
	case Aptitude:
		err = c.ApplyAptitudeUpgrade(attribute, upgrade)
	case Gauge:
		err = c.ApplyGaugeUpgrade(attribute, upgrade)
	case Rule:
		err = c.ApplyRuleUpgrade(attribute, upgrade)
	}

	return err
}

func (c *Character) ApplyCharacteristicUpgrade(characteristic Characteristic, u Upgrade) error {

	// Check if the characteristic is defined in the character, and define it if not.
	if _, found := c.Characteristics[characteristic.Name]; !found {
		c.Characteristics[characteristic.Name] = characteristic
	}

	// Increment the tier if the mark is default.
	if u.Mark == MarkDefault {
		char := c.Characteristics[u.Name]
		char.Tier++
		c.Characteristics[u.Name] = char
	}

	// Parse the characteristic's upgrade value.
	raw := strings.TrimSpace(strings.TrimLeft(u.Name, characteristic.Name))
	value, err := strconv.Atoi(raw)
	if err != nil {
		return NewError(InvalidCharacteristicValue)
	}

	// Update the characteristic value.
	char := c.Characteristics[u.Name]
	if strings.HasPrefix(raw, "+") || strings.HasPrefix(raw, "-") {
		char.Value += value
	} else {
		char.Value = value
	}
	c.Characteristics[c.Name] = char

	return nil
}

func (c *Character) ApplySkillUpgrade(skill Skill, u Upgrade) error {

	// Check if the skill is defined in the character, and define it if not.
	if _, found := c.Skills[skill.FullName()]; !found {
		c.Skills[skill.FullName()] = skill
	}

	// Increment the tier if the mark is default.
	if u.Mark == MarkDefault {
		s := c.Skills[skill.FullName()]
		s.Tier++
		c.Skills[skill.FullName()] = s
	}

	return nil
}

func (c *Character) ApplyTalentUpgrade(talent Talent, u Upgrade) error {

	// A talent must have a default mark.
	if u.Mark != MarkDefault {
		return NewError(InvalidUpgradeMark, u.Line)
	}

	// Check if the talent is defined in the character, and define it if not.
	if _, found := c.Talents[talent.FullName()]; !found {
		c.Talents[talent.FullName()] = talent
	}

	// Increment the value of the talent.
	t := c.Talents[talent.FullName()]
	t.Value++
	c.Talents[talent.FullName()] = t

	return nil
}

func (c *Character) ApplyAptitudeUpgrade(aptitude Aptitude, u Upgrade) error {

	c.Aptitudes[string(aptitude)] = aptitude
	return nil
}

func (c *Character) ApplyGaugeUpgrade(gauge Gauge, u Upgrade) error {

	// A talent must have a default mark.
	if u.Mark != MarkDefault {
		return NewError(InvalidUpgradeMark, u.Line)
	}

	// Check if the gauge is defined in the character, and define it if not.
	if _, found := c.Gauges[gauge.Name]; !found {
		c.Gauges[gauge.Name] = gauge
	}

	// Parse the gauge's upgrade value.
	raw := strings.TrimSpace(strings.TrimLeft(u.Name, gauge.Name))
	value, err := strconv.Atoi(raw)
	if err != nil {
		return NewError(InvalidGaugeValue)
	}

	// Update the gauge value.
	if strings.HasPrefix(raw, "+") || strings.HasPrefix(raw, "-") {
		c.Gauges[u.Name].Value += value
	} else {
		c.Gauges[c.Name].Value = up.Value
	}

	return nil
}

func (c *Character) ApplyRuleUpgrade(rule Rule, u Upgrade) error {

	// A talent must have a default mark.
	if u.Mark != MarkDefault {
		return NewError(InvalidUpgradeMark, u.Line)
	}

	// Check if the talent is defined in the character, and define it if not.
	if _, found := c.Rules[rule.Name]; !found {
		c.Rules[rule.Name] = rule
	}

	return nil
}

// Print the character sheet on the screen
func (c Character) Print() {
	// Print the name
	fmt.Printf("%s\t%s\n", theme.Title("Name"), c.Name)

	// Print the backgrounds
	types := []string{}
	for typ, _ := range c.Backgrounds {
		types = append(types, typ)
	}

	slice.Sort(types, func(i, j int) bool {
		return types[i] < types[j]
	})

	for _, typ := range types {
		backgrounds := []string{}

		for _, background := range c.Backgrounds[typ] {
			backgrounds = append(backgrounds, background.Name)
		}

		slice.Sort(backgrounds, func(i, j int) bool {
			return backgrounds[i] < backgrounds[j]
		})

		fmt.Printf("%s\t%s\n", theme.Title(strings.Title(typ)), strings.Title(strings.Join(backgrounds, ", ")))
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

/*
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
*/