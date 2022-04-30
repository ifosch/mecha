package jira

// Issue represents a Jira Issue.
type Issue struct{
	Key string `json:"key"`
	Fields struct{
		Status struct{
			Name string `json:"name"`
		} `json:"status"`
		StoryPoints string
	} `json:"fields"`
}
