package junit_input

import "encoding/xml"

type jUnitTestSuites struct {
	XMLName    xml.Name         `xml:"testsuites"`
	TestSuites []jUnitTestSuite `xml:"testsuite"`
}

type jUnitTestSuite struct {
	XMLName   xml.Name        `xml:"testsuite"`
	Name      string          `xml:"name,attr"`
	Tests     int             `xml:"tests,attr"`
	Failures  int             `xml:"failures,attr"`
	Errors    int             `xml:"errors,attr"`
	Skipped   int             `xml:"skipped,attr"`
	Time      string          `xml:"time,attr"`
	TestCases []jUnitTestCase `xml:"testcase"`
}

type jUnitTestCase struct {
	XMLName   xml.Name      `xml:"testcase"`
	ClassName string        `xml:"classname,attr"`
	Name      string        `xml:"name,attr"`
	Time      string        `xml:"time,attr"`
	Failure   *jUnitFailure `xml:"failure,omitempty"`
	Error     *jUnitError   `xml:"error,omitempty"`
	Skipped   *jUnitSkipped `xml:"skipped,omitempty"`
}

type jUnitFailure struct {
	XMLName xml.Name `xml:"failure"`
	Message string   `xml:"message,attr"`
	Type    string   `xml:"type,attr"`
	Text    string   `xml:",chardata"`
}

type jUnitError struct {
	XMLName xml.Name `xml:"error"`
	Message string   `xml:"message,attr"`
	Type    string   `xml:"type,attr"`
	Text    string   `xml:",chardata"`
}

type jUnitSkipped struct {
	XMLName xml.Name `xml:"skipped"`
	Message string   `xml:"message,attr,omitempty"`
}
