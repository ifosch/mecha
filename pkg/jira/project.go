package jira

import (
	"fmt"
)

// Project represents a Jira Project.
type Project struct{
	Name string `json:"name"`
	ID string `json:"id"`
	Key string `json:"key"`
	c *Client
}

// GetSprints returns all the sprints found for the Project, or an error.
func (p *Project) GetSprints(state string) (*SprintList, error) {
	if state == "" {
		state="active,future"
	}

	var boards BoardList
	err := p.c.getInterface("GET", fmt.Sprintf("/rest/agile/1.0/board?projectKeyOrId=%v", p.ID), &boards)
	if err != nil {
		return nil, err
	}

	var sprints, finalSprints SprintList
	finalSprints = SprintList{
		Values: []Sprint{},
	}
	for _, b := range boards.Values {
		err = p.c.getInterface("GET", fmt.Sprintf("/rest/agile/1.0/board/%v/sprint?state=%v", b.ID, state), &sprints)
		if err != nil && err.Error() != "unknown error, status code: 400" {
			return nil, err
		}
		for _, s := range sprints.Values {
			if s.BoardID == b.ID {
				s.c = p.c
				finalSprints.Values = append(finalSprints.Values, s)
			}
		}
	}

	return &finalSprints, nil
}
