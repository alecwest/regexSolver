package regexsolver

import (
	"regexp"
	"testing"
)

func TestTestEq(t *testing.T) {
	r1 := regexp.MustCompile("123")
	r2 := regexp.MustCompile("456")
	tables := []struct {
		regex1   []*regexp.Regexp
		regex2   []*regexp.Regexp
		expected bool
	}{
		{nil, nil, true},
		{
			[]*regexp.Regexp{
				r1,
			},
			nil,
			false,
		},
		{
			[]*regexp.Regexp{
				r1,
			},
			[]*regexp.Regexp{
				r1,
			},
			true,
		},
		{
			[]*regexp.Regexp{
				r1,
			},
			[]*regexp.Regexp{
				r1,
				r2,
			},
			false,
		},
	}
	for _, table := range tables {
		if testEq(table.regex1, table.regex2) != table.expected {
			t.Errorf("Got unexpected result when comparing %s and %s. Got %v", table.regex1, table.regex2, !table.expected)
		}
	}
}
