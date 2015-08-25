package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type Upgrade interface {
}

const (
	IS_A_VALUE = `[+-]?(\d)*$`
)

type upgradeFactory func(string) (Upgrade, error)

var upgradeFactories map[string]upgradeFactory

func init() {
	upgradeFactories["characteristic"] = NewCharacteristic
	// 	upgradeFactories["skill"] = NewSkill
	// 	upgradeFactories["talent"] = NewTalent
}

// returns the map and the remain of the given string
func chunkMark(raw string) (mark string, out string, err error) {
	out = raw
	out = strings.TrimSpace(out)
	split := strings.SplitN(out, " ", 2)
	if len(split) <= 1 {
		err = fmt.Errorf(`Incorrect format for raw. Expected " " in string`)
		return
	}
	mark = split[0]
	out = split[len(split)-1]
	return
}

// returns the xp and the remain of the given string
func chunkXp(raw string) (xp string, out string, err error) {
	out = raw
	split := strings.Split("(", out)
	if len(split) == 1 {
		return
	}
	if len(split) != 2 {
		err = fmt.Errorf(`Incorrect format for raw. Expected at most one "("`)
		return
	}
	out = split[0]
	xp = strings.TrimSpace(split[1])
	xp = xp[:len(xp)-4]
	return
}

// returns the value and the remain of the given string
func chunkValue(raw string) (value string, out string, err error) {
	out = raw
	out = strings.TrimSpace(out)
	split := strings.Split(" ", out)
	value = split[len(split)-1]

	// no need to test error, IS_A_VALUE is a correct regex
	match, _ := regexp.MatchString(IS_A_VALUE, value)
	if !match {
		return
	}
	// 	if value[0] == "+" || value[0] == "-" {
	// 		value = value[1:]
	// 	}
	return
}

// Transform the given line into an upgrade
func NewUpgrade(line string) (u Upgrade, err error) {
	
	raw := strings.TrimSpace(line)
	mark, raw, err := chunkMark(raw)
	if err != nil {
		err = fmt.Errorf("Incorrect upgrade in session.")
		return
	}

	xp, raw, err := chunkXp(raw)
	if err != nil {
		err = fmt.Errorf("Incorrect upgrade in session.")
		return
	}

	value, raw, err := chunkValue(raw)
	if err != nil {
		err = fmt.Errorf("Incorrect upgrade in session.")
		return
	}

	name := strings.TrimSpace(raw)
	a, err := GetAttributeByName(name)
	if err != nil {
		u, err = upgradeFactories["Special"](line)
	} else {
		u, err = upgradeFactories[reflect.TypeOf(a).String()](mark + name + value + xp)
	}

	return
}
