package main

import (
	"fmt"
	"io/ioutil"
	"utsukushii_generator/cmd"
	"utsukushii_generator/junit_input"
	"utsukushii_generator/model"
)

type InputConverter interface {
	Convert() model.TestReport
}

func main() {
	cmd.Execute()

	junitData, err := ioutil.ReadFile("junit.xml")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	junitReader := junit_input.JUnitReader{
		Raw: junitData,
	}

	report, err := junitReader.Convert()
	if err != nil {
		fmt.Println("Error convert junit:", err)
		return
	}

	jsonData, err := convertToJson(report)
	if err != nil {
		fmt.Println("Error convert json:", err)
		return
	}

	err = ioutil.WriteFile("../utsukushii_ui/utsukushii.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("Test report written to utsukushii.json")
}
