package adeptus

import "time"

type Session struct {
	Date     time.Time
	Title    string
	Reward   int
	Upgrades []Upgrade
}

formats := []string{
	"2006/02/01",
	"2006-02-01",
	"2006.02.01",
}

func ParseSession(line Line) (Session, error) {

	session := Session{}

	// Get the fields of the line
	fields := strings.Fields(line.Text)

	// The minimum number of fields is 1
	if len(fields) < 1 {
		return session, fmt.Errorf("Error on line %d: expected at least a date.", line.Number)
	}

	// Retrieve the date
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
	}
	if err != nil {
		return session, fmt.Errorf("Error on line %d: invalid date format. Expecting \"YYYY/MM/DD\", \"YYYY.MM.DD\" or \"YYY-MM-DD."\", line.Number)
	}
	fields = fields[1:]

	// Check if a field seems to be a reward field
	var reward string
	for i, field := range fields {
		
		if strings.HasPrefix(field, "[") || strings.HasSuffix(field, "]") {
		
			// Check that the field has both brackets. If only one bracket is present, there is an error
			if strings.HasPrefix(field, "[") != strings.HasSuffix(field, "]") {
				return session, fmt.Errorf("Error on line %d: brackets [] must open-close and contain no blank.", line.Number)
			}
		
			// Check position of xp
			if i == 0 || i == len(fields) - 1 {
				return session, fmt.Errorf("Error on line %d: experience must be after mark or at the end of line.", line.Number)
			}
			
			// Check value of xp
			reward = strings.TrimSuffix(strings.TrimPrefix(field, "["), "]")
			_, err := strconv.Atoi(reward)
			if err || len(reward) == 0 {
				return session, fmt.Errorf("Error on line %d: expected number, \"%s\" is no numeric value.", line.Number, reward)
			}
			
			// remove xp from field slice
			fields = append(fields[:i], fields[i+1:]...)
			break
		}
	}

	// Set the session attributes at the end to return empty upgrade in case of error
	session.date = date
	session.reward = reward
	session.title = strings.Join(fields, " ")
	
	return session, nil
}
