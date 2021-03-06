package regexsolver

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/jinzhu/copier"

	log "gopkg.in/inconshreveable/log15.v2"
)

// RegexPuzzle represents a two-dimensional object containing multiple cells.
// The RegexPuzzle is deemed valid if each expression applied to individual cells
// validates for itself and for all within
type RegexPuzzle struct {
	Cells    []*RegexCell
	CellRows []*RegexRow
}

// RandomSolve will pick letters randomly until the puzzle is solved
func (rp *RegexPuzzle) RandomSolve() {
	vals := strings.Split("abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(){}[]/=\\?+|-_',.\"<>`~", "")
	for !isSolved(rp) {
		for _, cell := range rp.Cells {
			cell.SetCellContent("")
		}
		for _, cell := range rp.Cells {
			cell.SetCellContent(vals[rand.Intn(len(vals)-1)])
		}
	}
}

// Solve will run a recursive backtracking algorithm to solve the puzzle
func (rp *RegexPuzzle) Solve() {
	rp = rp.solve(*rp)
}

func (rp *RegexPuzzle) solve(p RegexPuzzle) *RegexPuzzle {
	var nextCell RegexCell
	vals := strings.Split("abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*(){}[]/=\\?+|-_',.\"<>`~", "")
	temp := p.NextCell()
	if temp == nil {
		if isSolved(&p) {
			return &p
		}
		return nil
	}
	copier.Copy(nextCell, temp)

	for _, char := range vals {
		nextCell.SetCellContent(char)
		log.Debug(fmt.Sprintf("New cell with content %s is being added to the puzzle", nextCell.GetCellContent()))
		log.Debug(fmt.Sprintf("puzzle is %s", p))
		if isValidWithNewCell(&nextCell, &p) {
			p.SetNextCell(nextCell.GetCellContent())
			return rp.solve(p)
		}
	}

	return nil
}

// DeclareRow takes in all regex and applies it to a new row.
// The new row is added to the puzzle in order, with no new cells declared
func (rp *RegexPuzzle) DeclareRow(regex ...*regexp.Regexp) {
	row := &RegexRow{[]*RegexCell{}, []*regexp.Regexp{}}
	for _, r := range regex {
		row.AddExpression(r)
	}
	rp.CellRows = append(rp.CellRows, row)
}

// DeclareCell takes in all rows (parents) associated with a cell,
// assigns them each the new empty cell, and adds the cell to the
// full list of cells.
func (rp *RegexPuzzle) DeclareCell(parents ...*RegexRow) {
	cell := &RegexCell{}
	for _, p := range parents {
		for i := range rp.CellRows {
			if isEqRegex(p.Expressions, rp.CellRows[i].Expressions) {
				rp.CellRows[i].AddCell(cell)
				break
			}
		}
	}
	cell.SetCellContent(string(len(rp.Cells)))
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
					return row
				}
			}
		}
	}
	return nil
}

// NextCell returns the first cell that is not filled in
func (rp *RegexPuzzle) NextCell() *RegexCell {
	for _, cell := range rp.Cells {
		if len(cell.GetCellContent()) == 0 {
			return cell
		}
	}
	return nil
}

// NextRow returns the first row that is not filled in
func (rp *RegexPuzzle) NextRow() *RegexRow {
	for _, row := range rp.CellRows {
		if !row.IsFull() {
			return row
		}
	}
	return nil
}

// SetNextCell takes a string input and assigns it to the next
// empty cell.
func (rp *RegexPuzzle) SetNextCell(c string) error {
	err := rp.NextCell().SetCellContent(c)
	return err
}
