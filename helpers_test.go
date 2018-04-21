package regexsolver

import (
	"regexp"
	"testing"
)

func TestIsSolved(t *testing.T) {
	tables := []struct {
		puzzle   RegexPuzzle
		expected bool
	}{
		{
			RegexPuzzle{
				[]RegexCell{{"a"}, {"b"}, {"c"}},
				[]RegexRow{
					{[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{regexp.MustCompile("123")}},
				},
			}, false,
		},
		{
			RegexPuzzle{
				[]RegexCell{{"a"}, {"b"}, {"c"}},
				[]RegexRow{
					{[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{regexp.MustCompile("abc")}},
				},
			}, true,
		},
		{
			// Sample table from https://regexcrossword.com/challenges/beginner/puzzles/1
			RegexPuzzle{
				[]RegexCell{{"h"}, {"e"}, {"l"}, {"p"}},
				[]RegexRow{
					{[]*RegexCell{{"h"}, {"e"}}, []*regexp.Regexp{regexp.MustCompile("he|ll|o+")}},
					{[]*RegexCell{{"l"}, {"p"}}, []*regexp.Regexp{regexp.MustCompile("[please]+")}},
					{[]*RegexCell{{"h"}, {"l"}}, []*regexp.Regexp{regexp.MustCompile("[^speak]+")}},
					{[]*RegexCell{{"e"}, {"p"}}, []*regexp.Regexp{regexp.MustCompile("ep|ip|ef")}},
				},
			}, true,
		},
	}

	for _, table := range tables {
		if isSolved(&table.puzzle) != table.expected {
			t.Errorf("helper isSolved returned unexpected result, got %v for table %s", !table.expected, table.puzzle)
		}
	}
}

func TestTestEq(t *testing.T) {
	r1 := regexp.MustCompile("123")
	r2 := regexp.MustCompile("456")
	tables := []struct {
		regex1   []*regexp.Regexp
		regex2   []*regexp.Regexp
		expected bool
	}{
		{nil, nil, true},
		{[]*regexp.Regexp{r1}, nil, false},
		{[]*regexp.Regexp{r1}, []*regexp.Regexp{r1}, true},
		{[]*regexp.Regexp{r1}, []*regexp.Regexp{r1, r2}, false},
	}
	for _, table := range tables {
		if testEq(table.regex1, table.regex2) != table.expected {
			t.Errorf("Got unexpected result when comparing %s and %s. Got %v", table.regex1, table.regex2, !table.expected)
		}
	}
}
