package main

type Characteristic struct {
	label  string
	cost   int
	ignore bool
	value  int
}

// Transform the given line into an characteristic
func NewCharacteristic(line string) (Upgrade, error) {
	c := Characteristic{}
	return c, nil
}
