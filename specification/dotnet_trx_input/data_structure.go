package dotnet_trx_input

import "encoding/xml"

type testRun struct {
	XMLName xml.Name `xml:"TestRun"`
	Results results  `xml:"Results"`
}

type results struct {
	UnitTestResults []unitTestResult `xml:"UnitTestResult"`
}

type unitTestResult struct {
	TestName string `xml:"testName,attr"`
	Outcome  string `xml:"outcome,attr"`
	Duration string `xml:"duration,attr"`
	Output   output `xml:"Output"`
}

type output struct {
	ErrorInfo *errorInfo `xml:"ErrorInfo"`
}

type errorInfo struct {
	Message    string `xml:"Message"`
	StackTrace string `xml:"StackTrace"`
}
