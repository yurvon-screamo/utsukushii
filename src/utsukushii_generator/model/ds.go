package model

import "time"

type TestReport struct {
	Title     string        `json:"title"`
	Timestamp time.Time     `json:"timestamp"`
	Duration  time.Duration `json:"duration"`
	Total     int           `json:"total"`
	Coverage  int           `json:"coverage"`
	Success   int           `json:"success"`
	Dropped   int           `json:"dropped"`
	Skipped   int           `json:"skipped"`
	Groups    []TestGroup   `json:"groups"`
}

type TestGroup struct {
	Name     string        `json:"name"`
	State    TestCaseState `json:"state"`
	Duration time.Duration `json:"duration"`
	Tests    []Test        `json:"tests"`
}

type Test struct {
	Name     string        `json:"name"`
	State    TestCaseState `json:"state"`
	Duration time.Duration `json:"duration"`
}

type TestCaseState string

const (
	StateSuccess = "success"
	StateDropped = "dropped"
	StateSkipped = "skipped"
)
