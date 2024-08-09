package main

import (
	"encoding/json"
	"time"
	"utsukushii_generator/model"
)

func convertToJson(report *model.TestReport) ([]byte, error) {
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
			Duration: g.Duration.String(),
			Tests:    []*jsonTest{},
		}

		for _, t := range g.Tests {
			group.Tests = append(group.Tests, &jsonTest{
				Name:     t.Name,
				State:    string(t.State),
				Duration: t.Duration.String(),
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
	Coverage  int          `json:"coverage"`
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
