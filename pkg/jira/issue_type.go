package jira

// IssueType represents issue types for Jira.
type IssueType struct {
	Name string `json:"name"`
}

// String implements Stringer interface.
func (it IssueType) String() string {
	return it.Name
}
