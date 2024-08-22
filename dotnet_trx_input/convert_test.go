package dotnet_trx_input_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yurvon-screamo/utsukushii/dotnet_trx_input"
	"github.com/yurvon-screamo/utsukushii/model"
)

func TestConvert_ReturnsError_WhenXMLIsInvalid(t *testing.T) {
	// arrange
	invalidXML := []byte(`<TestRun><Invalid></TestRun>`)

	// act
	_, err := dotnet_trx_input.Convert(invalidXML)

	// assert
	assert.Error(t, err)
}

func TestConvert_ReturnsTestFileWithCorrectCountsAndRecords(t *testing.T) {
	// arrange
	validXML := []byte(`
		<TestRun>
			<Results>
				<UnitTestResult testName="Test1" outcome="Passed" duration="00:00:01.234" />
				<UnitTestResult testName="Test2" outcome="Failed" duration="00:00:02.345">
					<Output>
						<ErrorInfo>
							<Message>Shouldly.ShouldAssertException</Message>
							<StackTrace>at Saunter</StackTrace>
						</ErrorInfo>
					</Output>
				</UnitTestResult>
				<UnitTestResult testName="Test3" outcome="Skipped" duration="00:00:00.678" />
			</Results>
		</TestRun>
	`)

	expectedSuccess := 1
	expectedDropped := 1
	expectedSkipped := 1
	expectedTotal := 3

	expectedRecords := []*model.TestRecord{
		{
			Name:     "Test1",
			State:    model.StateSuccess,
			Duration: 1234 * time.Millisecond,
			Log:      stringPointer(""),
		},
		{
			Name:     "Test2",
			State:    model.StateDropped,
			Duration: 2345 * time.Millisecond,
			Log:      stringPointer("Shouldly.ShouldAssertException\nat Saunter"),
		},
		{
			Name:     "Test3",
			State:    model.StateSkipped,
			Duration: 678 * time.Millisecond,
			Log:      stringPointer(""),
		},
	}

	// act
	result, err := dotnet_trx_input.Convert(validXML)

	// assert
	assert.NoError(t, err, "unexpected error")

	assert.Equal(t, expectedSuccess, result.Success, "Success count mismatch")
	assert.Equal(t, expectedDropped, result.Dropped, "Dropped count mismatch")
	assert.Equal(t, expectedSkipped, result.Skipped, "Skipped count mismatch")
	assert.Equal(t, expectedTotal, len(result.Records), "Total number of records mismatch")

	for i, record := range result.Records {
		expectedRecord := expectedRecords[i]

		assert.Equal(t, expectedRecord.Name, record.Name, "Record name mismatch")
		assert.Equal(t, expectedRecord.State, record.State, "Record state mismatch")
		assert.Equal(t, expectedRecord.Duration, record.Duration, "Record duration mismatch")

		assert.NotNil(t, record.Log, "Expected log not to be nil")
		assert.Equal(t, *expectedRecord.Log, *record.Log, "Log content mismatch")
	}
}

func stringPointer(s string) *string {
	return &s
}
