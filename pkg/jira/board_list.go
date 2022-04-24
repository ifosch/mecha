package jira

// BoardList represents a list of Boards as returned by the Jira REST API.
type BoardList struct{
	Values []Board `json:"values"`
}
