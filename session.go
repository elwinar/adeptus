package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	date     time.Time
	label    string
	reward   int
	upgrades []Upgrade
}

// adds a label and a date to the session
// raw: <date>(<separator><label>)?(<separator><xp>)?
func (s Session) addLabel(raw string) (err error) {
	splits := strings.SplitN(strings.TrimSpace(raw), " ", 2)
	var err error
	var t time.Time

	if splits == nil {
		err = fmt.Errorf("Incorrect format for raw. Expected \" \" in string")
		return
	}

	date := strings.TrimSpace(splits[0])
	for k, f := range formats {
		t, err = time.Parse(f, date)
		if err == nil {
			swap := formats[0]
			formats[0] = f
			formats[k] = swap
			break
		}
	}
	if err != nil {
		return
	}
	s.date = t

	if len(splits) == 2 {
		s.label = strings.TrimSpace(splits[1])
	}
	return
}

// adds a reward to the session
// raw: <mark> ?<value> ?xp
func (s Session) addReward(raw string) error {
	value := strings.TrimSpace(raw[1 : len(raw)-3])
	tmp, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	s.reward = int(tmp)
	return nil
}

// adds an upgrade to the session
func (s Session) addUpgrade(raw string) error {

	u, err := NewUpgrade(raw)
	if err != nil {
		return err
	}
	s.upgrades = append(s.upgrades, u)
	return nil
}
