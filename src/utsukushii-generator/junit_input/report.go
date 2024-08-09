package junit_input

import (
	"encoding/xml"
	"strconv"
	"time"
	"utsukushii_generator/model"
)

type JUnitReader struct {
	Raw []byte
}

func (r *JUnitReader) Convert() (*model.TestReport, error) {
	var junitReport JUnitTestSuites
	err := xml.Unmarshal(r.Raw, &junitReport)
	if err != nil {
		return nil, err
	}

	report := convertJUnitToTestReport(junitReport)

	return report, nil
}

func convertJUnitToTestReport(report JUnitTestSuites) *model.TestReport {
	totalTests := 0
	totalFailures := 0
	totalSkipped := 0

	var totalDuration time.Duration
	totalDuration = 0

	var groups []model.TestGroup

	for _, suite := range report.TestSuites {
		if suite.Tests == 0 {
			continue
		}

		totalTests += suite.Tests
		totalFailures += suite.Failures
		totalSkipped += suite.Skipped

		durationSuiteRaw, _ := strconv.ParseFloat(suite.Time, 64)
		durationSuite := time.Duration(durationSuiteRaw * float64(time.Second))
		totalDuration += durationSuite

		group := model.TestGroup{
			Name:     suite.Name,
			State:    determineState(suite),
			Duration: durationSuite,
			Tests:    []model.Test{},
		}

		for _, tc := range suite.TestCases {
			durationCaseRaw, _ := strconv.ParseFloat(tc.Time, 64)
			test := model.Test{
				Name:     tc.Name,
				State:    determineTestState(tc),
				Duration: time.Duration(durationCaseRaw * float64(time.Second)),
			}
			group.Tests = append(group.Tests, test)
		}

		groups = append(groups, group)
	}

	return &model.TestReport{
		Title:     "Tester",
		Timestamp: time.Now().UTC(),
		Duration:  totalDuration,
		Total:     totalTests,
		Coverage:  0, // TODO
		Success:   totalTests - totalFailures - totalSkipped,
		Dropped:   totalFailures,
		Skipped:   totalSkipped,
		Groups:    groups,
	}
}

func determineState(testSuite JUnitTestSuite) model.TestCaseState {
	totalErr := testSuite.Errors + testSuite.Failures
	if totalErr > 0 {
		return model.StateDropped
	}

	if testSuite.Skipped == testSuite.Tests {
		return model.StateSkipped
	}

	return model.StateSuccess
}

func determineTestState(tc JUnitTestCase) model.TestCaseState {
	if tc.Failure != nil {
		return model.StateDropped
	}

	if tc.Skipped != nil {
		return model.StateSkipped
	}

	return model.StateSuccess
}
