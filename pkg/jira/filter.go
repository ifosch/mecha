package jira

// Filter represents a Jira Filter.
type Filter struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Key  string `json:"key"`
}

func (f Filter) String() string {
	return f.Name
}
