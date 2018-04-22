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
	// TODO figure out how to make a table for this
	var row1 RegexRow
	var row2 RegexRow
	row1.AddCell(&RegexCell{content: "a"})
	row1.AddCell(&RegexCell{content: "b"})
	row1.AddCell(&RegexCell{content: "c"})
	row1.AddExpression(regexp.MustCompile("[a-z]+"))
	row1.AddExpression(regexp.MustCompile("[^0-9]+"))
	result1 := row1.IsValidRow()
	expectedResult1 := true
	if result1 != expectedResult1 {
		t.Errorf("Wrong result got in validity test. Got %v, expected %v", result1, expectedResult1)
	}

	row2.AddCell(&RegexCell{content: "a"})
	row2.AddCell(&RegexCell{content: "1"})
	row2.AddCell(&RegexCell{content: "H"})
	row2.AddExpression(regexp.MustCompile("[b-h]+"))
	result2 := row2.IsValidRow()
	expectedResult2 := false
	if result2 != expectedResult2 {
		t.Errorf("Wrong result got in validity test. Got %v, expected %v", result2, expectedResult2)
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
