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
		t.Errorf("Unexpected number of cells in RegexPuzzle. Got %d, expected %d, rp is %s", resultNumCells, expectedNumCells, rp)
	}
	resultNumInRow1 := len(rp.CellRows[0].Cells)
	expectedNumInRow1 := 2
	if resultNumInRow1 != expectedNumInRow1 {
		t.Errorf("Unexpected number of cells in RegexRow. Got %d, expected %d, rp is %s", resultNumInRow1, expectedNumInRow1, rp)
	}
	resultNumInRow2 := len(rp.CellRows[1].Cells)
	expectedNumInRow2 := 1
	if resultNumInRow2 != expectedNumInRow2 {
		t.Errorf("Unexpected number of cells in RegexRow. Got %d, expected %d, rp is %s", resultNumInRow2, expectedNumInRow2, rp)
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

func TestNextRow(t *testing.T) {
	expectedRow1 := RegexRow{[]*RegexCell{{"a"}, {"b"}, {""}}, []*regexp.Regexp{}}
	expectedRow2 := RegexRow{[]*RegexCell{{""}, {"b"}, {""}}, []*regexp.Regexp{}}
	expectedRow3 := RegexRow{[]*RegexCell{{""}, {"p"}}, []*regexp.Regexp{regexp.MustCompile("[please]+")}}
	tables := []struct {
		puzzle   RegexPuzzle
		expected *RegexRow
	}{
		{
			RegexPuzzle{
				[]RegexCell{{"a"}, {"b"}, {"c"}},
				[]RegexRow{{[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{}}}},
			nil,
		},
		{
			RegexPuzzle{
				[]RegexCell{{"a"}, {"b"}, {"c"}},
				[]RegexRow{{[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{}}, expectedRow1, {[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{}}}},
			&expectedRow1,
		},
		{
			RegexPuzzle{
				[]RegexCell{{"a"}, {"b"}, {"c"}},
				[]RegexRow{{[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{}}, expectedRow2, {[]*RegexCell{{"a"}, {"b"}, {"c"}}, []*regexp.Regexp{}}}},
			&expectedRow2,
		},
		{
			// Sample table from https://regexcrossword.com/challenges/beginner/puzzles/1
			RegexPuzzle{
				[]RegexCell{{"h"}, {"e"}, {"l"}, {"p"}},
				[]RegexRow{{[]*RegexCell{{"h"}, {"e"}}, []*regexp.Regexp{regexp.MustCompile("he|ll|o+")}}, expectedRow3, {[]*RegexCell{{"h"}, {""}}, []*regexp.Regexp{regexp.MustCompile("[^speak]+")}}, {[]*RegexCell{{"e"}, {"p"}}, []*regexp.Regexp{regexp.MustCompile("ep|ip|ef")}}}},
			&expectedRow3,
		},
	}

	for _, table := range tables {
		result := table.puzzle.NextRow()
		if !isEqRows(result, table.expected) {
			t.Errorf("Got unexpected row from NextRow function. Got %s, expected %s", result, table.expected)
		}
	}
}
