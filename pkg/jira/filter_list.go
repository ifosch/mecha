package jira

import (
	"fmt"
)

// FilterList is a list of Filters.
type FilterList []Filter

// String is the Stringer implementation.
func (fl FilterList) String() string {
	output := ""

	for _, f := range fl {
		output += fmt.Sprintf("- %v: %v\n", f.ID, f.Name)
	}

	return output
}
