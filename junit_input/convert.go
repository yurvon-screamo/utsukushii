package junit_input

import (
	"encoding/xml"
	"strconv"
	"time"

	"github.com/yurvon-screamo/utsukushii/model"
)

func Convert(raw []byte) (*model.TestFile, error) {
	var junitReport jUnitTestSuites
	err := xml.Unmarshal(raw, &junitReport)
	if err != nil {
		return nil, err
	}

	tests := convertJUnitToTestReport(junitReport)
	return tests, nil
}

func convertJUnitToTestReport(report jUnitTestSuites) *model.TestFile {
	result := &model.TestFile{
		Success: 0,
		Dropped: 0,
		Skipped: 0,
		Records: []*model.TestRecord{},
	}

	for _, suite := range report.TestSuites {
		if suite.Tests == 0 {
			continue
		}

		durationSuiteRaw, _ := strconv.ParseFloat(suite.Time, 64)
		durationSuite := time.Duration(durationSuiteRaw * float64(time.Second))

		group := model.TestRecord{
			Name:     suite.Name,
			State:    determineState(suite),
			Duration: durationSuite,
			Tests:    []*model.TestRecord{},
		}

		for _, tc := range suite.TestCases {
			durationCaseRaw, _ := strconv.ParseFloat(tc.Time, 64)
			test := model.TestRecord{
				Name:     tc.Name,
				State:    determineTestState(tc),
				Duration: time.Duration(durationCaseRaw * float64(time.Second)),
			}
			group.Tests = append(group.Tests, &test)
		}

		result.Records = append(result.Records, &group)
	}

	return result
}

func determineState(testSuite jUnitTestSuite) model.TestCaseState {
	totalErr := testSuite.Errors + testSuite.Failures
	if totalErr > 0 {
		return model.StateDropped
	}

	if testSuite.Skipped == testSuite.Tests {
		return model.StateSkipped
	}

	return model.StateSuccess
}

func determineTestState(tc jUnitTestCase) model.TestCaseState {
	if tc.Failure != nil {
		return model.StateDropped
	}

	if tc.Skipped != nil {
		return model.StateSkipped
	}

	return model.StateSuccess
}
