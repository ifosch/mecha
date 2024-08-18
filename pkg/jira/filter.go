package jira

// Filter represents a Jira Filter.
type Filter struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Key  string `json:"key"`
	JQL  string `json:"jql"`
	c    *Client
}

// GetColumns returns the list of columns for the filter, or an error.
func (f *Filter) GetColumns() (*ColumnList, error) {
	return f.c.GetFilterColumns(f.ID)
}

// GetRows returns the list of issues for the search of the related
// JQL, or an error.
func (f *Filter) GetIssues() (*IssueList, error) {
	cl, err := f.GetColumns()
	if err != nil {
		return nil, err
	}
	return f.c.Search(f.JQL, cl)
}

func (f Filter) String() string {
	return f.Name
}
