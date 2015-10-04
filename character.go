package main

import (
	"fmt"

	"github.com/bradfitz/slice"
)

// Character is the type representing a role playing character
type Character struct {
	Name            string
	Histories       map[string][]History
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

	// Retrieve the name from the sheet header.
	if len(h.Name) == 0 {
		return nil, fmt.Errorf("empty name")
	}
	c.Name = h.Name

	// Apply the initial characteristics from the sheet.
	c.Characteristics = make(map[*Characteristic]int)
	for _, characteristic := range s.Characteristics {

		// Identify name and value.
		name, value, _, err := IdentifyCharacteristic(characteristic.Name)
		if err != nil {
			return nil, err
		}

		// Retrieve characteristic from given it's name.
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

	// Check all characteristics from are defined for the character
checkCharacteristics:
	for _, u := range u.Characteristics {
		for c := range c.Characteristics {
			if c.Name == u.Name {
				continue checkCharacteristics
			}
		}
		return nil, fmt.Errorf("charactersitic %s of not defined for character", u.Name)
	}

	// Make the character's gauges, skills and talents maps
	c.Skills = make(map[*Skill]int)
	c.Talents = make(map[*Talent]int)
	c.Gauges = make(map[*Gauge]int)

	// Apply each Meta.
	c.Histories = make(map[string][]History)
	for typ, meta := range h.Metas {

		histories, found := u.Histories[typ]

		// Check the history type exists in
		if !found {
			return nil, fmt.Errorf("undefined history %s in universe", typ)
		}

		c.Histories[typ] = []History{}

	metasLoop:
		for _, m := range meta {

			// Search the history corresponding to the provided meta
			for _, h := range histories {
				if m.Label != h.Name {
					continue
				}

				// Apply the history
				c.Histories[typ] = append(c.Histories[typ], h)
				err := c.ApplyHistory(h, u)
				if err != nil {
					return nil, err
				}

				continue metasLoop
			}
			return nil, fmt.Errorf("history %s not defined for history type %s in universe", m.Label, typ)
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

// ApplyHistory changes the character's trait according to the history values
func (c *Character) ApplyHistory(h History, u Universe) error {

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
			return fmt.Errorf(`unexpected cost on upgrade line %d: mark "-" expects no cost value`, up.Line)
		}
		payed = true

	// Cost is hard defined.
	case up.Cost != nil:
		c.Spent += *up.Cost
		payed = true
	}

	// Identify characteristic.
	name, value, sign, err := IdentifyCharacteristic(up.Name)
	if err == nil {

		// Check the characteristic exists.
		characteristic, found := un.FindCharacteristic(name)
		if !found {
			return fmt.Errorf(`provides upgrade "%s" but does not define the charactersitic "%s"`, up.Name, name)
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

	// Sort aptitudes by name and remove duplicates
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
	fmt.Printf("%s\t%s\n", theme.Title("Name"), c.Name)

	for label, histories := range c.Histories {
		for _, history := range histories {
			fmt.Printf("%s\t%s\n", theme.Title(label), history.Name)
		}
	}

	fmt.Printf("\n%s\t%d/%d\n", theme.Title("Experience"), c.Spent, c.Experience)

	fmt.Printf("\n%s\n", theme.Title("Characteristics"))
	for c, value := range c.Characteristics {
		fmt.Printf("%s\t%s\n", c.Name, theme.Value(value))
	}

	fmt.Printf("\n%s\n", theme.Title("Aptitudes"))
	for _, a := range c.Aptitudes {
		fmt.Printf("- %s\n", a)
	}

	fmt.Printf("\n%s\n", theme.Title("Talents"))
	for t, v := range c.Talents {
		if v != 1 {
			fmt.Printf("- %s %s\n", t.Name, theme.Value(v))
		} else {
			fmt.Printf("- %s\n", t.Name)
		}
	}

	fmt.Printf("\n%s\n", theme.Title("Skills"))
	for s, v := range c.Skills {
		fmt.Printf("- %s %s\n", s.Name, theme.Value(v))
	}

	fmt.Printf("\n%s\n", theme.Title("Rules"))
	for _, r := range c.Rules {
		fmt.Printf("- %s %s\n", r.Name, theme.Value(r.Description))
	}
}
