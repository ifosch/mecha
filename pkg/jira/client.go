package jira

import (
	"context"
	"errors"
	"fmt"
	"encoding/json"
	"net/http"
	"time"
)

// Client represents a basic client object to access the
// API.
type Client struct {
	baseURL string
	userName string
	apiKey string
	HTTPClient *http.Client
	ctx context.Context
}

type errorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// NewClient constructs a new Client object with the API token defined
// in `apiKey` argument.
func NewClient(baseURL, userName, apiKey string, ctx context.Context) *Client {
	if ctx == nil {
		ctx = context.Background()
	}
	return &Client{
		baseURL: baseURL,
		userName: userName,
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
		ctx: ctx, 
	}
}

// GetProjects returns a ProjectList with all the projects found, or
// an error.
func (c *Client) GetProjects()  (pl *ProjectList, err error) {
	pl = &ProjectList{}
	err = c.get("GET", "/rest/api/2/project", pl)
	if err != nil {
		return nil, err
	}

	return
}

// FindProject returns a project with a name matching `projectName`
// argument, or an error.
func (c *Client) FindProject(projectName string) (*Project, error) {
	if projectName == "" {
		return nil, fmt.Errorf("you must specify a project name")
	}

	pl, err := c.GetProjects()
	if err != nil {
		return nil, err
	}

	for _, p := range *pl {
		if p.Name == projectName {
			p.c = c
			return &p, nil
		}
	}
	return nil, fmt.Errorf("project %v not found", projectName)
}

func (c *Client) get(method, url string, response interface{}) error {
	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s%s", c.baseURL, url),
		nil,
	)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.userName,c.apiKey)
	req = req.WithContext(c.ctx)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK ||
	   res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		err = json.NewDecoder(res.Body).Decode(&errRes)
		if err == nil {
			if errRes.Message == "" {
				return fmt.Errorf(
					"unknown error, status code: %d",
					res.StatusCode,
				)
			}

			return errors.New(errRes.Message)
		}
	}

	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		return err
	}

	return nil
}
