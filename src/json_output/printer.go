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
		Groups:    []*jsonGroup{},
	}

	for _, g := range report.Groups {
		group := &jsonGroup{
			Name:     g.Name,
			State:    string(g.State),
			Duration: strconv.FormatInt(g.Duration.Milliseconds(), 10) + "ms",
			Tests:    []*jsonTest{},
		}

		for _, t := range g.Tests {
			group.Tests = append(group.Tests, &jsonTest{
				Name:     t.Name,
				State:    string(t.State),
				Duration: strconv.FormatInt(t.Duration.Milliseconds(), 10) + "ms",
			})
		}

		o.Groups = append(o.Groups, group)
	}

	return json.MarshalIndent(o, "", "  ")
}

type jsonRoot struct {
	Title     string       `json:"title"`
	Timestamp string       `json:"timestamp"`
	Duration  string       `json:"duration"`
	Total     int          `json:"total"`
	Coverage  int16        `json:"coverage"`
	Success   int          `json:"success"`
	Dropped   int          `json:"dropped"`
	Skipped   int          `json:"skipped"`
	Groups    []*jsonGroup `json:"groups"`
}

type jsonGroup struct {
	Name     string      `json:"name"`
	State    string      `json:"state"`
	Duration string      `json:"duration"`
	Tests    []*jsonTest `json:"tests"`
}

type jsonTest struct {
	Name     string `json:"name"`
	State    string `json:"state"`
	Duration string `json:"duration"`
}
