package regexsolver

import "testing"

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
		if len(row.Cells) != table.numExpected {
			t.Errorf("RegexCell was not sucessfully added. Counted %d, expected %d", len(row.Cells), table.numExpected)
		}
	}
}
