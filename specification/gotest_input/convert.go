package gotest_input

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/yurvon-screamo/utsukushii/model"
)

func Convert(raw []byte) (*model.TestFile, error) {
	rawStr := string(raw)
	lines := strings.Split(rawStr, "\n")

	// key - package name, value - group tests
	packageMap := make(map[string]*model.TestRecord)
	// key - test or package name, value - test out
	logMap := make(map[string]string)

	result := &model.TestFile{
		Records: []*model.TestRecord{},
	}

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		test := &goTestOutputRow{}
		if err := json.Unmarshal([]byte(line), test); err != nil {
			return nil, err
		}

		if _, exists := packageMap[test.Package]; !exists {
			record := &model.TestRecord{
				Name:  test.Package,
				Tests: []*model.TestRecord{},
			}
			packageMap[test.Package] = record
		}

		if test.Test == "" {
			if test.Action == "fail" {
				packageMap[test.Package].State = model.StateDropped
			} else if test.Action == "output" {
				logMap[test.Package] += test.Output
			}
			continue
		}

		if test.Action == "output" {
			logMap[test.Test] += test.Output
			continue
		}

		state := getTestState(result, test)
		if state == "" {
			continue
		}

		current := &model.TestRecord{
			Name:     test.Test,
			State:    state,
			Duration: (time.Duration(test.Elapsed * float64(time.Second))),
		}

		group := packageMap[test.Package]
		group.Duration += current.Duration

		if test.Action == "fail" {
			group.State = model.StateDropped
		} else if group.State != model.StateDropped {
			if test.Action == "skip" && len(group.Tests) == 0 {
				group.State = model.StateSkipped
			} else if test.Action == "pass" {
				group.State = model.StateSuccess
			}
		}

		group.Tests = append(group.Tests, current)
	}

	for _, group := range packageMap {
		if group.State == "" {
			continue
		}

		if group.State != model.StateSuccess && len(group.Tests) == 0 {
			if l, v := logMap[group.Name]; v {
				l = strings.TrimRight(l, "\n")
				group.Log = &l
				result.Dropped++
			}
		}

		for _, test := range group.Tests {
			if test.State != model.StateSuccess {
				if l, v := logMap[test.Name]; v {
					l = strings.TrimRight(l, "\n")
					test.Log = &l
				}
			}
		}

		result.Records = append(result.Records, group)
	}

	return result, nil
}

func getTestState(result *model.TestFile, test *goTestOutputRow) model.TestCaseState {
	if test.Action == "fail" {
		result.Dropped++
		return model.StateDropped
	}

	if test.Action == "skip" {
		result.Skipped++
		return model.StateSkipped
	}

	if test.Action == "pass" {
		result.Success++
		return model.StateSuccess
	}

	return ""
}
