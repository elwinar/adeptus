type Session struct {
	date 		dateTime
	label		string
	reward		int
	upgrades	[]Upgrade
}

var formats []string

func init() {
	formats.push("2006/02/03")
	formats.push("2006-02-03")
	formats.push("2006_02_03")
	formats.push("2006.02.03")
	formats.push("20060203")
}

// adds a label and a date to the session
// raw: <date> <label>
(s Session) func addLabel(raw string) error {
	splits := strings.Split(raw, " ")
	
	// don't know if this may happend
	date, f := splits[0]
	if !f {
		err := fmt.Errorf("Incorrect format for raw. Expected \" \" in string")
		return err
	}
	
	for k, f := range formats {
		d, err := time.Parse(f, date)
		if err == nil {
			swap := formats[0]
			formats[0] = f
			formats[k] = swap
			break
		}
	}
	if err != nil {
		return err
	}
	s.date = d
	
	_, f = splits[1]
	if f {
		s.label = strings.join(splits[1:])
	}
	return nil
}

// adds a reward to the session
// raw: <mark> ?<value> ?xp
(s Session) func addReward(raw string) {
	value := strings.TrimSpace(raw[1:-2])
	s.reward = int(strconv.ParseInt(value, 10, 64))
}

// adds an upgrade to the session
(s Session) func addUpgrade(raw string) error {
	
	u, err := NewUpgrade(raw)
	if err != nil {
		return err
	}
	s.upgrades.push(u)
	return nil
}