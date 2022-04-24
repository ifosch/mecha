package jira

import (
	"time"
)

// Sprint represents a Jira Sprint for a Project.
type Sprint struct {
	ID int `json:"id"`
	State string `json:"state"`
	Name string `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate time.Time `json:"endDate"`
	CompleteDate time.Time `json:"completeDate"`
	BoardID int `json:"originBoardId"`
}
