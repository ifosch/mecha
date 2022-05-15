package jira

import (
	"fmt"
)

// ProjectList is a list of Projects.
type ProjectList []Project

// String is the Stringer implementation.
func (pl ProjectList) String() string {
	output := ""

	for _, p := range pl {
		output += fmt.Sprintf("- %v\n", p.Name)
	}

	return output
}
