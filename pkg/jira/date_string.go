package jira

import (
	"time"
)

// DateString represents date based fields in Jira.
type DateString struct {
	string
}

// String implements Stringer interface.
func (ds DateString) String() string {
	inLayout := "2006-01-02T15:04:05.999999999"
	time, _ := time.Parse(inLayout, ds.string[:len(ds.string)-5])
	return time.Format("2006-01-02")
}
