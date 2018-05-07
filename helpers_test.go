package regexsolver

import (
	"regexp"
	"testing"
)

func TestIsSolved(t *testing.T) {
	tables := []struct {
		puzzle   *RegexPuzzle
		expected bool
	}{
		{
			&RegexPuzzle{
				[]*RegexCell{{"a"}, {"b"}, {"c"}},
				[]*RegexRow{
					{[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{regexp.MustCompile("123")}},
				},
			}, false,
		},
		{
			&RegexPuzzle{
				[]*RegexCell{{"a"}, {"b"}, {"c"}},
				[]*RegexRow{
					{[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{regexp.MustCompile("abc")}},
				},
			}, true,
		},
		{
			// Sample table from https://regexcrossword.com/challenges/beginner/puzzles/1
			&RegexPuzzle{
				[]*RegexCell{{"h"}, {"e"}, {"l"}, {"p"}},
				[]*RegexRow{
					{[]*RegexCell{{"h"}, {"e"}}, []*regexp.Regexp{regexp.MustCompile("he|ll|o+")}},
					{[]*RegexCell{{"l"}, {"p"}}, []*regexp.Regexp{regexp.MustCompile("[please]+")}},
					{[]*RegexCell{{"h"}, {"l"}}, []*regexp.Regexp{regexp.MustCompile("[^speak]+")}},
					{[]*RegexCell{{"e"}, {"p"}}, []*regexp.Regexp{regexp.MustCompile("ep|ip|ef")}},
				},
			}, true,
		},
	}

	for _, table := range tables {
		if isSolved(table.puzzle) != table.expected {
			t.Errorf("helper isSolved returned unexpected result, got %v for table %s", !table.expected, table.puzzle)
		}
	}
}

func TestIsEqRegex(t *testing.T) {
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
		if isEqRegex(table.regex1, table.regex2) != table.expected {
			t.Errorf("Got unexpected result when comparing %s and %s. Got %v", table.regex1, table.regex2, !table.expected)
		}
	}
}

func TestIsEqCell(t *testing.T) {
	tables := []struct {
		cell1    *RegexCell
		cell2    *RegexCell
		expected bool
	}{
		{nil, nil, true},
		{&RegexCell{"a"}, nil, false},
		{&RegexCell{"a"}, &RegexCell{"a"}, true},
		{&RegexCell{"a"}, &RegexCell{"b"}, false},
	}
	for _, table := range tables {
		if isEqCell(table.cell1, table.cell2) != table.expected {
			t.Errorf("Unexpected result from isEqCells when comparing %s and %s. Got %v", table.cell1, table.cell2, !table.expected)
		}
	}
}

func TestIsEqCells(t *testing.T) {
	c1 := RegexCell{"a"}
	c2 := RegexCell{"b"}
	c3 := RegexCell{"c"}
	tables := []struct {
		row1     *RegexRow
		row2     *RegexRow
		expected bool
	}{
		{&RegexRow{nil, []*regexp.Regexp{}}, &RegexRow{nil, []*regexp.Regexp{}}, true},
		{&RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{}}, &RegexRow{nil, []*regexp.Regexp{}}, false},
		{&RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{}}, &RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{}}, true},
		{&RegexRow{[]*RegexCell{&c1, &c3}, []*regexp.Regexp{}}, &RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{}}, false},
		{&RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{}}, &RegexRow{[]*RegexCell{&c1, &RegexCell{"d"}, &c3}, []*regexp.Regexp{}}, false},
	}
	for _, table := range tables {
		if isEqCells(table.row1.Cells, table.row2.Cells) != table.expected {
			t.Errorf("Unexpected result from isEqCells when comparing %s and %s. Got %v", table.row1.Cells, table.row2.Cells, !table.expected)
		}
	}
}

func TestIsEqRows(t *testing.T) {
	r1 := regexp.MustCompile("123")
	r2 := regexp.MustCompile("456")
	c1 := RegexCell{"a"}
	c2 := RegexCell{"b"}
	c3 := RegexCell{"c"}
	tables := []struct {
		row1     *RegexRow
		row2     *RegexRow
		expected bool
	}{
		{nil, nil, true},
		{&RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{r1, r2}}, nil, false},
		{&RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{r1, r2}}, &RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{r1, r2}}, true},
		{&RegexRow{[]*RegexCell{&c2, &c3}, []*regexp.Regexp{r1, r2}}, &RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{r1, r2}}, false},
		{&RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{r2}}, &RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{r1, r2}}, false},
		{&RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{r1, r2}}, &RegexRow{[]*RegexCell{&c1, &RegexCell{"d"}, &c3}, []*regexp.Regexp{r1, r2}}, false},
		{&RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{r1, r2}}, &RegexRow{[]*RegexCell{&c1, &c2, &c3}, []*regexp.Regexp{regexp.MustCompile("aaa"), r2}}, false},
	}
	for _, table := range tables {
		if isEqRows(table.row1, table.row2) != table.expected {
			t.Errorf("Unexpected result from isEqRows when comparing %s and %s. Got %v", table.row1, table.row2, !table.expected)
		}
	}
}

func TestIsValid(t *testing.T) {
	tables := []struct {
		puzzle   *RegexPuzzle
		expected bool
	}{
		{
			// Sample table from https://regexcrossword.com/challenges/beginner/puzzles/1
			&RegexPuzzle{
				[]*RegexCell{{"h"}, {"e"}, {"l"}, {"p"}},
				[]*RegexRow{
					{[]*RegexCell{{"h"}, {"e"}}, []*regexp.Regexp{regexp.MustCompile("he|ll|o+")}},
					{[]*RegexCell{{"l"}, {"p"}}, []*regexp.Regexp{regexp.MustCompile("[please]+")}},
					{[]*RegexCell{{"h"}, {"l"}}, []*regexp.Regexp{regexp.MustCompile("[^speak]+")}},
					{[]*RegexCell{{"e"}, {"p"}}, []*regexp.Regexp{regexp.MustCompile("ep|ip|ef")}},
				},
			}, true,
		},
		{
			&RegexPuzzle{
				[]*RegexCell{{"h"}, {"e"}, {"l"}, {"d"}},
				[]*RegexRow{
					{[]*RegexCell{{"h"}, {"e"}}, []*regexp.Regexp{regexp.MustCompile("he|ll|o+")}},
					{[]*RegexCell{{"l"}, {"d"}}, []*regexp.Regexp{regexp.MustCompile("[please]+")}},
					{[]*RegexCell{{"h"}, {"l"}}, []*regexp.Regexp{regexp.MustCompile("[^speak]+")}},
					{[]*RegexCell{{"e"}, {"d"}}, []*regexp.Regexp{regexp.MustCompile("ep|ip|ef")}},
				},
			}, false,
		},
	}
	for _, table := range tables {
		if isValid(table.puzzle) != table.expected {
			t.Errorf("Unexpected result from isValid function call on puzzle %s, expected %v", table.puzzle, table.expected)
		}
	}
}

func TestIsValidWithNewCell(t *testing.T) {
	c1 := &RegexCell{"h"}
	c2 := &RegexCell{""}
	c3 := &RegexCell{"l"}
	c4 := &RegexCell{"p"}
	c5 := &RegexCell{"e"}
	c6 := &RegexCell{""}
	tables := []struct {
		puzzle   *RegexPuzzle
		newValue string
		expected bool
	}{
		{
			// Sample table from https://regexcrossword.com/challenges/beginner/puzzles/1
			&RegexPuzzle{
				[]*RegexCell{c1, c2, c3, c4},
				[]*RegexRow{
					{[]*RegexCell{c1, c2}, []*regexp.Regexp{regexp.MustCompile("he|ll|o+")}},
					{[]*RegexCell{c3, c4}, []*regexp.Regexp{regexp.MustCompile("[please]+")}},
					{[]*RegexCell{c1, c3}, []*regexp.Regexp{regexp.MustCompile("[^speak]+")}},
					{[]*RegexCell{c2, c4}, []*regexp.Regexp{regexp.MustCompile("ep|ip|ef")}},
				},
			}, "e", true,
		},
		{
			&RegexPuzzle{
				[]*RegexCell{c1, c5, c3, c6},
				[]*RegexRow{
					{[]*RegexCell{c1, c5}, []*regexp.Regexp{regexp.MustCompile("he|ll|o+")}},
					{[]*RegexCell{c3, c6}, []*regexp.Regexp{regexp.MustCompile("[please]+")}},
					{[]*RegexCell{c1, c3}, []*regexp.Regexp{regexp.MustCompile("[^speak]+")}},
					{[]*RegexCell{c5, c6}, []*regexp.Regexp{regexp.MustCompile("ep|ip|ef")}},
				},
			}, "d", false,
		},
	}
	for _, table := range tables {
		newCell := &RegexCell{table.newValue}
		if isValidWithNewCell(newCell, table.puzzle) != table.expected {
			t.Errorf("Unexpected result from isValidWithNewCell function call on puzzle %s with cell %s, expected %v", table.puzzle, newCell, table.expected)
		}
		if table.puzzle.NextCell() == nil {
			t.Errorf("isValidWithNewCell unintentionally added content to the puzzle %s", table.puzzle)
		}
	}
}
