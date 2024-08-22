package dotnet_trx_input

import (
	"encoding/xml"
	"strconv"
	"strings"
	"time"

	"github.com/yurvon-screamo/utsukushii/model"
)

func Convert(raw []byte) (*model.TestFile, error) {
	var testRun testRun
	err := xml.Unmarshal(raw, &testRun)
	if err != nil {
		return nil, err
	}

	testFile := &model.TestFile{
		Records: []*model.TestRecord{},
	}

	for _, result := range testRun.Results.UnitTestResults {
		duration := parseDuration(result.Duration)

		var state model.TestCaseState
		switch result.Outcome {
		case "Passed":
			state = model.StateSuccess
			testFile.Success++
		case "Failed":
			state = model.StateDropped
			testFile.Dropped++
		case "Skipped":
			state = model.StateSkipped
			testFile.Skipped++
		default:
			state = model.StateDropped
			testFile.Dropped++
		}

		var log string
		if result.Output.ErrorInfo != nil {
			log = result.Output.ErrorInfo.Message + "\n" + result.Output.ErrorInfo.StackTrace
		}

		testRecord := &model.TestRecord{
			Name:     result.TestName,
			State:    state,
			Duration: duration,
			Log:      &log,
		}

		testFile.Records = append(testFile.Records, testRecord)
	}

	return testFile, nil
}

func parseDuration(duration string) time.Duration {
	if duration == "" {
		return time.Duration(0)
	}

	partsGlobal := strings.Split(duration, ".")
	parts := strings.Split(partsGlobal[0], ":")

	ms, _ := strconv.ParseInt(partsGlobal[1], 10, 64)
	sec, _ := strconv.ParseInt(parts[2], 10, 64)
	min, _ := strconv.ParseInt(parts[1], 10, 64)
	hour, _ := strconv.ParseInt(parts[0], 10, 64)

	return time.Duration(ms)*time.Millisecond +
		time.Duration(sec)*time.Second +
		time.Duration(min)*time.Minute +
		time.Duration(hour)*time.Hour
}
