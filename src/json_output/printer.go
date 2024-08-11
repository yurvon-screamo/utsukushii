package json_output

import (
	"encoding/json"
	"strconv"
	"time"
	"utsukushii_generator/model"
)

type JsonPrinter struct {
	Report *model.TestReport
}

func (p *JsonPrinter) Print() ([]byte, error) {
	report := p.Report

	o := jsonRoot{
		Title:     report.Title,
		Timestamp: report.Timestamp.Format(time.RFC822),
		Duration:  report.Duration.String(),
		Total:     report.Total,
		Coverage:  report.Coverage,
		Success:   report.Success,
		Dropped:   report.Dropped,
		Skipped:   report.Skipped,
		Tests:     []*jsonTest{},
	}

	o.Tests = mapTestRecords(report.Tests)

	return json.MarshalIndent(o, "", "  ")
}

func mapTestRecords(sourceTests []*model.TestRecord) []*jsonTest {
	jsonTests := []*jsonTest{}

	for _, g := range sourceTests {
		test := &jsonTest{
			Name:     g.Name,
			State:    string(g.State),
			Duration: strconv.FormatInt(g.Duration.Milliseconds(), 10) + "ms",
		}

		if g.Tests != nil && len(g.Tests) > 0 {
			test.Tests = mapTestRecords(g.Tests)
		}

		jsonTests = append(jsonTests, test)
	}

	return jsonTests
}

type jsonRoot struct {
	Title     string      `json:"title"`
	Timestamp string      `json:"timestamp"`
	Duration  string      `json:"duration"`
	Total     int         `json:"total"`
	Coverage  int16       `json:"coverage"`
	Success   int         `json:"success"`
	Dropped   int         `json:"dropped"`
	Skipped   int         `json:"skipped"`
	Tests     []*jsonTest `json:"tests"`
}

type jsonTest struct {
	Name     string      `json:"name"`
	State    string      `json:"state"`
	Duration string      `json:"duration"`
	Tests    []*jsonTest `json:"tests"`
}
