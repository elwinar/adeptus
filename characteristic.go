struct Characteristic {
	label	string
	cost	int
	ignore	boolean
	value	int
}

// Transform the given line into an characteristic
func NewCharacteristic(line string) (Characteristic, error) {
	c := Characteristic{}
	return c, nil
}