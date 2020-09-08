package charcount

import (
	"strings"
	"testing"
)



func TestCharcount(t *testing.T) {
	var tests = []string{"a b cdc"}
	for _, test := range tests {
		f := strings.NewReader(test)
		Charcount2(f)
	}
}