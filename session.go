package adeptus

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Session struct {
	Date     time.Time
	Title    string
	Reward   int
	Upgrades []Upgrade
}

type upgradeParser func(Line) (Upgrade, error)

var formats []string = []string{
	"2006/01/02",
	"2006-01-02",
	"2006.01.02",
	"20060102",
}

// ParseSession generate a Session from a block of lines
func ParseSession(block []Line) (Session, error) {
	return parseSession(block, ParseUpgrade)
}

// non-exported function for parseSession:
// dependency injection (parser)
func parseSession(block []Line, parse upgradeParser) (Session, error) {

	session := Session{}
	if len(block) < 1 {
		return session, fmt.Errorf("Unexpected block size.")
	}
	line := block[0]

	// Get the fields of the line
	fields := strings.Fields(line.Text)

	// The minimum number of fields is 1
	if len(fields) < 1 {
		return session, fmt.Errorf("Error on line %d: expected at least a date.", line.Number)
	}

	// Search the right date format
	var err error
	var date time.Time
	for i, f := range formats {
		date, err = time.Parse(f, fields[0])
		if err != nil {
			continue
		}
		// pushes the format found in first position
		swap := formats[0]
		formats[0] = f
		formats[i] = swap
		break
	}
	if err != nil {
		return session, fmt.Errorf("Error on line %d: invalid date format. Expecting \"YYYY/MM/DD\", \"YYYY.MM.DD\" or \"YYY-MM-DD.\"", line.Number)
	}
	
	fields = fields[1:]

	// Check if a field seems to be a reward field
	var reward int
	for i, field := range fields {

		if strings.HasPrefix(field, "[") || strings.HasSuffix(field, "]") {

			// Check that the field has both brackets. If only one bracket is present, there is an error
			if strings.HasPrefix(field, "[") != strings.HasSuffix(field, "]") {
				return session, fmt.Errorf("Error on line %d: brackets [] must open-close and contain no blank.", line.Number)
			}

			// Check position of xp
			if i != 0 && i != len(fields)-1 {
				return session, fmt.Errorf("Error on line %d: experience must be in second or last position of the line.", line.Number)
			}

			// Check value of xp
			xp := strings.TrimSuffix(strings.TrimPrefix(field, "["), "]")
			reward, err = strconv.Atoi(xp)
			if err != nil || len(xp) == 0 {
				return session, fmt.Errorf("Error on line %d: expected number, \"%s\" is no numeric value.", line.Number, xp)
			}

			// remove xp from field slice
			fields = append(fields[:i], fields[i+1:]...)
			break
		}
	}

	// Set the session attributes at the end to return empty upgrade in case of error
	session.Date = date
	session.Reward = reward
	session.Title = strings.Join(fields, " ")

	// Parse upgrades
	block = block[1:]
	for _, line := range block {
		u, err := parse(line)
		if err != nil {
			return Session{}, err
		}
		session.Upgrades = append(session.Upgrades, u)
	}

	return session, nil
}
