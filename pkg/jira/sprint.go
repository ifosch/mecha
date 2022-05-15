package jira

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Sprint represents a Jira Sprint for a Project.
type Sprint struct {
	ID           int       `json:"id"`
	State        string    `json:"state"`
	Name         string    `json:"name"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	CompleteDate time.Time `json:"completeDate"`
	BoardID      int       `json:"originBoardId"`
	p            *Project
	c            *Client
}

func getSep(input string, seps map[string]string) (sep string) {
	for r, sep := range seps {
		match, _ := regexp.MatchString(r, input)
		if match {
			return sep
		}
	}
	return ""
}

// NextSprintName returns the next sprint name according to
// calculation rules.
func (s *Sprint) NextSprintName() (string, error) {
	sprintNameSeps := map[string]string{
		"[A-Z0-9]+ Sprint [0-9]+$":          " ",
		"[A-Z0-9]+ Sprint [0-9]{4}-[0-9]+$": "-",
	}
	sep := getSep(s.Name, sprintNameSeps)

	nameWords := strings.Split(s.Name, sep)
	sprintNumber, err := strconv.Atoi(nameWords[len(nameWords)-1])
	if err != nil {
		return "", err
	}
	sprintNumber++
	nameWords[len(nameWords)-1] = fmt.Sprintf("%v", sprintNumber)
	return strings.Join(nameWords, sep), nil
}

// Start starts the sprint, or returns an error.
func (s *Sprint) Start() error {
	updateSprintInputData := &updateSprintInput{
		State:     "active",
		StartDate: time.Now(),
		EndDate:   time.Now().Local().Add(time.Hour * 24 * 14),
	}
	updateSprintInputJSON, err := json.Marshal(updateSprintInputData)
	if err != nil {
		return err
	}

	_, err = s.c.post(
		fmt.Sprintf("/rest/agile/1.0/sprint/%v", s.ID),
		updateSprintInputJSON,
	)
	if err != nil {
		return err
	}

	return nil
}

// Complete sets the sprint as complete, or returns an error.
func (s *Sprint) Complete() error {
	updateSprintInputData := &updateSprintInput{
		State:     "closed",
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
	}
	updateSprintInputJSON, err := json.Marshal(updateSprintInputData)
	if err != nil {
		return err
	}

	_, err = s.c.post(
		fmt.Sprintf("/rest/agile/1.0/sprint/%v", s.ID),
		updateSprintInputJSON,
	)
	if err != nil {
		return err
	}

	return nil
}

type updateSprintInput struct {
	State     string    `json:"state"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
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

type moveIssuesInput struct {
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
			ss[i.Fields.Status.Name]["Stories"]++
			ss[i.Fields.Status.Name]["SP"] += SP
		} else {
			ss[i.Fields.Status.Name] = map[string]int{
				"Stories":    1,
				"SP":         SP,
				"Missing SP": 0,
			}
		}
		if SP == 0 {
			ss[i.Fields.Status.Name]["Missing SP"]++
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
			Key: strings.ReplaceAll(issue.Path("key").String(), "\"", ""),
			Fields: struct {
				Status struct {
					Name string `json:"name"`
				} `json:"status"`
				StoryPoints string
			}{
				Status: struct {
					Name string `json:"name"`
				}{
					Name: strings.ReplaceAll(issue.Path("fields.status.name").String(), "\"", ""),
				},
				StoryPoints: SP.String(),
			},
		}
		il.Issues = append(il.Issues, &i)
	}

	return &il, nil
}
