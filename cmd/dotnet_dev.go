package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/yurvon-screamo/utsukushii/model"
	"github.com/yurvon-screamo/utsukushii/specification/dotnet_trx_input"
	"github.com/yurvon-screamo/utsukushii/specification/json_output"
)

func handleDotnetDev(cmd *cobra.Command, args []string) {
	report := model.TestReport{
		Title:     dotnetDevTitle,
		Coverage:  dotnetDevCoverage,
		Timestamp: time.Now().UTC(),
		Duration:  0,
	}

	stdin := os.Stdin
	stdinScanner := bufio.NewScanner(stdin)
	stdinScanner.Split(bufio.ScanLines)

	var fileTrx string

	for stdinScanner.Scan() {
		line := string(stdinScanner.Bytes())
		if strings.HasSuffix(line, ".trx") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				fileTrx = strings.TrimSpace(parts[1])
			}
		}
	}

	if fileTrx == "" {
		fmt.Println("Not found trx file path")
		return
	}

	raw, err := os.ReadFile(fileTrx)
	if err != nil {
		fmt.Println("Error while read trx report:", err)
		return
	}

	tests, err := dotnet_trx_input.Convert(raw)
	if err != nil {
		fmt.Println("Error convert dotnet trx:", err)
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

var dotnetDevTitle string
var dotnetDevCoverage int16

var dotnetDevCmd = &cobra.Command{
	Use:     "dotnet-dev",
	Short:   "Generate and serve report content from dotnet test trx logger",
	Long:    `Generate and serve report from dotnet test trx logger`,
	Example: "\ndotnet test --logger trx | utsukushii dotnet-dev",
	Run:     handleDotnetDev,
}

func init() {
	rootCmd.AddCommand(dotnetDevCmd)

	dotnetDevCmd.Flags().StringVarP(&dotnetDevTitle, "title", "t", "Utsukushii Report", "Report title")
	dotnetDevCmd.Flags().Int16Var(&dotnetDevCoverage, "coverage", 0, "Code coverage in percent")

	dotnetDevCmd.Flags().StringVarP(&addr, "addr", "a", ":8080", "Address for listen and serve")
	dotnetDevCmd.Flags().BoolVar(&open, "open-browser", true, "Open default browser with report after start")
	dotnetDevCmd.Flags().StringVarP(&content, "content", "c", "utsukushii.json", "Utsukushii content file for serve")
}
