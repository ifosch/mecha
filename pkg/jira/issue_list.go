package jira

import (
	"fmt"
	"strings"
)

// IssueList represents a list of Jira Issues as returned by the Jira API.
type IssueList struct {
	Issues []*Issue `json:"issues"`
}

// FilterStatus returns a filtered `IssueList` with the issues with
// states specified in the `states` argument (allows comma separated
// values), or an error.
func (il *IssueList) FilterStatus(status string) (*IssueList, error) {
	statuses := strings.Split(status, ",")

	filtered := &IssueList{}
	for _, i := range il.Issues {
		for _, status := range statuses {
			if i.Fields.Status.Name == status {
				filtered.Issues = append(filtered.Issues, i)
			}
		}
	}
	return filtered, nil
}

// String implements Stringer interface.
func (il *IssueList) String() string {
	output := ""

	for _, i := range il.Issues {
		output += fmt.Sprintf("  - %v %v (%v)\n", i.Key, i.Fields.Status.Name, i.Fields.StoryPoints)
	}

	return output
}
