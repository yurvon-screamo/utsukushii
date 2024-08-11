package cmd

import (
	"fmt"
	"io/ioutil"
	"utsukushii_generator/json_output"
	"utsukushii_generator/junit_input"

	"github.com/spf13/cobra"
)

func handleGen(cmd *cobra.Command, args []string) {
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

var title string
var junit string
var coverage int16
var output string

var genCmd = &cobra.Command{
	Use:     "gen",
	Short:   "Print the version number of Hugo",
	Long:    `All software has versions. This is Hugo's`,
	Example: "\nutsukushii_generator --title MyApplication --coverage 67 --junit myJunitReport.xml --output utsukushii.json",
	Run:     handleGen,
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&title, "title", "t", "Utsukushii Report", "Report title")
	genCmd.Flags().StringVar(&junit, "junit", "", "Junit report path for generate")
	genCmd.Flags().Int16Var(&coverage, "coverage", 0, "Code coverage in percent")
	genCmd.Flags().StringVarP(&output, "output", "o", "utsukushii.json", "Target utsukushii file path")

	genCmd.MarkFlagRequired("junit")
}
