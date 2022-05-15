package jira

import (
	"fmt"
)

// SprintList represents a list of Sprints, as returned by the Jira REST API.
type SprintList struct {
	Values []Sprint `json:"values"`
}

// String returns a string representation of a SprintList.
func (sl SprintList) String() string {
	output := ""

	for _, s := range sl.Values {
		output += fmt.Sprintf("- %v (%v)\n", s.Name, s.State)
	}

	return output
}
