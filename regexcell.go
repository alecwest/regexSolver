package regexsolver

import "fmt"

// RegexCell represents a single character, and all of the regex expressions
// that are applied to it to confirm its validity
type RegexCell struct {
	content string
}

// GetCellContent returns the cell's contents
func (rc *RegexCell) GetCellContent() string {
	return rc.content
}

// SetCellContent sets the RegexCell content value
// Value must be no more than one character in length
func (rc *RegexCell) SetCellContent(c string) error {
	if rc == nil {
		return fmt.Errorf(fmt.Sprintf("Cell was nil"))
	} else if len(c) <= 1 {
		rc.content = c
	} else {
		return fmt.Errorf(fmt.Sprintf("Cell content cannot be greater than one character in length, but got %s", c))
	}
	return nil
}
