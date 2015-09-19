package parser

// Meta is a header and a collection of associated options
type Meta struct {
	Label   string
	Options []string
}

func NewMeta(raw string) (*Meta, error) {
	return &Meta{
		Label: raw,
	}, nil
}
