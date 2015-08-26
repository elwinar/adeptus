package main

import (
	"fmt"
)

type Upgrade interface {
}

type rawUpgrade struct {
	mark       string
	label      string
	value      string
	experience string
}

func (r rawUpgrade) Format() string {
	return fmt.Sprintf("rawUpgrade\n\tmark: %s\n\tlabel: %s\n\tvalue: %s\n\texperience: %s\n", r.mark, r.label, r.value, r.experience)
}
