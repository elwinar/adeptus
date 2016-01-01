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
	Backgrounds     map[string]Background
	Aptitudes       map[string]Aptitude
	Characteristics map[string]Characteristic
	Skills          map[string]Skill
	Talents         map[string]Talent
	Gauges          map[string]Gauge
	Rules           map[string]Rule
	Spells          map[string]Spell
	Experience      int
	Spent           int
	History         []Upgrade
}

// NewCharacter creates a new character from the given sheet and universe.
func NewCharacter(universe Universe, sheet Sheet) (*Character, error) {

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
		Spells:          make(map[string]Spell),
		Experience:      0,
		Spent:           0,
	}

	// The characteristics described in the header of the sheet are parsed as upgrades
	for _, upgrade := range sheet.Characteristics {

		// Get the characteristic from the universe
		characteristic, found := universe.FindCharacteristic(upgrade)
		if !found {
			return nil, NewError(UndefinedCharacteristic, upgrade.Line)
		}

		// Check it is not already applied
		_, found = c.Characteristics[characteristic.Name]
		if found {
			return nil, NewError(DuplicateUpgrade, upgrade.Line)
		}

		// Apply the upgrade
		err := characteristic.Apply(&c, upgrade)
		if err != nil {
			return nil, err
		}
	}

	// Next are the backgrounds
	for typ, metas := range sheet.Header.Metas {

		for _, meta := range metas {

			// Find the background corresponding to the meta
			background, found := universe.FindBackground(typ, meta.Label)
			if !found {
				return nil, NewError(UndefinedBackground, meta.Line, typ, meta.Label)
			}

			err := background.Apply(&c, universe)
			if err != nil {
				return nil, err
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
				return nil, err
			}
		}
	}

	return &c, nil
}

// Intersect return the number of aptitudes of the given slice
// that are in the character's aptitudes.
func (c *Character) Intersect(aptitudes []Aptitude) int {

	count := 0
	for _, aptitude := range aptitudes {
		if _, found := c.Aptitudes[string(aptitude)]; found {
			count++
		}
	}
	return count
}

// ApplyUpgrade changes the character's attributes according to the given upgrade.
func (c *Character) ApplyUpgrade(upgrade Upgrade, universe Universe) error {

	// Find the attribute corresponding to the upgrade., and initialize a new
	// rule if there isn't any.
	coster, found := universe.FindCoster(upgrade)
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

	// Apply the upgrade.
	err := coster.Apply(c, upgrade)
	c.Spent += *upgrade.Cost

	// If there is no error, add the upgrade to the history.
	if err == nil {
		c.History = append(c.History, upgrade)
	}

	return err
}

// Print the character sheet on the screen
func (c *Character) Print() {
	// Print the name
	fmt.Printf("%s\t%s\n", theme.Title("Name"), c.Name)

	// Print the backgrounds
	backgrounds := []Background{}

	for _, background := range c.Backgrounds {
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

	for _, aptitude := range c.Aptitudes {
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
	fmt.Printf("\n%s\t%d/%d\n", theme.Title("Experience"), c.Spent, c.Experience)

	// Print the characteristics
	fmt.Printf("\n%s\n", theme.Title("Characteristics"))

	characteristics := []Characteristic{}
	for _, characteristic := range c.Characteristics {
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
	for _, gauge := range c.Gauges {
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
	for _, skill := range c.Skills {
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
	for _, talent := range c.Talents {
		talents = append(talents, talent)
	}

	slice.Sort(talents, func(i, j int) bool {
		return talents[i].FullName() < talents[j].FullName()
	})

	w = tabwriter.NewWriter(os.Stdout, 10, 1, 2, ' ', 0)
	for _, talent := range talents {
		if talent.Value != 1 {
			fmt.Fprintf(w, "%s (%d)\t%s\n", strings.Title(talent.FullName()), talent.Value, talent.Description)
		} else {
			fmt.Fprintf(w, "%s\t%s\n", strings.Title(talent.FullName()), talent.Description)
		}
	}
	w.Flush()

	// Print the spells

	if len(c.Spells) != 0 {

		fmt.Printf("\n%s\n", theme.Title("Spells"))

		spells := []Spell{}

		for _, spell := range c.Spells {
			spells = append(spells, spell)
		}

		slice.Sort(spells, func(i, j int) bool {
			return spells[i].Name < spells[j].Name
		})

		for _, spell := range spells {
			fmt.Fprintf(w, "%s\t%s\n", strings.Title(spell.Name), spell.Description)
		}
		w.Flush()
	}

	// Print the special rules

	if len(c.Rules) != 0 {
		fmt.Printf("\n%s\n", theme.Title("Rules"))

		rules := []Rule{}

		for _, rule := range c.Rules {
			rules = append(rules, rule)
		}

		slice.Sort(rules, func(i, j int) bool {
			return rules[i].Name < rules[j].Name
		})

		for _, rule := range rules {
			fmt.Printf("%s\t%s\n", strings.Title(rule.Name), rule.Description)
		}
		w.Flush()
	}
}

// PrintHistory displays the history of expences of the character.
func (c *Character) PrintHistory() {
	// Print the name.
	fmt.Printf("%s\t%s\n", theme.Title("Name"), c.Name)

	// Print the experience
	fmt.Printf("\n%s\t%d/%d\n", theme.Title("Experience"), c.Spent, c.Experience)

	// Print the history.
	fmt.Printf("\n%s\n", theme.Title("History"))

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 2, ' ', 0)
	for _, upgrade := range c.History {
		if upgrade.Cost != nil {
			fmt.Fprintf(w, "%d\t%s\n", *upgrade.Cost, strings.Title(upgrade.Name))
		} else {
			fmt.Fprintf(w, "%d\t%s\n", 0, strings.Title(upgrade.Name))
		}
	}
	w.Flush()
}

// Suggest the next purchasable upgrades of the character.
func (c *Character) Suggest(max int, all bool) {

	// Aggregate each coster into a unique slice of costers.
	costers := []Coster{}
	for _, upgrade := range universe.Characteristics {
		costers = append(costers, upgrade)
	}

	for _, upgrade := range universe.Skills {
		costers = append(costers, upgrade)
	}

	for _, upgrade := range universe.Talents {
		costers = append(costers, upgrade)
	}

	for _, upgrade := range universe.Spells {
		costers = append(costers, upgrade)
	}

	for _, upgrade := range universe.Gauges {
		upgrade.Value = 1
		costers = append(costers, upgrade)
	}

	// Default max value equals to the remaining XP.
	if max == 0 {
		max = c.Experience - c.Spent
	}

	// The slice of appliable upgrades.
	var appliable []Upgrade

	// Attempt to apply each coster once.
	for _, coster := range costers {
		var upgrade Upgrade

		// Don't propose the upgrade its cost cannot be defined
		cost, err := coster.Cost(universe, *c)
		if err != nil {
			continue
		}

		// Don't propose the upgrade if it is free.
		if cost == 0 {
			continue
		}

		// Don't propose the upgrade if it is too expensive.
		if !all && max < cost {
			continue
		}

		upgrade.Cost = &cost
		upgrade.Mark = MarkApply

		switch t := coster.(type) {

		case Characteristic:
			upgrade.Name = fmt.Sprintf("%s +%d", t.Name, 5)
			err = t.Apply(c, upgrade)

		case Skill:
			upgrade.Name = fmt.Sprintf("%s", t.Name)
			err = t.Apply(c, upgrade)

		case Talent:
			upgrade.Name = fmt.Sprintf("%s", t.Name)
			err = t.Apply(c, upgrade)

		case Spell:
			upgrade.Name = fmt.Sprintf("%s", t.Name)
			err = t.Apply(c, upgrade)

		case Gauge:
			upgrade.Name = fmt.Sprintf("%s +%d", t.Name, 1)
			err = t.Apply(c, upgrade)
		}

		if err != nil {
			continue
		}

		appliable = append(appliable, upgrade)
	}

	// Sort by cost then name.
	slice.Sort(appliable, func(i, j int) bool {
		ci, cj := *appliable[i].Cost, *appliable[j].Cost
		if ci == cj {
			return appliable[i].Name < appliable[j].Name
		}
		return ci < cj
	})

	// Print the name.
	fmt.Printf("%s\t%s\n", theme.Title("Name"), c.Name)

	// Print the experience
	fmt.Printf("\n%s\t%d/%d\n", theme.Title("Experience"), c.Spent, c.Experience)

	// Print the history.
	fmt.Printf("\n%s\n", theme.Title("Suggestions"))

	w := tabwriter.NewWriter(os.Stdout, 10, 1, 2, ' ', 0)
	for i, upgrade := range appliable {
		if i > 0 && *appliable[i-1].Cost != *upgrade.Cost {
			fmt.Fprintln(w)
		}
		fmt.Fprintf(w, "%s\t%s\n", theme.Value(*upgrade.Cost), strings.Title(upgrade.Name))
	}
	w.Flush()
}
