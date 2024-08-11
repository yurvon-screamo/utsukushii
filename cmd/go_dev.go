package cmd

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/yurvon-screamo/utsukushii/gotest_input"
	"github.com/yurvon-screamo/utsukushii/json_output"
	"github.com/yurvon-screamo/utsukushii/model"
)

func handleGoDev(cmd *cobra.Command, args []string) {
	report := model.TestReport{
		Title:     goDevTitle,
		Coverage:  goDevCoverage,
		Timestamp: time.Now().UTC(),
		Duration:  0,
	}

	stdin := os.Stdin
	stdinScanner := bufio.NewScanner(stdin)
	stdinScanner.Split(bufio.ScanLines)

	raw := []byte{}
	for stdinScanner.Scan() {
		raw = append(raw, stdinScanner.Bytes()...)
		raw = append(raw, byte('\n'))
	}

	tests, err := gotest_input.Convert(raw)
	if err != nil {
		fmt.Println("Error convert gotest:", err)
		return
	}

	report.Enrich(tests)

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

	handleServe(cmd, args)
}

var goDevTitle string
var goDevCoverage int16

var goDevCmd = &cobra.Command{
	Use:     "go-dev",
	Short:   "Generate and serve report content from go test stdout",
	Long:    `Generate and serve report from go test stdout`,
	Example: "\ngo test -v --json ./... > utsukushii go-dev",
	Run:     handleGoDev,
}

func init() {
	rootCmd.AddCommand(goDevCmd)

	goDevCmd.Flags().StringVarP(&goDevTitle, "title", "t", "Utsukushii Report", "Report title")
	goDevCmd.Flags().Int16Var(&goDevCoverage, "coverage", 0, "Code coverage in percent")

	goDevCmd.Flags().StringVarP(&addr, "addr", "a", ":8080", "Address for listen and serve")
	goDevCmd.Flags().BoolVar(&open, "open-browser", true, "Open default browser with report after start")
	goDevCmd.Flags().StringVarP(&content, "content", "c", "utsukushii.json", "Utsukushii content file for serve")
}
