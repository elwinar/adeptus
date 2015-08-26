package adeptus

type Sheet struct {
	Header map[string]string
	Sessions []Session
}

func ParseSheet(file io.Reader) (Sheet, error) {
	return Sheet{}, nil
}
