package jira

// Board represents a Jira Board.
type Board struct{
	ID int `json:"id"`
	Name string `json:"name"`
}
