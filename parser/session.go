package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Session blocks describe a game session, with its reward and upgrades to the
// character.
type Session struct {
	Date     time.Time
	Title    string
	Reward   *int
	Upgrades []Upgrade
}

// formats list the recognized date formats for session headlines
var formats = []string{
	"2006/01/02",
	"2006-01-02",
	"2006.01.02",
}

// parseSession parse a block of line into a Session, and return an error in the
// event of an invalid line.
func parseSession(block []line) (Session, error) {
	// Check the block is non-empty
	if len(block) == 0 {
		return Session{}, fmt.Errorf("unexpected block size")
	}

	// The first line of the block is always the headline
	headline := block[0]

	// Get the fields of the line
	fields := strings.Fields(headline.Text)

	// The line should at least have a date field
	if len(fields) == 0 {
		return Session{}, fmt.Errorf("line %d: expected at least a date", headline.Number)
	}

	// Check if the first field is a recognized date
	var err error
	var date time.Time
	for i, format := range formats {
		// Try the format
		date, err = time.Parse(format, fields[0])
		if err != nil {
			continue
		}

		// Put the format in the first
		formats[0], formats[i] = formats[i], formats[0]

		// The format is good, stop trying
		break
	}

	// If we have an error, that's because no format matched
	if err != nil {
		return Session{}, fmt.Errorf("line %d: invalid date format", headline.Number)
	}

	// Check if a field seems to be a reward field
	var reward *int
	for i, field := range fields[1:] {
		// If one end has the brackets but not the other, that's an error:
		// brackets does by pairs, and are forbidden in the title
		if strings.HasPrefix(field, "[") != strings.HasSuffix(field, "]") {
			return Session{}, fmt.Errorf("line %d: brackets [] must open-close and contain no blank", headline.Number)
		}

		// If the brackets are absents, that's not a reward, so skip the field.
		// Note that as both ends have brackets (or not), we just need to test
		// one of them.
		if !strings.HasPrefix(field, "[") {
			break
		}
		
		// There can be only one reward on the line
		if reward != nil {
			return Session{}, fmt.Errorf("line %d: the can be only one reward on the headline", headline.Number)
		}

		// Check position of the reward
		if i != 0 && i != len(fields)-1 {
			return Session{}, fmt.Errorf("line %d: reward must be in second or last position of the line", headline.Number)
		}

		// Trim the field to get the raw reward
		raw := strings.Trim(field, "[]")

		// Parse the reward
		r, err := strconv.Atoi(raw)
		if err != nil {
			return Session{}, fmt.Errorf("line %d: expected integer, got \"%s\"", headline.Number, raw)
		}
		reward = &r

		// Remove the field from the slice
		fields = append(fields[:i], fields[i+1:]...)
	}

	// The remaining fields are the title
	title := strings.Join(fields, " ")

	// Parse the other lines as upgrades
	upgrades := []Upgrade{}
	for _, line := range block[1:] {
		upgrade, err := parseUpgrade(line)
		if err != nil {
			return Session{}, err
		}

		upgrades = append(upgrades, upgrade)
	}

	// Return the session
	return Session{
		Date:     date,
		Reward:   reward,
		Title:    title,
		Upgrades: upgrades,
	}, nil
}
