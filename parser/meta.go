package parser

import(
	"fmt"
	"strings"
)

// Meta is a header and a collection of associated options
type Meta struct {
	Label   string
	Options []string
}

func NewMeta(raw string) (Meta, error) {
	
	if strings.Contains(raw, "(") {
			return Meta{}, fmt.Errorf("incorrect meta format")
	}
	
	return Meta{
		Label: raw,
	}, nil
}
