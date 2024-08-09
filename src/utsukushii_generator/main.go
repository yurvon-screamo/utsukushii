package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"utsukushii_generator/junit_input"
	"utsukushii_generator/model"

	"github.com/spf13/cobra"
)

type InputConverter interface {
	Convert() model.TestReport
}

func main() {
	Execute()

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

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
		love by spf13 and friends in Go.
		Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(*cmd)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
