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
		{&RegexCell{"a"}, 1},
		{&RegexCell{""}, 2},
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

func TestCellInRow(t *testing.T) {
	c1 := &RegexCell{"a"}
	c2 := &RegexCell{"b"}
	c3 := &RegexCell{"c"}
	c4 := &RegexCell{"d"}
	tables := []struct {
		row      *RegexRow
		cell     *RegexCell
		expected bool
	}{
		{
			&RegexRow{[]*RegexCell{{"e"}, {"f"}, c1}, []*regexp.Regexp{}},
			c1,
			true,
		},
		{
			&RegexRow{[]*RegexCell{{"e"}, {"f"}, c2}, []*regexp.Regexp{}},
			c3,
			false,
		},
		{
			&RegexRow{[]*RegexCell{c3, {"f"}, c2}, []*regexp.Regexp{}},
			c3,
			true,
		},
		{
			&RegexRow{[]*RegexCell{c4, {"f"}, c2}, []*regexp.Regexp{}},
			c3,
			false,
		},
	}
	for _, table := range tables {
		if table.row.CellInRow(table.cell) != table.expected {
			t.Errorf("Unexpected result from CellInRow for row %s and cell %s, expected %v", table.row, table.cell, table.expected)
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
	row1.AddCell(&RegexCell{"a"})
	row1.AddCell(&RegexCell{"b"})
	row1.AddCell(&RegexCell{"c"})
	result1 := row1.join()
	expectedResult1 := "abc"
	if result1 != expectedResult1 {
		t.Errorf("Wrong result got in join test. Got %v, expected %v", result1, expectedResult1)
	}
}
