package jira

import (
	"fmt"
	"os"
	"strings"
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
	c *Client
}

// GetIssues returns all the issues found for the Sprint, or an error.
func (s *Sprint) GetIssues() (*IssueList, error) {
	container, err := s.c.getContainer("GET", fmt.Sprintf("/rest/agile/1.0/sprint/%v/issue", s.ID))
	if err != nil {
		return nil, err
	}

	il := IssueList{}
	for _, issue := range container.S("issues").Children() {
		path := fmt.Sprintf("fields.%s", os.Getenv("JIRA_SP_FIELD"))
		SP := issue.Path(path)
		i := Issue{
			Key: strings.ReplaceAll(fmt.Sprintf("%s", issue.Path("key")), "\"", ""),
			Fields: struct{
				Status struct{
					Name string `json:"name"`
				} `json:"status"`
				StoryPoints string
			}{
				Status: struct{
					Name string `json:"name"`
				}{
					Name: strings.ReplaceAll(fmt.Sprintf("%s", issue.Path("fields.status.name")), "\"", ""),
				},
				StoryPoints: fmt.Sprintf("%s", SP),
			},
		}
		il.Issues = append(il.Issues, &i)
	}

	return &il, nil
}
