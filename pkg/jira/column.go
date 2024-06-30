package jira

// Column represents a column in a filter.
type Column struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func (c Column) String() string {
	return c.Label
}
