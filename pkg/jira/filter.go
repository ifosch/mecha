package jira

// Filter represents a Jira Filter.
type Filter struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Key  string `json:"key"`
	c    *Client
}

// GetColumns returns the list of columns for the filter, or an error.
func (f *Filter) GetColumns() (*ColumnList, error) {
	return f.c.GetFilterColumns(f.ID)
}

func (f Filter) String() string {
	return f.Name
}
