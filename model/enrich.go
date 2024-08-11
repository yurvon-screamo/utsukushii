package model

type TestFile struct {
	Success int
	Dropped int
	Skipped int
	Records []*TestRecord
}

func (r *TestReport) Enrich(file *TestFile) {
	r.Success += file.Success
	r.Dropped += file.Dropped
	r.Skipped += file.Skipped
	r.Total += (file.Dropped + file.Skipped + file.Success)

	for i := 0; i < len(file.Records); i++ {
		testRecord := file.Records[i]
		r.Tests = append(r.Tests, testRecord)

		r.Duration += testRecord.Duration
	}
}
