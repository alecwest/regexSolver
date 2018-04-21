package regexsolver

import (
	"regexp"
	"strings"
)

// RegexRow represents a linear array of RegexCells and all regular expressions that
// are applied to the row, much like they would be applied to a normal string
type RegexRow struct {
	Cells       []*RegexCell
	Expressions []*regexp.Regexp
}

// AddCell adds a RegexCell to a collection of RegexCells.
func (rr *RegexRow) AddCell(rc *RegexCell) {
	rr.Cells = append(rr.Cells, rc)
}

// AddExpression adds a Regexp element to the list of Expressions that
// apply to this row.
func (rr *RegexRow) AddExpression(e *regexp.Regexp) {
	rr.Expressions = append(rr.Expressions, e)
}

// IsFull returns a boolean indicating whether or not all cells
// within the row have been set to some value
func (rr *RegexRow) IsFull() bool {
	for _, cell := range rr.Cells {
		if len(cell.GetCellContent()) == 0 {
			return false
		}
	}
	return true
}

// IsValidRow returns a boolean indicating whether or not all regular
// expressions that apply to the given row are valid against it.
func (rr *RegexRow) IsValidRow() bool {
	for _, test := range rr.Expressions {
		if !test.MatchString(rr.join()) {
			return false
		}
	}
	return true
}

func (rr *RegexRow) join() string {
	var result []string
	for _, cell := range rr.Cells {
		result = append(result, cell.GetCellContent())
	}
	return strings.Join(result, "")
}
