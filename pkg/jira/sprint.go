package jira

import (
	"fmt"
	"os"
	"strconv"
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

// GetStats returns stats for this specific Sprint, or an error.
func (s *Sprint) GetStats() (map[string]map[string]int, error) {
	il, err := s.GetIssues()
	if err != nil {
		return nil, err
	}

	ss := map[string]map[string]int{}
	for _, i := range il.Issues {
		SP, err := strconv.Atoi(i.Fields.StoryPoints)
		if err != nil {
			SP = 0
		}

		if _, ok := ss[i.Fields.Status.Name]; ok {
			ss[i.Fields.Status.Name]["Stories"] += 1
			ss[i.Fields.Status.Name]["SP"] += SP
		} else {
			ss[i.Fields.Status.Name] = map[string]int{
				"Stories": 1,
				"SP": SP,
				"Missing SP": 0,
			}
		}
		if SP == 0 {
			ss[i.Fields.Status.Name]["Missing SP"] += 1
		}
	}

	return ss, nil
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
