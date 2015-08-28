package adeptus

import (
	"errors"
	"os"
	"reflect"
	"testing"
)

func successSessionParser(_ []Line) (Session, error) {
	return Session{}, nil
}
func failSessionParser(_ []Line) (Session, error) {
	return Session{}, errors.New("fail")
}
func successHeaderParser(_ []Line) (Header, error) {
	return Header{}, nil
}
func failHeaderParser(_ []Line) (Header, error) {
	return Header{}, errors.New("fail")
}

func Test_parseSheet(t *testing.T) {
	cases := []struct {
		in string
		sessionParser
		headerParser
		out Sheet
		err bool
	}{
		{
			in:            "tests/empty-file.40k",
			headerParser:  successHeaderParser,
			sessionParser: successSessionParser,
			out:           Sheet{},
			err:           true,
		},
		{
			in:            "tests/comments-only.40k",
			headerParser:  successHeaderParser,
			sessionParser: successSessionParser,
			out:           Sheet{},
			err:           true,
		},
		{
			in:            "tests/fail.40k",
			headerParser:  failHeaderParser,
			sessionParser: successSessionParser,
			out:           Sheet{},
			err:           true,
		},
		{
			in:            "tests/fail.40k",
			headerParser:  successHeaderParser,
			sessionParser: failSessionParser,
			out:           Sheet{},
			err:           true,
		},
		{
			in:            "tests/success.40k",
			headerParser:  successHeaderParser,
			sessionParser: successSessionParser,
			out:           Sheet{},
			err:           false,
		},
	}

	for i, c := range cases {
		in, err := os.Open(c.in)
		if err != nil {
			t.Fatalf("Unable to open file %s.", c.in)
		}
		out, err := parseSheet(in, c.headerParser, c.sessionParser)
		if (err != nil) != c.err {
			t.Logf("Unexpected error on case %d:", i+1)
			t.Logf("	Having %s", err)
			t.Fail()
			continue
		}
		if !reflect.DeepEqual(out, c.out) {
			t.Logf("Unexpected output on case %d:", i+1)
			t.Logf("	Expected %v", c.out)
			t.Logf("	Having %v", out)
			t.Fail()
		}
	}

}
