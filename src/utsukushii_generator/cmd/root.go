package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"utsukushii_generator/json_output"
	"utsukushii_generator/junit_input"

	"github.com/spf13/cobra"
)

func handle(cmd *cobra.Command, args []string) {
	junitData, err := ioutil.ReadFile(junit)
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

	report.Title = title
	report.Coverage = coverage

	jsonPrinter := json_output.JsonPrinter{
		Report: report,
	}

	jsonData, err := jsonPrinter.Print()
	if err != nil {
		fmt.Println("Error convert json:", err)
		return
	}

	err = ioutil.WriteFile(output, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Printf("Test report written to %s\n", output)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var title string
var junit string
var coverage int16
var output string

var rootCmd = &cobra.Command{
	Use:     "utsukushii_generator",
	Short:   "utsukushii test-gen",
	Long:    `utsukushii is a utility for converting your reports to utsukushii format.`,
	Example: "\nutsukushii_generator --title MyApplication --coverage 67 --junit myJunitReport.xml --output utsukushii.json",
	Run:     handle,
}

func init() {
	rootCmd.Flags().StringVarP(&title, "title", "t", "Utsukushii Report", "Report title")
	rootCmd.Flags().StringVar(&junit, "junit", "", "Junit report path for generate")
	rootCmd.Flags().Int16Var(&coverage, "coverage", 0, "Code coverage in percent")
	rootCmd.Flags().StringVarP(&output, "output", "o", "utsukushii.json", "Target utsukushii file path")

	rootCmd.MarkFlagRequired("junit")
}
