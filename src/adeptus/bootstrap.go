package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/urfave/cli.v1"
)

// Bootstrap open and parse universe and character sheet.
func Bootstrap(ctx *cli.Context) (Universe, *Character, error) {
	// Open and parse the universe
	files, err := filepath.Glob(ctx.GlobalString("universe") + "/*.yaml")
	if err != nil {
		return Universe{}, nil, fmt.Errorf("%s %s", theme.Error("unable to open universe:"), err)
	}
	
	var universe Universe
	for _, f := range files {
		u, err := os.Open(f)
		if err != nil {
			return Universe{}, nil, fmt.Errorf("%s %s", theme.Error("unable to open universe:"), err)
		}
		defer func() {
			_ = u.Close()
		}()
		tmp, err := ParseUniverse(u)
		if err != nil {
			return Universe{}, nil, fmt.Errorf("%s %s", theme.Error("corrupted universe:"), err)
		}
		
		universe, err = MergeUniverses(universe, tmp)
		if err != nil {
			return Universe{}, nil, fmt.Errorf("%s %s", theme.Error("corrupted universe:"), err)
		}
	}

	// Open and parse character sheet.
	args := ctx.Args()
	if len(args) == 0 {
		return Universe{}, nil, fmt.Errorf("%s undefined character", theme.Error("unable to open character sheet:"))
	}
	name := args[len(args)-1]
	c, err := os.Open(name)
	if err != nil {
		return Universe{}, nil, fmt.Errorf("%s %s", theme.Error("unable to open character sheet:"), err)
	}
	defer func() {
		_ = c.Close()
	}()
	sheet, err := ParseSheet(c)
	if err != nil {
		return Universe{}, nil, fmt.Errorf("%s %s", theme.Error("corrupted character sheet:"), err)
	}

	// Create character with the sheet
	character, err := NewCharacter(universe, sheet)
	if err != nil {
		return Universe{}, nil, fmt.Errorf("%s %s", theme.Error("unable to create character:"), err)
	}

	return universe, character, nil
}

// MergeUniverses two universes into one.
func MergeUniverses(u1, u2 Universe) (Universe, error) {
	
	duplicates := make(map[string]struct{})
	
	// Merge backgounds.
	for label := range u2.Backgrounds {
		if u1.Backgrounds == nil {
			u1.Backgrounds = make(map[string][]Background)
		}
		u1.Backgrounds[label] = append(u1.Backgrounds[label], u2.Backgrounds[label]...)
		for _, b := range u1.Backgrounds[label] {
			if _, ok := duplicates[b.Name]; ok {
				return Universe{}, fmt.Errorf("background %s - %s already defined", label, b.Name)
			}
			duplicates[b.Name] = struct{}{}
		}
	}
	
	// Merge aptitudes.
	u1.Aptitudes = append(u1.Aptitudes, u2.Aptitudes...)
	duplicates = make(map[string]struct{})
	for _, a := range u1.Aptitudes {
		if _, ok := duplicates[string(a)]; ok {
			return Universe{}, fmt.Errorf("aptitude %s already defined", a)
		}
		duplicates[string(a)] = struct{}{}
	}
	
	// Merge characteristics.
	u1.Characteristics = append(u1.Characteristics, u2.Characteristics...)
	duplicates = make(map[string]struct{})
	for _, a := range u1.Characteristics {
		if _, ok := duplicates[a.Name]; ok {
			return Universe{}, fmt.Errorf("characteristic %s already defined", a.Name)
		}
		duplicates[a.Name] = struct{}{}
	}
	
	// Merge gauges.
	u1.Gauges = append(u1.Gauges, u2.Gauges...)
	duplicates = make(map[string]struct{})
	for _, a := range u1.Gauges {
		if _, ok := duplicates[a.Name]; ok {
			return Universe{}, fmt.Errorf("gauge %s already defined", a.Name)
		}
		duplicates[a.Name] = struct{}{}
	}
	
	// Merge Skills.
	u1.Skills = append(u1.Skills, u2.Skills...)
	duplicates = make(map[string]struct{})
	for _, a := range u1.Skills {
		if _, ok := duplicates[a.Name]; ok {
			return Universe{}, fmt.Errorf("skill %s already defined", a.Name)
		}
		duplicates[a.Name] = struct{}{}
	}
	
	// Merge Talents.
	u1.Talents = append(u1.Talents, u2.Talents...)
	duplicates = make(map[string]struct{})
	for _, a := range u1.Talents {
		if _, ok := duplicates[a.Name]; ok {
			return Universe{}, fmt.Errorf("talent %s already defined", a.Name)
		}
		duplicates[a.Name] = struct{}{}
	}
	
	// Merge Spells.
	u1.Spells = append(u1.Spells, u2.Spells...)
	duplicates = make(map[string]struct{})
	for _, a := range u1.Spells {
		if _, ok := duplicates[a.Name]; ok {
			return Universe{}, fmt.Errorf("talent %s already defined", a.Name)
		}
		duplicates[a.Name] = struct{}{}
	}
	
	// Merge Costs.
	if u1.Costs != nil && u2.Costs != nil {
			return Universe{}, fmt.Errorf("costs already defined")
	}
	if u1.Costs == nil {
		u1.Costs = u2.Costs
	}
	
	return u1, nil
}
