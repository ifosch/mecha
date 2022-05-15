package jira

import (
	"encoding/json"
	"fmt"
)

// Project represents a Jira Project.
type Project struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Key  string `json:"key"`
	c    *Client
}

// CreateSprint creates a new sprint in the first board and returns its name, or returns an error.
func (p *Project) CreateSprint() (string, error) {
	lastCreatedSprint, err := p.GetLastCreatedSprint()
	if err != nil {
		return "", err
	}

	newName, err := lastCreatedSprint.NextSprintName()
	if err != nil {
		return "", nil
	}

	createSprintInputData := &createSprintInput{
		Name:    newName,
		BoardID: lastCreatedSprint.BoardID,
	}
	createSprintInputJSON, err := json.Marshal(createSprintInputData)
	if err != nil {
		return "", err
	}

	_, err = p.c.post(
		"/rest/agile/1.0/sprint",
		createSprintInputJSON,
	)
	if err != nil {
		return "", err
	}

	return newName, nil
}

type createSprintInput struct {
	Name    string `json:"name"`
	BoardID int    `json:"originBoardId"`
}

// GetSprints returns all the sprints found for the Project, or an error.
func (p *Project) GetSprints(state string) (*SprintList, error) {
	if state == "" {
		state = "active,future"
	}

	var boards BoardList
	err := p.c.getInterface(fmt.Sprintf("/rest/agile/1.0/board?projectKeyOrId=%v", p.ID), &boards)
	if err != nil {
		return nil, err
	}

	var sprints, finalSprints SprintList
	finalSprints = SprintList{
		Values: []Sprint{},
	}
	for _, b := range boards.Values {
		err = p.c.getInterface(fmt.Sprintf("/rest/agile/1.0/board/%v/sprint?state=%v", b.ID, state), &sprints)
		if err != nil && err.Error() != "unknown error, status code: 400" {
			return nil, err
		}
		for _, s := range sprints.Values {
			if s.BoardID == b.ID {
				s.c = p.c
				s.p = p
				finalSprints.Values = append(finalSprints.Values, s)
			}
		}
	}

	return &finalSprints, nil
}

// GetLastCreatedSprint return last created Sprint, or an error.
func (p *Project) GetLastCreatedSprint() (*Sprint, error) {
	sprints, err := p.GetSprints("closed,active,future")
	if err != nil {
		return nil, err
	}
	return &sprints.Values[len(sprints.Values)-1], nil
}

// GetCurrentSprint return current active Sprint, or an error.
func (p *Project) GetCurrentSprint() (*Sprint, error) {
	activeSprints, err := p.GetSprints("active")
	if err != nil {
		return nil, err
	}
	return &activeSprints.Values[0], nil
}

// GetNextSprint returns next Sprint, or an error.
func (p *Project) GetNextSprint() (*Sprint, error) {
	futureSprints, err := p.GetSprints("future")
	if err != nil {
		return nil, err
	}
	return &futureSprints.Values[0], nil
}
