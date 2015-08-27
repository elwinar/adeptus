package adeptus

type Line struct {
	Number int
	Text string
}

func (l Line) isValid() bool {
	return l.Text != ""
}