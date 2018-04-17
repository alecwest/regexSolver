package regexsolver

import (
	"testing"
)

func TestValidCellContent(t *testing.T) {
	var cell RegexCell
	tables := []struct {
		content      string
		contentAfter string
		expectError  bool
	}{
		{"a", "a", false},
		{"", "", false},
		{"ab", "", true},
	}
	for _, table := range tables {
		err := cell.SetCellContent(table.content)
		if err != nil && !table.expectError {
			t.Errorf("SetCellContent failed unexpectedly. Passed: %s, got: %s", table.content, err.Error())
		} else if err == nil && table.expectError {
			t.Errorf("SetCellContent didn't fail as expected. Passed: %s", table.content)
		} else if cell.GetCellContent() != table.contentAfter {
			t.Errorf("SetCellContent didn't set cell content as expected. Passed: %s, expected: %s", table.content, table.contentAfter)
		}
	}
}
