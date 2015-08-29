package parser

const ()

type Error struct {
	Line int
	Code int
}

func NewError(line, code int) Error {
	return Error{
		Line: line,
		Code: code,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("line %d: %d", e.Line, e.Code)
}
