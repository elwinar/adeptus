package adeptus

import(
	"fmt"
	"strings"
)

type Header struct {
	Name		string
	Origin		string
	Background	string
	Role		string
	Tarot		string
}

func ParseHeader(block []Line) (Header, error) {
	
	header := Header{}
	for _, line := range block {
		fields := strings.SplitN(block, ":", 2)
		if len(fields) < 2 {
			return Header{}, fmt.Errorf("Error on line %d: Expected pair key:value.", line.Number)
		}
		
		key := strings.TrimSpace(strings.ToLower(fields[0]))
		value := strings.TrimSpace(fields[1])
		switch key {
			case "name":
				header.Name = value
			case "origin":
				header.Origin = value
			case "background":
				header.Background = value
			case "role":
				header.Role = value
			case "tarot":
				header.Tarot = value
			default:
				return Header{}, fmt.Errorf("Error on line %d: Undefined key: \"%s\".", line.Number, key)
		}
	}

	return header, nil
}