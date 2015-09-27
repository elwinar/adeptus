package main

import (
	"fmt"
	"strconv"
	"text/tabwriter"
	"os"

	"github.com/elwinar/adeptus/parser"
	"github.com/elwinar/adeptus/universe"
)

// Character is the type representing a role playing character
type Character struct {
	Name            string
	Histories		map[string]universe.History
	Tarot           universe.Tarot
	Aptitudes       []universe.Aptitude
	Characteristics map[*universe.Characteristic]int
	Skills			map[*universe.Skill]int
	Talents			map[*universe.Talent]int
	Gauges			map[*universe.Gauge]int
	Rules			[]universe.Rule
	Experience		int
	Spent			int
}

// NewCharacter creates a new character given a sheet
func NewCharacter(u universe.Universe, s parser.Sheet) (*Character, error) {
	
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
	
	// Apply each Meta.
	c.Histories = make(map[string]universe.History)
	metasLoop:
	for typ, meta := range h.Metas {
		
		// Treat specific case of tarot
		if typ == "tarot" {

			dice, err := strconv.Atoi(meta.Label)
			if err != nil {
				return nil, fmt.Errorf("expecting numeric tarot")
			} 
			tarot, found := u.FindTarot(dice)
			if !found {
				return nil, fmt.Errorf("tarot %s not found", dice)
			}
			err = c.ApplyHistory(tarot.History, u)
			if err != nil {
				return nil, err
			}
			c.Tarot = tarot
			continue metasLoop
		}
		
		histories, found := u.Histories[typ]
		
		// Check the history type exists in universe.
		if !found {
			return nil, fmt.Errorf("undefined history %s in universe", typ)
		}
		
		// Search the hystory corresponding to the provided meta
		for _, h := range histories {
			if meta.Label != h.Name {
				continue
			}
			
			// Apply the history
			c.Histories[typ] = h
			err := c.ApplyHistory(h, u)
			if err != nil {
				return nil, err
			}
			
			continue metasLoop
		}
		return nil, fmt.Errorf("history %s not defined for history type %s in universe", meta.Label, typ)
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
func (c *Character) ApplyHistory(h universe.History, u universe.Universe) error {
	
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
func (c *Character) ApplyUpgrade(up parser.Upgrade, un universe.Universe) error {
	
	
	// Defer payment of upgrade
	var payed bool
	var coster universe.Coster
			
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
			return fmt.Errorf(`universe provides upgrade "%s" but does not define the charactersitic "%s"`, up.Name, name)
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
	rule := universe.Rule{
		Name: name,
		Description: speciality,
	}
	c.Rules = append(c.Rules, rule)
	
	return nil
}

// Debug prints the current values of the character
func (c Character) Debug() {
	
	w := new(tabwriter.Writer)

	// Format in tab-separated columns.
	w.Init(os.Stdout, 4, 8, 0, '\t', 0)
	
	fmt.Fprintf(w, "Name\t%s\n", c.Name)
	for label, history := range c.Histories {
		fmt.Fprintf(w, "%s\t%s\n", label, history.Name)
	}
	fmt.Fprintf(w, "Tarot\t%s\n", c.Tarot.Name)
	fmt.Fprintf(w, "\nExperience:\t%d/%d\n", c.Spent, c.Experience)
	fmt.Fprintf(w, "\nCharacteristics\n")
	for c, value := range c.Characteristics {
		fmt.Fprintf(w, "%s\t%d\n", c.Name, value)
	}
	fmt.Fprintf(w, "\nAptitudes\n")
	for _, a := range c.Aptitudes {
		fmt.Fprintf(w, "\t%s\n", a)
	}
	fmt.Fprintf(w, "\nTalents\n")
	for t, v := range c.Talents {
		fmt.Fprintf(w, "%s\t%d\n", t.Name, v)
	}
	fmt.Fprintf(w, "\nSkills\n")
	for s, v := range c.Skills {
		fmt.Fprintf(w, "%s\t%d\n", s.Name, v)
	}
	fmt.Fprintf(w, "\nRules\n")
	for _, r := range c.Rules {
		fmt.Fprintf(w, "%s\t%s\n", r.Name, r.Description)
	}
	w.Flush()
}