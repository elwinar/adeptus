interface Upgrade {
}

type factory func([]line) (Upgrade, err)

var factories [string]factory

func init() {
	factories['characteristic'] = NewCharacteristic;
	factories['skill'] = NewSkill;
	factories['talent'] = NewTalent;
}

// Transform the given line into an upgrade
func Parse(line string, []data) (Upgrade, error) {
}