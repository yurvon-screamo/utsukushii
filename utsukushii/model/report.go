package model

import "time"

type TestReport struct {
	Title     string        `json:"title"`
	Timestamp time.Time     `json:"timestamp"`
	Duration  time.Duration `json:"duration"`
	Total     int           `json:"total"`
	Coverage  int16         `json:"coverage"`
	Success   int           `json:"success"`
	Dropped   int           `json:"dropped"`
	Skipped   int           `json:"skipped"`
	Tests     []*TestRecord `json:"tests"`
}

type TestRecord struct {
	Name     string        `json:"name"`
	State    TestCaseState `json:"state"`
	Duration time.Duration `json:"duration"`
	Tests    []*TestRecord `json:"tests"`
}

type TestCaseState string

const (
	StateSuccess = "success"
	StateDropped = "dropped"
	StateSkipped = "skipped"
)
