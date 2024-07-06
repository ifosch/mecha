package jira

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// IssueFields represents detailed fields for Jira Issues.
type IssueFields struct {
	Created     DateString  `json:"created"`
	Updated     DateString  `json:"updated"`
	Status      IssueStatus `json:"status"`
	Summary     string      `json:"summary"`
	IssueType   IssueType   `json:"issuetype"`
	Project     Project     `json:"project"`
	StoryPoints string
	Labels      []string `json:"labels"`
}

// UnmarshalJSON is the implementation for IssueFields to decode from
// JSON.
func (issueFields *IssueFields) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var JSON map[string]interface{}

	if err := json.Unmarshal(data, &JSON); err != nil {
		return err
	}

	issueFields.Created = DateString{JSON["created"].(string)}
	issueFields.Updated = DateString{JSON["updated"].(string)}
	status := JSON["status"].(map[string]interface{})
	issueFields.Status = IssueStatus{
		Name: status["name"].(string),
	}
	issueFields.Summary = JSON["summary"].(string)
	issueType := JSON["issuetype"].(map[string]interface{})
	issueFields.IssueType = IssueType{
		Name: issueType["name"].(string),
	}
	project := JSON["project"].(map[string]interface{})
	issueFields.Project = Project{
		Name: project["name"].(string),
		ID:   project["id"].(string),
		Key:  project["key"].(string),
	}
	if JSON[os.Getenv("JIRA_SP_FIELD")] != nil {
		var sp float64
		var ok bool
		if sp, ok = JSON[os.Getenv("JIRA_SP_FIELD")].(float64); !ok {
			return fmt.Errorf("error converting %v (%v) to float64", os.Getenv("JIRA_SP_FIELD"), JSON[os.Getenv("JIRA_SP_FIELD")])
		}
		issueFields.StoryPoints = strconv.FormatFloat(sp, 'f', -1, 64)
	} else {
		issueFields.StoryPoints = "-"
	}

	for _, l := range JSON["labels"].([]interface{}) {
		issueFields.Labels = append(issueFields.Labels, l.(string))
	}

	return nil
}

// Issue represents a Jira Issue.
type Issue struct {
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}
