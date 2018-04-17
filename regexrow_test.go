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
