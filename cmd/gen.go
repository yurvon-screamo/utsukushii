package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/yurvon-screamo/utsukushii/gotest_input"
	"github.com/yurvon-screamo/utsukushii/json_output"
	"github.com/yurvon-screamo/utsukushii/junit_input"
	"github.com/yurvon-screamo/utsukushii/model"
)

func handleGen(cmd *cobra.Command, args []string) {
	report := model.TestReport{
		Title:     title,
		Coverage:  coverage,
		Duration:  0,
		Timestamp: time.Now().UTC(),
	}

	for _, junit := range junits {
		junitData, err := os.ReadFile(junit)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		tests, err := junit_input.Convert(junitData)
		if err != nil {
			fmt.Println("Error convert junit:", err)
			return
		}

		report.Enrich(tests)
	}

	for _, goJson := range goTests {
		goJsonData, err := os.ReadFile(goJson)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		tests, err := gotest_input.Convert(goJsonData)
		if err != nil {
			fmt.Println("Error convert gotest:", err)
			return
		}

		report.Enrich(tests)
	}

	jsonPrinter := json_output.JsonPrinter{
		Report: &report,
	}

	jsonData, err := jsonPrinter.Print()
	if err != nil {
		fmt.Println("Error convert json:", err)
		return
	}

	err = os.WriteFile(output, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Printf("Test report written to %s\n", output)
}

var title string
var goTests []string
var junits []string
var coverage int16
var output string

var genCmd = &cobra.Command{
	Use:     "gen",
	Short:   "Generate report content",
	Long:    `Generate report json-content file from your junit report or go test output`,
	Example: "\nutsukushii gen --title MyApplication --coverage 67 --junit myJunitReport.xml --go-json-test myGoJson.log --output utsukushii.json",
	Run:     handleGen,
}

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&title, "title", "t", "Utsukushii Report", "Report title")
	genCmd.Flags().StringArrayVar(&junits, "junit", nil, "Junit reports paths for generate")
	genCmd.Flags().StringArrayVar(&goTests, "go-json-test", nil, "Paths to files with 'go test --json' stdout")
	genCmd.Flags().Int16Var(&coverage, "coverage", 0, "Code coverage in percent")
	genCmd.Flags().StringVarP(&output, "output", "o", "utsukushii.json", "Target utsukushii file path")

	genCmd.MarkFlagsOneRequired("junit", "go-json-test")
}
