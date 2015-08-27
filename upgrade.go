package adeptus

var (
	regex_xp = regexp.MustCompile(`\(?\d+xp\)?`) // Match `150xp` and `(150xp)`
)

type Upgrade interface {
}

type RawUpgrade struct {
	mark string
	name string
	cost string
	line	int
}

// ParseUpgrade generate an upgrade from a raw line
func ParseUpgrade(line int, raw string) (RawUpgrade, error) {
	upgrade := RawUpgrade{
		line: line,
	}
	
	// Get the fields of the line
	fields := strings.Fields(raw)
	
	// The minimum number of fields is 2
	if len(fields) < 2 {
		return upgrade, fmt.Errorf("Error on line %d: expected at least mark and label.", line)
	}
	
	// Check that the mark is a valid one
	if !in(fields[0], []string{"*", "+", "-"}) {
		return upgrade, fmt.Errorf("Error on line %d: %s is not a valid mark", line, fields[0])
	}
	
	// Set the upgrade mark
	upgrade.mark = fields[0]
	fields = fields[1:]
	
	// Check if a field seems to be a cost field
	for i, field := range fields {
		if !regex_xp.MatchString(field) {
			continue
		}
		
		upgrade.cost = regex_xp.FindString(field)
		fields = append(fields[:i], fields[i+1:]...)
		break
	}
	
	// The remaining line is the name of the upgrade
	upgrade.name = strings.Join(fields, " ")
	
	return upgrade, nil
}
