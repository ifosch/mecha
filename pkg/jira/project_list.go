package jira

import (
	"fmt"
)

type ProjectList []struct{
	Name string `json:"name"`
}

func (pl ProjectList) String() string {
	output := ""

	for _, p := range pl {
		output += fmt.Sprintf("- %v\n", p.Name)
	}

	return output
}
