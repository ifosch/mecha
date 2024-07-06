package jira

// IssueStatus represents issue statuses for Jira.
type IssueStatus struct {
	Name string `json:"name"`
}

// String implements Stringer interface.
func (it IssueStatus) String() string {
	return it.Name
}
