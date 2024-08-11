package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReportIsEnriched(t *testing.T) {
	// arrange
	report := &TestReport{
		Tests: []*TestRecord{},
	}
	file := &TestFile{
		Success: 2,
		Dropped: 1,
		Skipped: 1,
		Records: []*TestRecord{
			{
				Name:     "success 1",
				State:    StateSuccess,
				Duration: 0,
			},
			{
				Name:     "success 2",
				State:    StateSuccess,
				Duration: 22 * time.Second,
			},
			{
				Name:     "dropped 1",
				State:    StateDropped,
				Duration: 0 * time.Second,
			},
			{
				Name:     "skipped 1",
				State:    StateSkipped,
				Duration: 11 * time.Second,
			},
		},
	}

	// act
	report.Enrich(file)

	// assert
	assert.Equal(t, file.Success, report.Success, "Success should be equal")
	assert.Equal(t, file.Dropped, report.Dropped, "Dropped should be equal")
	assert.Equal(t, file.Skipped, report.Skipped, "Skipped should be equal")
	assert.Equal(t, len(file.Records), len(report.Tests), "Skipped should be equal")
	assert.Equal(t, 33*time.Second, report.Duration, "Duration should be equal sum of file duration")

	for i, e := range file.Records {
		actual := report.Tests[i]

		assert.Equal(t, e.Name, actual.Name, "Test name should be equal")
		assert.Equal(t, e.State, actual.State, "Test state should be equal")
		assert.Equal(t, e.Duration, actual.Duration, "Test duration should be equal")
	}
}
