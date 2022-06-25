package scanutil

import (
	"bytes"
	"testing"
)

func TestScannNText(t *testing.T) {
	testCases := []struct {
		desc     string
		in       string
		def      string
		expected string
	}{
		{
			desc:     "empty input (%s)",
			in:       "\n",
			def:      "daniel",
			expected: "daniel",
		},
		{
			desc:     "a valid input (%s)",
			in:       "do\n",
			def:      "daniel",
			expected: "do",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			buff := new(bytes.Buffer)
			buff.WriteString(tC.in)
			in = buff
			rtext := ScannNText(tC.desc, 1, tC.def, func(s string) bool { return true })
			if rtext != tC.expected {
				// fmt.Printf("rtext: %v\n", rtext)
				t.Fail()
			}
		})
	}
}

func TestIsExt(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		valid bool
	}{
		{
			desc:  "a valid extension",
			input: ".js",
			valid: true,
		},
		{
			desc:  "an invalid extension",
			input: "js",
			valid: false,
		},
		{
			desc:  "an invalid extension",
			input: ".js.ss",
			valid: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if IsExt(tC.input) != tC.valid {
				t.Fail()
			}
		})
	}
}

func TestIsSep(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		valid bool
	}{
		{
			desc:  "a valid separator",
			input: "|",
			valid: true,
		},
		{
			desc:  "an invalid separator",
			input: " ",
			valid: false,
		},
		{
			desc:  "an invalid separator",
			input: "+",
			valid: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if IsSep(tC.input) != tC.valid {
				t.Fail()
			}
		})
	}
}
