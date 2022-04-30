package jira

import (
	"fmt"
)

// IssueList represents a list of Jira Issues as returned by the Jira API.
type IssueList struct {
	Issues []*Issue `json:"issues"`
}

// String implements Stringer interface.
func (il *IssueList) String() string {
	output := ""

	for _, i := range il.Issues {
		output += fmt.Sprintf("  - %v %v (%v)\n", i.Key, i.Fields.Status.Name, i.Fields.StoryPoints)
	}

	return output
}
