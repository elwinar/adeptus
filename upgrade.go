interface Upgrade {
}

const(
	IS_A_VALUE = "[+-]?(\d)*$"
)

type factory func(string) (Upgrade, err)

var factories [string]factory

func init() {
	factories['characteristic'] = NewCharacteristic;
	factories['skill'] = NewSkill;
	factories['talent'] = NewTalent;
}

// returns the map and the remain of the given string
func chunkMark(raw string) (string, string, error) {
	raw = strings.TrimSpace(raw)
	split := strings.SplitN(raw, " ", 2)
	if len(split) <= 1 {
		err := fmt.Errorf("Incorrect format for raw. Expected \" \" in string")
		return nil, nil, err
	}
	mark := split[0]
	raw = split[len(split)-1]
	return mark, raw, nil
}

// returns the xp and the remain of the given string
func chunkXp(raw string) (*int, string, error) {	
	split := strings.Split("(", raw)
	if len(split) == 1 {
		return nil, raw, nil
	}
	if len(split) != 2 {
		err = fmt.Errorf("Incorrect format for raw. Expected at most one \"(\"")
		return nil, nil, err
	}
	raw = split[0]
	experience := strings.TrimSpace(split[1])
	experience = experience[:-3]
	tmp, err := strconv.ParseInt(experience, 10, 32)
	if err != nil {
		return nil, nil, err
	}
	xp = int8(tmp)
	return &xp, raw, nil
	
}

// returns the value and the remain of the given string
func chunkValue(raw string) (*int, string, error) {	
	raw = strings.TrimSpace(raw)
	split := strings.Split(" ", raw)
	val := split[len(split) -1]
	if !regexp.Match(IS_A_VALUE, val) {
		return nil, raw, nil
	}
	if val[0] == "+" || val[0] == "-" {
		val = val[1:]
	}
	tmp, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return nil, nil, err
	}
	value = int8(tmp)
	return &value, raw, nil
}

// Transform the given line into an upgrade
func NewUpgrade(line string) (Upgrade, error) {
	raw := strings.TrimSpace(line)
	mark, raw, err := chunkMark(raw)
	if err != nil {
		err = fmt.Errorf("Incorrect upgrade in session.")
		return nil, err
	}
	
	xp, raw, err := chunkXp(raw)
	if err != nil {
		err = fmt.Errorf("Incorrect upgrade in session.")
		return nil, err
	}
	
	value, raw, err := chunkValue(raw)
	if err != nil {
		err = fmt.Errorf("Incorrect upgrade in session.")
		return nil, err
	}
	
	name := strings.TrimSpace(raw)
	a, err := GetAttributeByName(name)
	if err != nil {
		type = "custom"
	}
	else {
		type = a.type
	}
	
	return nil, nil
}