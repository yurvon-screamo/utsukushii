package gotest_input

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yurvon-screamo/utsukushii/model"
)

func TestInvalidGoJsonTestReportIsReturnError(t *testing.T) {
	// arrange
	const input = "helpme"

	// act
	tests, err := Convert([]byte(input))

	// assert
	assert.Nil(t, tests, "Tests must be nil on invalid go test content")
	assert.NotNil(t, err, "Error must be not nil on invalid go test content")
}

func TestValidGoJsonTestReportIsConverted(t *testing.T) {
	// arrange
	const junitInput = `{"Time":"2024-08-11T23:19:22.5735729+03:00","Action":"start","Package":"github.com/yurvon-screamo/utsukushii"}
{"Time":"2024-08-11T23:19:22.5735729+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii","Output":"FAIL\tgithub.com/yurvon-screamo/utsukushii [build failed]\n"}
{"Time":"2024-08-11T23:19:22.5735729+03:00","Action":"fail","Package":"github.com/yurvon-screamo/utsukushii","Elapsed":0}
{"Time":"2024-08-11T23:19:22.5735729+03:00","Action":"start","Package":"github.com/yurvon-screamo/utsukushii/cmd"}
{"Time":"2024-08-11T23:19:22.5735729+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/cmd","Output":"FAIL\tgithub.com/yurvon-screamo/utsukushii/cmd [build failed]\n"}
{"Time":"2024-08-11T23:19:22.5735729+03:00","Action":"fail","Package":"github.com/yurvon-screamo/utsukushii/cmd","Elapsed":0}
{"Time":"2024-08-11T23:19:22.5735729+03:00","Action":"start","Package":"github.com/yurvon-screamo/utsukushii/gotest_input"}
{"Time":"2024-08-11T23:19:22.5735729+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/gotest_input","Output":"FAIL\tgithub.com/yurvon-screamo/utsukushii/gotest_input [build failed]\n"}
{"Time":"2024-08-11T23:19:22.5735729+03:00","Action":"fail","Package":"github.com/yurvon-screamo/utsukushii/gotest_input","Elapsed":0}
{"Time":"2024-08-11T23:19:22.5740932+03:00","Action":"start","Package":"github.com/yurvon-screamo/utsukushii/json_output"}
{"Time":"2024-08-11T23:19:22.5740932+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/json_output","Output":"?   \tgithub.com/yurvon-screamo/utsukushii/json_output\t[no test files]\n"}
{"Time":"2024-08-11T23:19:22.5740932+03:00","Action":"skip","Package":"github.com/yurvon-screamo/utsukushii/json_output","Elapsed":0}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"start","Package":"github.com/yurvon-screamo/utsukushii/junit_input"}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"start","Package":"github.com/yurvon-screamo/utsukushii/model"}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"run","Package":"github.com/yurvon-screamo/utsukushii/model","Test":"TestReportIsEnriched"}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/model","Test":"TestReportIsEnriched","Output":"=== RUN   TestReportIsEnriched\n"}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"start","Package":"github.com/yurvon-screamo/utsukushii/utsukushii_ui"}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/model","Test":"TestReportIsEnriched","Output":"--- PASS: TestReportIsEnriched (0.00s)\n"}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"pass","Package":"github.com/yurvon-screamo/utsukushii/model","Test":"TestReportIsEnriched","Elapsed":0}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/utsukushii_ui","Output":"?   \tgithub.com/yurvon-screamo/utsukushii/utsukushii_ui\t[no test files]\n"}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/model","Output":"PASS\n"}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/model","Output":"ok  \tgithub.com/yurvon-screamo/utsukushii/model\t(cached)\n"}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"pass","Package":"github.com/yurvon-screamo/utsukushii/model","Elapsed":0}
{"Time":"2024-08-11T23:19:22.9764591+03:00","Action":"skip","Package":"github.com/yurvon-screamo/utsukushii/utsukushii_ui","Elapsed":0}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"run","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestInvalidJunitReportIsReturnError"}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestInvalidJunitReportIsReturnError","Output":"=== RUN   TestInvalidJunitReportIsReturnError\n"}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestInvalidJunitReportIsReturnError","Output":"    convert_test.go:19: \n"}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestInvalidJunitReportIsReturnError","Output":"        \tError Trace:\tD:/github/utsukushii/junit_input/convert_test.go:19\n"}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestInvalidJunitReportIsReturnError","Output":"        \tError:      \tExpected value not to be nil.\n"}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestInvalidJunitReportIsReturnError","Output":"        \tTest:       \tTestInvalidJunitReportIsReturnError\n"}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestInvalidJunitReportIsReturnError","Output":"        \tMessages:   \tTests must be nil on invalid junit content\n"}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestInvalidJunitReportIsReturnError","Output":"--- FAIL: TestInvalidJunitReportIsReturnError (0.00s)\n"}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"fail","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestInvalidJunitReportIsReturnError","Elapsed":0}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"run","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestValidJunitReportIsConverted"}
{"Time":"2024-08-11T23:19:23.2440136+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestValidJunitReportIsConverted","Output":"=== RUN   TestValidJunitReportIsConverted\n"}
{"Time":"2024-08-11T23:19:23.2445449+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestValidJunitReportIsConverted","Output":"--- PASS: TestValidJunitReportIsConverted (0.00s)\n"}
{"Time":"2024-08-11T23:19:23.2445449+03:00","Action":"pass","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Test":"TestValidJunitReportIsConverted","Elapsed":0}
{"Time":"2024-08-11T23:19:23.2445449+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Output":"FAIL\n"}
{"Time":"2024-08-11T23:19:23.2458028+03:00","Action":"output","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Output":"FAIL\tgithub.com/yurvon-screamo/utsukushii/junit_input\t0.269s\n"}
{"Time":"2024-08-11T23:19:23.2458028+03:00","Action":"fail","Package":"github.com/yurvon-screamo/utsukushii/junit_input","Elapsed":0.269}`

	// act
	tests, err := Convert([]byte(junitInput))

	// assert
	assert.Nil(t, err, "Error must be nil on valid go test content")
	assert.NotNil(t, tests, "Tests must be not nil on valid go test content")

	assert.Equal(t, 2, tests.Success, "Success must be 2")
	assert.Equal(t, 5, tests.Dropped, "Dropped must be 5")
	assert.Equal(t, 0, tests.Skipped, "Skipped must be 0")

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
