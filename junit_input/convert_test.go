package junit_input_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yurvon-screamo/utsukushii/junit_input"
	"github.com/yurvon-screamo/utsukushii/model"
)

func TestInvalidJunitReportIsReturnError(t *testing.T) {
	// arrange
	const junitInput = "helpme"

	// act
	tests, err := junit_input.Convert([]byte(junitInput))

	// assert
	assert.Nil(t, tests, "Tests must be nil on invalid junit content")
	assert.NotNil(t, err, "Error must be not nil on invalid junit content")
}

func TestValidJunitReportIsConverted(t *testing.T) {
	// arrange
	const junitInput = `<?xml version="1.0" encoding="UTF-8"?>
<testsuites tests="4" failures="1" skipped="1">
	<testsuite name="github.com/yurvon-screamo/utsukushii/model" tests="1" failures="0" errors="0" id="0" hostname="DESKTOP-FPT2UJE" time="1.000" timestamp="2024-08-11T17:57:20+03:00">
		<testcase name="TestReportIsEnriched" classname="github.com/yurvon-screamo/utsukushii/model" time="1.00"></testcase>
	</testsuite>
	<testsuite name="github.com/yurvon-screamo/utsukushii" tests="0" failures="0" errors="0" id="1" hostname="DESKTOP-FPT2UJE" time="0.000" timestamp="2024-08-11T17:57:20+03:00"></testsuite>
	<testsuite name="github.com/yurvon-screamo/utsukushii/cmd" tests="0" failures="0" errors="0" id="2" hostname="DESKTOP-FPT2UJE" time="0.000" timestamp="2024-08-11T17:57:20+03:00"></testsuite>
	<testsuite name="github.com/yurvon-screamo/utsukushii/json_output" tests="0" failures="0" errors="0" id="3" hostname="DESKTOP-FPT2UJE" time="0.000" timestamp="2024-08-11T17:57:20+03:00"></testsuite>
	<testsuite name="github.com/yurvon-screamo/utsukushii/utsukushii_ui" tests="0" failures="0" errors="0" id="4" hostname="DESKTOP-FPT2UJE" time="0.000" timestamp="2024-08-11T17:57:21+03:00"></testsuite>
	<testsuite name="github.com/yurvon-screamo/utsukushii/junit_input" tests="3" failures="1" errors="0" id="5" hostname="DESKTOP-FPT2UJE" skipped="1" time="0.391" timestamp="2024-08-11T17:57:21+03:00">
		<testcase name="TestSkip" classname="github.com/yurvon-screamo/utsukushii/junit_input" time="0.000">
			<skipped message="Skipped"></skipped>
		</testcase>
		<testcase name="TestInvalidJunitReportIsReturnError" classname="github.com/yurvon-screamo/utsukushii/junit_input" time="0.000"></testcase>
		<testcase name="TestValidJunitReportIsConverted" classname="github.com/yurvon-screamo/utsukushii/junit_input" time="1.000">
			<failure message="Failed"><![CDATA[    convert_test.go:53: 
        	Error Trace:	D:/github/utsukushii/junit_input/convert_test.go:53
        	Error:      	Not equal: 
        	            	expected: 10
        	            	actual  : 1
        	Test:       	TestValidJunitReportIsConverted
        	Messages:   	Skipped must be 1]]></failure>
		</testcase>
	</testsuite>
</testsuites>`

	// act
	tests, err := junit_input.Convert([]byte(junitInput))

	// assert
	assert.Nil(t, err, "Error must be nil on valid junit content")
	assert.NotNil(t, tests, "Tests must be not nil on valid junit content")

	assert.Equal(t, 2, tests.Success, "Success must be 2")
	assert.Equal(t, 1, tests.Dropped, "Dropped must be 1")
	assert.Equal(t, 1, tests.Skipped, "Skipped must be 1")

	group1 := tests.Records[0]
	assertJUnitGroup(t, group1, "github.com/yurvon-screamo/utsukushii/model", 1*time.Second, model.StateSuccess, 1)
	assertJUnitTest(t, group1.Tests[0], "TestReportIsEnriched", 1*time.Second, model.StateSuccess, nil)

	group2 := tests.Records[1]
	log := "Skipped"
	assertJUnitGroup(t, group2, "github.com/yurvon-screamo/utsukushii/junit_input", 391*time.Millisecond, model.StateDropped, 3)
	assertJUnitTest(t, group2.Tests[0], "TestSkip", 0*time.Second, model.StateSkipped, &log)
	assertJUnitTest(t, group2.Tests[1], "TestInvalidJunitReportIsReturnError", 0*time.Second, model.StateSuccess, nil)
	log = `convert_test.go:53: 
        	Error Trace:	D:/github/utsukushii/junit_input/convert_test.go:53
        	Error:      	Not equal: 
        	            	expected: 10
        	            	actual  : 1
        	Test:       	TestValidJunitReportIsConverted
        	Messages:   	Skipped must be 1`
	assertJUnitTest(t, group2.Tests[2], "TestValidJunitReportIsConverted", 1*time.Second, model.StateDropped, &log)
}

func assertJUnitGroup(
	t *testing.T,
	group *model.TestRecord,
	name string,
	duration time.Duration,
	testCase model.TestCaseState,
	len int) {

	assert.Equal(t, name, group.Name)
	assert.Equal(t, duration, group.Duration)
	assert.Equal(t, testCase, group.State)
	assert.NotNil(t, group.Tests)
	assert.Len(t, group.Tests, len)
}

func assertJUnitTest(
	t *testing.T,
	test *model.TestRecord,
	name string,
	duration time.Duration,
	testCase model.TestCaseState,
	log *string) {

	assert.Equal(t, name, test.Name)
	assert.Equal(t, duration, test.Duration)
	assert.Equal(t, testCase, test.State)
	assert.Equal(t, log, test.Log)
	assert.Nil(t, test.Tests)
}
