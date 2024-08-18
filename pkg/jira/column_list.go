package jira

import (
	"fmt"
)

// ColumnList is a list of Columns.
type ColumnList []Column

// GetLabelNames returns a list of all the label names for the
// columns.
func (cl ColumnList) GetLabelNames() []string {
	labelNames := []string{}

	for _, c := range cl {
		labelNames = append(labelNames, c.Value)
	}

	return labelNames
}

// GetLabels returns a list of all the labels for the columns.
func (cl ColumnList) GetLabels() []string {
	labels := []string{}

	for _, c := range cl {
		labels = append(labels, c.Label)
	}

	return labels
}

// String is the Stringer implementation.
func (cl ColumnList) String() string {
	output := ""

	for _, c := range cl {
		output += fmt.Sprintf("%v\t", c.Label)
	}
	output += "\n"

	return output
}
