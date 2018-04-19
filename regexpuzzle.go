package regexsolver

import (
	"regexp"
)

// RegexPuzzle represents a two-dimensional object containing multiple cells.
// The RegexPuzzle is deemed valid if each expression applied to individual cells
// validates for itself and for all within
type RegexPuzzle struct {
	Cells    []RegexCell
	CellRows []RegexRow
}

func (rp *RegexPuzzle) Solve() {

}

// DeclareRow takes in all regex and applies it to a new row.
// The new row is added to the puzzle in order, with no new cells declared
func (rp *RegexPuzzle) DeclareRow(regex ...*regexp.Regexp) {
	var row RegexRow
	for _, r := range regex {
		row.AddExpression(r)
	}
	rp.CellRows = append(rp.CellRows, row)
}

// DeclareCell takes in all rows (parents) associated with a cell,
// assigns them each the new empty cell, and adds the cell to the
// full list of cells.
func (rp *RegexPuzzle) DeclareCell(parents ...*RegexRow) {
	var cell RegexCell
	for _, p := range parents {
		p.AddCell(&cell)
	}
	rp.Cells = append(rp.Cells, cell)
}

// GetRowByRegex returns the row that is associated with all of the passed regex
func (rp *RegexPuzzle) GetRowByRegex(regex ...*regexp.Regexp) *RegexRow {
	for _, row := range rp.CellRows {
		match := 0
		for _, expr := range row.Expressions {
			for _, reg := range regex {
				if expr == reg {
					match = match + 1
				}
				if match == len(regex) {
					return &row
				}
			}
		}
	}
	return nil
}
