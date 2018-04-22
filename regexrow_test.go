package regexsolver

import (
	"regexp"
	"testing"
)

func TestAddCell(t *testing.T) {
	var row RegexRow
	tables := []struct {
		cell        *RegexCell
		numExpected int
	}{
		{&RegexCell{content: "a"}, 1},
		{&RegexCell{content: ""}, 2},
	}
	for _, table := range tables {
		row.AddCell(table.cell)
		res := len(row.Cells)
		if res != table.numExpected {
			t.Errorf("RegexCell was not sucessfully added. Counted %d, expected %d", res, table.numExpected)
		}
	}
}

func TestAddExpression(t *testing.T) {
	var row RegexRow
	tables := []struct {
		expression  *regexp.Regexp
		numExpected int
	}{
		{regexp.MustCompile("a"), 1},
		{regexp.MustCompile("[a-z]"), 2},
	}
	for _, table := range tables {
		row.AddExpression(table.expression)
		res := len(row.Expressions)
		if res != table.numExpected {
			t.Errorf("Regex expression was not successufully added. Counted %d, expected %d", res, table.numExpected)
		}
	}
}

func TestIsFull(t *testing.T) {
	tables := []struct {
		row      RegexRow
		expected bool
	}{
		{
			RegexRow{[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{}},
			true,
		},
		{
			RegexRow{[]*RegexCell{{"a"}, {""}, {"c"}}, []*regexp.Regexp{}},
			false,
		},
		{
			RegexRow{[]*RegexCell{{""}, {""}, {""}}, []*regexp.Regexp{}},
			false,
		},
	}
	for _, table := range tables {
		if table.row.IsFull() != table.expected {
			t.Errorf("Unexpected result when calling IsFull() on %s, expected %v", table.row, table.expected)
		}
	}
}

func TestIsValidRow(t *testing.T) {
	tables := []struct {
		row      RegexRow
		expected bool
	}{
		{
			RegexRow{[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{regexp.MustCompile("[a-z]+"), regexp.MustCompile("[^0-9]+")}},
			true,
		},
		{
			RegexRow{[]*RegexCell{{"a"}, {"1"}, {"H"}}, []*regexp.Regexp{regexp.MustCompile("[b-h]+")}},
			false,
		},
	}
	for _, table := range tables {
		if table.row.IsValidRow() != table.expected {
			t.Errorf("Unexpected result from IsValidRow function on row %s. Expected %v", table.row, table.expected)
		}
	}
}

func TestJoin(t *testing.T) {
	var row1 RegexRow
	row1.AddCell(&RegexCell{content: "a"})
	row1.AddCell(&RegexCell{content: "b"})
	row1.AddCell(&RegexCell{content: "c"})
	result1 := row1.join()
	expectedResult1 := "abc"
	if result1 != expectedResult1 {
		t.Errorf("Wrong result got in join test. Got %v, expected %v", result1, expectedResult1)
	}
}
