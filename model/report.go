package model

import "time"

type TestReport struct {
	Title     string
	Timestamp time.Time
	Duration  time.Duration
	Total     int
	Coverage  int16
	Success   int
	Dropped   int
	Skipped   int
	Tests     []*TestRecord
}

type TestRecord struct {
	Name     string
	State    TestCaseState
	Duration time.Duration
	Tests    []*TestRecord
	Log      *string
}

type TestCaseState string

const (
	StateSuccess = "success"
	StateDropped = "dropped"
	StateSkipped = "skipped"
)
