package gotest_input_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yurvon-screamo/utsukushii/gotest_input"
	"github.com/yurvon-screamo/utsukushii/model"
)

func TestConvert_WhenAllTestsPass_ShouldReturnSuccess(t *testing.T) {
	// arrange
	testOutput := []byte(`
	{"Time":"2023-08-22T10:00:00Z","Action":"pass","Package":"mypackage","Test":"TestExample","Elapsed":0.005}
	{"Time":"2023-08-22T10:00:00Z","Action":"pass","Package":"mypackage","Elapsed":0.005}
	`)

	expected := &model.TestFile{
		Records: []*model.TestRecord{
			{
				Name:     "mypackage",
				State:    model.StateSuccess,
				Duration: 5 * time.Millisecond,
				Tests: []*model.TestRecord{
					{
						Name:     "TestExample",
						State:    model.StateSuccess,
						Duration: 5 * time.Millisecond,
					},
				},
			},
		},
		Success: 1,
	}

	// act
	actual, err := gotest_input.Convert(testOutput)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, expected)
	assert.Equal(t, expected.Success, actual.Success)
	assert.Equal(t, expected.Dropped, actual.Dropped)
	assert.Equal(t, expected.Skipped, actual.Skipped)

	assertTestRecords(t, expected.Records, actual.Records)
}

func TestConvert_WhenTestFails_ShouldReturnDropped(t *testing.T) {
	// arrange
	testOutput := []byte(`
	{"Time":"2023-08-22T10:00:00Z","action":"fail","Package":"mypackage","Test":"TestExample","Elapsed":0.005}
	{"Time":"2023-08-22T10:00:00Z","action":"fail","Package":"mypackage","Elapsed":0.005}
	`)

	expected := &model.TestFile{
		Records: []*model.TestRecord{
			{
				Name:     "mypackage",
				State:    model.StateDropped,
				Duration: 5 * time.Millisecond,
				Tests: []*model.TestRecord{
					{
						Name:     "TestExample",
						State:    model.StateDropped,
						Duration: 5 * time.Millisecond,
					},
				},
			},
		},
		Dropped: 1,
	}

	// act
	actual, err := gotest_input.Convert(testOutput)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected.Success, actual.Success)
	assert.Equal(t, expected.Dropped, actual.Dropped)
	assert.Equal(t, expected.Skipped, actual.Skipped)

	assertTestRecords(t, expected.Records, actual.Records)
}

func TestConvert_WhenTestIsSkipped_ShouldReturnSkipped(t *testing.T) {
	// arrange
	testOutput := []byte(`
	{"Time":"2023-08-22T10:00:00Z","Action":"skip","Package":"mypackage","Test":"TestExample","Elapsed":0.000}
	{"Time":"2023-08-22T10:00:00Z","Action":"skip","Package":"mypackage","Elapsed":0.000}
	`)

	expected := &model.TestFile{
		Records: []*model.TestRecord{
			{
				Name:  "mypackage",
				State: model.StateSkipped,
				Tests: []*model.TestRecord{
					{
						Name:     "TestExample",
						State:    model.StateSkipped,
						Duration: 0,
					},
				},
			},
		},
		Skipped: 1,
	}

	// act
	actual, err := gotest_input.Convert(testOutput)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected.Success, actual.Success)
	assert.Equal(t, expected.Dropped, actual.Dropped)
	assert.Equal(t, expected.Skipped, actual.Skipped)

	assertTestRecords(t, expected.Records, actual.Records)
}

func TestConvert_WhenMixedResults_ShouldReturnCorrectStates(t *testing.T) {
	// arrange
	testOutput := []byte(`
	{"Time":"2023-08-22T10:00:00Z","Action":"pass","Package":"mypackage","Test":"TestSuccess","Elapsed":0.003}
	{"Time":"2023-08-22T10:00:00Z","Action":"fail","Package":"mypackage","Test":"TestFailure","Elapsed":0.005}
	{"Time":"2023-08-22T10:00:00Z","Action":"skip","Package":"mypackage","Test":"TestSkip","Elapsed":0.000}
	{"Time":"2023-08-22T10:00:00Z","Action":"fail","Package":"mypackage","Elapsed":0.008}
	`)

	expected := &model.TestFile{
		Records: []*model.TestRecord{
			{
				Name:     "mypackage",
				State:    model.StateDropped,
				Duration: 8 * time.Millisecond,
				Tests: []*model.TestRecord{
					{
						Name:     "TestSuccess",
						State:    model.StateSuccess,
						Duration: 3 * time.Millisecond,
					},
					{
						Name:     "TestFailure",
						State:    model.StateDropped,
						Duration: 5 * time.Millisecond,
					},
					{
						Name:     "TestSkip",
						State:    model.StateSkipped,
						Duration: 0,
					},
				},
			},
		},
		Success: 1,
		Dropped: 1,
		Skipped: 1,
	}

	// act
	actual, err := gotest_input.Convert(testOutput)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, expected.Success, actual.Success)
	assert.Equal(t, expected.Dropped, actual.Dropped)
	assert.Equal(t, expected.Skipped, actual.Skipped)

	assertTestRecords(t, expected.Records, actual.Records)
}

func TestConvert_WhenInvalidJSON_ShouldReturnError(t *testing.T) {
	// arrange
	testOutput := []byte(`invalid json`)

	// act
	actual, err := gotest_input.Convert(testOutput)

	// assert
	assert.Error(t, err)
	assert.Nil(t, actual)
}

// assertTestRecords compares two slices of TestRecord for equality by checking each field in each record.
func assertTestRecords(t *testing.T, expected []*model.TestRecord, actual []*model.TestRecord) {
	assert.Equal(t, len(expected), len(actual), "Number of records does not match")

	for i := range expected {
		// Assert the basic properties of the TestRecord
		assert.Equal(t, expected[i].Name, actual[i].Name, "Record name does not match at index %d", i)
		assert.Equal(t, expected[i].State, actual[i].State, "Record state does not match at index %d", i)
		assert.Equal(t, expected[i].Duration, actual[i].Duration, "Record duration does not match at index %d", i)

		if expected[i].Log != nil && actual[i].Log != nil {
			assert.Equal(t, *expected[i].Log, *actual[i].Log, "Record log does not match at index %d", i)
		} else {
			assert.Equal(t, expected[i].Log, actual[i].Log, "Record log does not match at index %d", i)
		}

		// Recursively assert the nested TestRecords
		assertTestRecords(t, expected[i].Tests, actual[i].Tests)
	}
}
