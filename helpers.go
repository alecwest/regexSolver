package regexsolver

import (
	"regexp"
)

func isSolved(p *RegexPuzzle) bool {
	for _, row := range p.CellRows {
		if !row.IsValidRow() || !row.IsFull() {
			return false
		}
	}
	return true
}

func isEqRegex(a, b []*regexp.Regexp) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isEqCell(a, b *RegexCell) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.GetCellContent() == b.GetCellContent()
}

func isEqCells(a, b []*RegexCell) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].GetCellContent() != b[i].GetCellContent() {
			return false
		}
	}
	return true
}

func isEqRows(a, b *RegexRow) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a.Cells) != len(b.Cells) || len(a.Expressions) != len(b.Expressions) {
		return false
	}
	if !isEqCells(a.Cells, b.Cells) || !isEqRegex(a.Expressions, b.Expressions) {
		return false
	}
	return true
}

func isValid(puzzle *RegexPuzzle) bool {
	for _, row := range puzzle.CellRows {
		if !row.IsValidRow() {
			return false
		}
	}
	return true
}

func isValidWithNewCell(cell *RegexCell, puzzle *RegexPuzzle) bool {
	nextCell := puzzle.NextCell()
	if nextCell == nil {
		return false
	}
	nextCell.SetCellContent(cell.GetCellContent())
	result := isValid(puzzle)
	nextCell.SetCellContent("")
	return result
}
