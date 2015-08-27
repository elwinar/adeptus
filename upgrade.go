package adeptus

var (
	regex_xp = regexp.MustCompile(`\(?\d+xp\)?`) // Match `150xp` and `(150xp)`
)

type Upgrade interface {
	Mark string
	Name string
	Cost string
	Line int
}

// ParseUpgrade generate an upgrade from a raw line
func ParseUpgrade(raw string, line int) (Upgrade, error) {
	// Initialize a new upgrade
	upgrade := Upgrade{
		Line: line,
	}
	
	// Get the fields of the line
	fields := strings.Fields(raw)
	
	// The minimum number of fields is 2
	if len(fields) < 2 {
		return Upgrade{}, fmt.Errorf("not enought")
	}
	
	// Check that the mark is a valid one
	if !in(fields[0], []string{"*", "+", "-"}) {
		return Upgrade{}, fmt.Errorf("%s isn't a valid mark", fields[0])
	}
	
	// Set the upgrade mark
	upgrade.Mark = fields[0]
	fields = fields[1:]
	
	// Check if a field seems to be a cost field
	for i, field := range fields {
		if !regex_xp.MatchString(field) {
			continue
		}
		
		upgrade.Cost = regex_xp.FindString(field)
		fields = append(fields[:i], fields[i+1:]...)
		break
	}
	
	// The remaining line is the name of the upgrade
	upgrade.Name = strings.Join(fields, " ")
	
	return upgrade, nil
}
