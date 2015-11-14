package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

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

// NewCharacter creates a new character from the given sheet and universe.
func NewCharacter(universe Universe, sheet Sheet) (Character, error) {

	// Create a character
	c := Character{
		Name:            sheet.Header.Name,
		Backgrounds:     make(map[string]Background),
		Aptitudes:       make(map[string]Aptitude),
		Characteristics: make(map[string]Characteristic),
		Skills:          make(map[string]Skill),
		Talents:         make(map[string]Talent),
		Gauges:          make(map[string]Gauge),
		Rules:           make(map[string]Rule),
		Experience:      0,
		Spent:           0,
	}

	// The characteristics described in the header of the sheet are parsed as upgrades
	for _, upgrade := range sheet.Characteristics {

		// Get the characteristic from the universe
		characteristic, found := universe.FindCharacteristic(upgrade.Name)
		if !found {
			return c, NewError(UndefinedCharacteristic, upgrade.Line)
		}

		// Check it is not already applied
		_, found = c.Characteristics[characteristic.Name]
		if found {
			return c, NewError(DuplicateCharacteristic, upgrade.Line)
		}

		// Apply the upgrade
		err := c.ApplyCharacteristicUpgrade(characteristic, upgrade)
		if err != nil {
			return c, err
		}
	}

	// Next are the backgrounds
	for typ, metas := range sheet.Header.Metas {

		for _, meta := range metas {

			// Find the background corresponding to the meta
			background, found := universe.FindBackground(typ, meta.Label)
			if !found {
				return c, NewError(UndefinedBackground, meta.Line, typ, meta.Label)
			}

			err := c.ApplyBackground(background, universe)
			if err != nil {
				return c, err
			}
		}
	}

	// Next are the sessions
	for _, session := range sheet.Sessions {

		// Apply the experience gain if needed
		if session.Reward != nil {
			c.Experience += *session.Reward
		}

		// Apply each upgrade in order
		for _, upgrade := range session.Upgrades {
			err := c.ApplyUpgrade(upgrade, universe)
			if err != nil {
				return c, err
			}
		}
	}

	return c, nil
}

// CountMatchingAptitudes return the number of aptitudes of the given slice
// that are in the character's aptitudes.
func (character Character) CountMatchingAptitudes(aptitudes []Aptitude) int {

	count := 0
	for _, aptitude := range aptitudes {
		if _, found := character.Aptitudes[string(aptitude)]; found {
			count++
		}
	}
	return count
}

// ApplyBackground changes the character's trait according to the history values
func (character *Character) ApplyBackground(background Background, universe Universe) error {

	cost := 0

	// For each upgrade associated to the history, apply each option.
	for _, upgrade := range background.Upgrades {
		err := character.ApplyUpgrade(Upgrade{
			Mark: MarkSpecial,
			Name: upgrade,
			Cost: &cost,
		}, universe)
		if err != nil {
			return err
		}
	}

	// Add the background to the character's backgrounds
	character.Backgrounds[background.Name] = background

	return nil
}

// ApplyUpgrade changes the character's attributes according to the given upgrade.
func (character *Character) ApplyUpgrade(upgrade Upgrade, universe Universe) error {

	var err error

	// Find the attribute corresponding to the upgrade., and initialize a new
	// rule if there isn't any.
	coster, found := universe.FindCoster(upgrade.Name)
	if !found {
		coster = Rule{
			Name: upgrade.Name,
		}
	}

	// If no cost is defined, compute it on the fly.
	if upgrade.Cost == nil {
		cost, err := coster.Cost(universe, *character)
		if err != nil {
			return err
		}

		upgrade.Cost = &cost
	}

	// Update the spent experience.
	character.Spent += *upgrade.Cost

	// Apply the upgrade depending on the target attribute.
	switch attribute := coster.(type) {
	case Characteristic:
		err = character.ApplyCharacteristicUpgrade(attribute, upgrade)
	case Skill:
		err = character.ApplySkillUpgrade(attribute, upgrade)
	case Talent:
		err = character.ApplyTalentUpgrade(attribute, upgrade)
	case Aptitude:
		err = character.ApplyAptitudeUpgrade(attribute, upgrade)
	case Gauge:
		err = character.ApplyGaugeUpgrade(attribute, upgrade)
	case Rule:
		err = character.ApplyRuleUpgrade(attribute, upgrade)
	}

	return err
}

func (character *Character) ApplyCharacteristicUpgrade(characteristic Characteristic, upgrade Upgrade) error {

	// Get the attribute from the character's characteristic map.
	c, found := character.Characteristics[characteristic.Name]
	if !found {
		c = characteristic
	}

	// Increment the tier if the mark is default.
	if upgrade.Mark == MarkDefault {
		c.Tier++
	}

	// Parse the characteristic's upgrade value.
	raw := strings.TrimSpace(strings.TrimLeft(upgrade.Name, characteristic.Name))
	value, err := strconv.Atoi(raw)
	if err != nil {
		return NewError(InvalidCharacteristicValue)
	}

	// Update the characteristic value.
	if strings.HasPrefix(raw, "+") || strings.HasPrefix(raw, "-") {
		c.Value += value
	} else {
		c.Value = value
	}

	character.Characteristics[c.Name] = c

	return nil
}

func (character *Character) ApplySkillUpgrade(skill Skill, upgrade Upgrade) error {

	// Get the skill from the character's skill map.
	s, found := character.Skills[skill.FullName()]
	if !found {
		s = skill
	}

	// Increment the tier
	s.Tier++

	// Put the skill back on the map.
	character.Skills[skill.FullName()] = s

	return nil
}

func (character *Character) ApplyTalentUpgrade(talent Talent, upgrade Upgrade) error {

	// Get the talent from the character.
	t, found := character.Talents[talent.FullName()]
	if !found {
		t = talent
	}

	// Increment the value of the talent.
	t.Value++

	// Put it back on the map.
	character.Talents[talent.FullName()] = t

	return nil
}

func (character *Character) ApplyAptitudeUpgrade(aptitude Aptitude, upgrade Upgrade) error {

	// Add the aptitude to the character's aptitudes.
	character.Aptitudes[string(aptitude)] = aptitude

	return nil
}

func (character *Character) ApplyGaugeUpgrade(gauge Gauge, upgrade Upgrade) error {

	// Get the gauge from the character.
	g, found := character.Gauges[gauge.Name]
	if !found {
		g = gauge
	}

	// Parse the gauge's upgrade value.
	raw := strings.TrimSpace(strings.TrimLeft(upgrade.Name, g.Name))
	value, err := strconv.Atoi(raw)
	if err != nil {
		return NewError(InvalidGaugeValue)
	}

	// Update the gauge value.
	if strings.HasPrefix(raw, "+") || strings.HasPrefix(raw, "-") {
		g.Value += value
	} else {
		g.Value = value
	}

	// Set the gauge back on the map.
	character.Gauges[g.Name] = g

	return nil
}

func (character *Character) ApplyRuleUpgrade(rule Rule, upgrade Upgrade) error {

	// Add the rule to the character's rules.
	character.Rules[rule.Name] = rule

	return nil
}

// Print the character sheet on the screen
func (character Character) Print() {
	// Print the name
	fmt.Printf("%s\t%s\n", theme.Title("Name"), character.Name)

	// Print the backgrounds
	backgrounds := []Background{}

	for _, background := range character.Backgrounds {
		backgrounds = append(backgrounds, background)
	}

	slice.Sort(backgrounds, func(i, j int) bool {
		if backgrounds[i].Type != backgrounds[j].Type {
			return backgrounds[i].Type < backgrounds[j].Type
		}

		return backgrounds[i].Name < backgrounds[j].Name
	})

	for _, background := range backgrounds {
		fmt.Printf("%s\t%s\n", theme.Title(strings.Title(background.Type)), strings.Title(background.Name))
	}

	// Print the aptitudes
	aptitudes := []Aptitude{}

	for _, aptitude := range character.Aptitudes {
		aptitudes = append(aptitudes, aptitude)
	}

	slice.Sort(aptitudes, func(i, j int) bool {
		return aptitudes[i] < aptitudes[j]
	})

	fmt.Printf("\n%s (%s)\n", theme.Title("Aptitudes"), theme.Value(fmt.Sprintf("%d", len(aptitudes))))
	for _, aptitude := range aptitudes {
		fmt.Printf("%s\n", strings.Title(string(aptitude)))
	}

	// Print the experience
	fmt.Printf("\n%s\t%d/%d\n", theme.Title("Experience"), character.Spent, character.Experience)

	// Print the characteristics
	fmt.Printf("\n%s\n", theme.Title("Characteristics"))

	characteristics := []Characteristic{}
	for _, characteristic := range character.Characteristics {
		characteristics = append(characteristics, characteristic)
	}

	slice.Sort(characteristics, func(i, j int) bool {
		return characteristics[i].Name < characteristics[j].Name
	})

	for _, characteristic := range characteristics {
		fmt.Printf("%s\t%s %s\n", characteristic.Name, theme.Value(characteristic.Value), theme.Value(characteristic.Level()))
	}

	// Print the gauges
	fmt.Printf("\n%s\n", theme.Title("Gauges"))

	gauges := []Gauge{}
	for _, gauge := range character.Gauges {
		gauges = append(gauges, gauge)
	}

	slice.Sort(gauges, func(i, j int) bool {
		return gauges[i].Name < gauges[j].Name
	})

	for _, gauge := range gauges {
		fmt.Printf("%s\t%s\n", gauge.Name, theme.Value(gauge.Value))
	}

	// Print the skills using a tabwriter
	fmt.Printf("\n%s\n", theme.Title("Skills"))

	skills := []Skill{}
	for _, skill := range character.Skills {
		skills = append(skills, skill)
	}

	slice.Sort(skills, func(i, j int) bool {
		return skills[i].FullName() < skills[j].FullName()
	})

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 2, ' ', 0)
	for _, skill := range skills {
		fmt.Fprintf(w, "%s\t+%s\n", strings.Title(skill.FullName()), theme.Value((skill.Tier-1)*10))
	}
	w.Flush()

	// Print the talents
	fmt.Printf("\n%s\n", theme.Title("Talents"))

	talents := []Talent{}
	for _, talent := range character.Talents {
		talents = append(talents, talent)
	}

	slice.Sort(talents, func(i, j int) bool {
		return talents[i].FullName() < talents[j].FullName()
	})

	for _, talent := range talents {
		if talent.Value != 1 {
			fmt.Printf("%s (%s)\n", strings.Title(talent.FullName()), theme.Value(talent.Value))
		} else {
			fmt.Printf("%s\n", strings.Title(talent.FullName()))
		}
	}

	// Print the special rules
	fmt.Printf("\n%s\n", theme.Title("Rules"))

	rules := []Rule{}

	for _, rule := range character.Rules {
		rules = append(rules, rule)
	}

	slice.Sort(rules, func(i, j int) bool {
		return rules[i].Name < rules[j].Name
	})

	for _, rule := range rules {
		fmt.Printf("%s\t%s\n", strings.Title(rule.Name), rule.Description)
	}
}
