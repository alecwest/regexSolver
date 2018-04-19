package regexsolver

import (
	"regexp"
	"testing"
)

func TestDeclareRow(t *testing.T) {
	var rp RegexPuzzle
	rp.DeclareRow(regexp.MustCompile("abc"), regexp.MustCompile("bcd"))
	result := len(rp.CellRows)
	expected := 1
	if result != expected {
		t.Errorf("Unexpected number of elements in RegexPuzzle. Got %d, expected %d", result, expected)
	}
}

func TestDeclareCell(t *testing.T) {
	var rp RegexPuzzle
	regex1 := regexp.MustCompile("abc")
	regex2 := regexp.MustCompile("bcd")
	regex3 := regexp.MustCompile("123")
	regex4 := regexp.MustCompile("234")
	rp.DeclareRow(regex1, regex2)
	rp.DeclareRow(regex3, regex4)
	rp.DeclareCell(
		rp.GetRowByRegex(regex1, regex2),
		rp.GetRowByRegex(regex4, regex3), // Order shouldn't matter
	)
	rp.DeclareCell(
		rp.GetRowByRegex(regex2, regex1),
	)
	resultNumCells := len(rp.Cells)
	expectedNumCells := 2
	if resultNumCells != expectedNumCells {
		t.Errorf("Unexpected number of cells in RegexPuzzle. Got %d, expected %d", resultNumCells, expectedNumCells)
	}
	resultNumInRow1 := len(rp.GetRowByRegex(regex1, regex2).Cells)
	expectedNumInRow1 := 2
	if resultNumInRow1 != expectedNumInRow1 {
		t.Errorf("Unexpected number of cells in RegexRow. Got %d, expected %d", resultNumInRow1, expectedNumInRow1)
	}
	resultNumInRow2 := len(rp.GetRowByRegex(regex3, regex4).Cells)
	expectedNumInRow2 := 1
	if resultNumInRow2 != expectedNumInRow2 {
		t.Errorf("Unexpected number of cells in RegexRow. Got %d, expected %d", resultNumInRow2, expectedNumInRow2)
	}
}

func TestGetRowByRegex(t *testing.T) {
	var rp RegexPuzzle
	regex1 := regexp.MustCompile("abc")
	regex2 := regexp.MustCompile("bcd")
	regex3 := regexp.MustCompile("123")
	regex4 := regexp.MustCompile("234")
	rp.DeclareRow(regex1, regex2)
	rp.DeclareRow(regex3, regex4)

	row := rp.GetRowByRegex(regex1, regex2)
	if row == nil {
		t.Errorf("Row was not found!")
	}
	row2 := rp.GetRowByRegex(regex4, regex3)
	if row2 == nil {
		t.Errorf("Row was not found!")
	}
	row3 := rp.GetRowByRegex(regex1, regex4)
	if row3 != nil {
		t.Errorf("Row should not have been found!")
	}
}
