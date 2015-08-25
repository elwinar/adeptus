package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	REWARD_REGEX = `^( )*\+( )*(\d)*( )*xp$`
)

func Parse(filename string) (Character, error) {
	c := Character{}
	f, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("Unable to open character: \"%s.\"", filename)
		return c, err
	}
	scanner := bufio.NewScanner(f)

	h, err := scanHeader(scanner)
	if err != nil {
		err = fmt.Errorf("Incorrect formating in header: \"%s.\"", filename)
		return c, err
	}
	c.AddHeader(h)

	it := 0
	for scanner.Scan() {
		it++
		s, err := scanSession(scanner)
		if err != nil {
			err = fmt.Errorf("Incorrect formating in session %d: \"%s.\"", it, filename)
			return c, err
		}
		c.AddSession(s)
	}

	err = scanner.Err()
	if err != nil {
		err = fmt.Errorf("Error during scan: \"%s.\"", filename)
		return c, err
	}

	return c, nil
}

// Scans the pairs key:value in the lines and returns the header
func scanHeader(scanner *bufio.Scanner) (Header, error) {
	h := Header{}

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "\n" {
			return h, nil
		}
		err := h.addMetadata(text)
		if err != nil {
			return h, err
		}
	}

	return h, nil
}

// Reads line until the session_end token is reached and returns the session
func scanSession(scanner *bufio.Scanner) (Session, error) {
	s := Session{}

	// scan label
	text := strings.TrimSpace(scanner.Text())
	err := s.addLabel(text)
	if err != nil {
		return s, err
	}

	// scan potential reward
	text = strings.TrimSpace(scanner.Text())
	// no need to test error since REWARD_REGEX is correct
	match, _ := regexp.MatchString(REWARD_REGEX, text)
	if match {
		s.addReward(text)
	}

	// scan upgrades
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "\n" {
			return s, nil
		}
		err = s.addUpgrade(text)
		if err != nil {
			return s, err
		}
	}
	return s, nil
}
