package jira

import (
	"encoding/json"
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
	p *Project
	c *Client
}

// Complete sets the sprint as complete, or returns an error.
func (s *Sprint) Complete() error {
	completeSprintInputData := &completeSprintInput{
		State: "closed",
	}
	completeSprintInputJSON, err := json.Marshal(completeSprintInputData)
	if err != nil {
		return err
	}

	_, err = s.c.post(
		fmt.Sprintf("/rest/agile/1.0/sprint/%v", s.ID),
		completeSprintInputJSON,
	)
	if err != nil {
		return err
	}

	return nil
}

type completeSprintInput struct{
	State string `json:"state"`
}

// MoveIssuesToNextSprint moves all issues defined in `issues`
// argument to next future sprint, or returns an error.
func (s *Sprint) MoveIssuesToNextSprint(issues *IssueList) error {
	future, err := s.p.GetNextSprint()
	if err != nil {
		return err
	}

	moveIssuesInputData := &moveIssuesInput{
		Issues: []string{},
	}
	for _, i := range issues.Issues {
		moveIssuesInputData.Issues = append(moveIssuesInputData.Issues, i.Key)
	}
	moveIssuesInputJSON, err := json.Marshal(moveIssuesInputData)
	if err != nil {
		return err
	}

	_, err = s.c.post(
		fmt.Sprintf("/rest/agile/1.0/sprint/%v/issue", future.ID),
		moveIssuesInputJSON,
	)
	return err
}

type moveIssuesInput struct{
	Issues []string `json:"issues"`
}

// GetIssuesByStatus returns an `*IssueList` with the issues matching
// the specified status (can be a comma separated list), or an error.
func (s *Sprint) GetIssuesByStatus(statuses string) (*IssueList, error) {
	il, err := s.GetIssues()
	if err != nil {
		return nil, err
	}

	filtered, err := il.FilterStatus(statuses)
	if err != nil {
		return nil, err
	}

	return filtered, err
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
	container, err := s.c.getContainer(fmt.Sprintf("/rest/agile/1.0/sprint/%v/issue", s.ID))
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
