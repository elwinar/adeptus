package universe

type Universe struct {
	Origins         []Origin
	Backgrounds     []Background
	Roles           []Role
	Tarots          []Tarot
	Aptitudes       []Aptitude
	Characteristics []Characteristic
	Gauges          []Gauge
	Skills          []Skill
	Talents         []Talent
}

func ParseUniverse(file io.Reader) (Universe, error) {
	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return Universe{}, fmt.Errorf("unable to read universe: %s", err.Error())
	}
	
	universe := Universe{}
	err := json.Unmarshall(file, &universe)
	if err != nil {
		return Universe{}, fmt.Errorf("unable to parse universe: %s", err.Error())
	}
	
	return universe, nil
}
