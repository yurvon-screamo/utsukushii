package dev_mode

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/yurvon-screamo/utsukushii/model"
	"github.com/yurvon-screamo/utsukushii/serve"
	"github.com/yurvon-screamo/utsukushii/specification/dotnet_trx_input"
	"github.com/yurvon-screamo/utsukushii/specification/gotest_input"
	"github.com/yurvon-screamo/utsukushii/specification/json_output"
)

func RunDevMode(config *DevModeConfig) error {
	jsonData, err := getJsonData(config)
	if err != nil {
		fmt.Println("Error convert json:", err)
		return err
	}

	content := &serve.UtsukushiiContent{
		JsonContent: jsonData,
	}

	http.HandleFunc("/dev_mode/check", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })

	http.HandleFunc("/dev_mode/cmd", func(w http.ResponseWriter, r *http.Request) {
		jsonData, err = getJsonData(config)
		if err != nil {
			fmt.Println("Error convert json:", err)
			w.Write([]byte(err.Error()))
		}

		content.JsonContent = jsonData
	})

	serve.RunUiHandle()
	serve.RunUtsukushiiContentHandle(content)

	serve.ListenAndServe(config.ServeAddress, config.OpenBrowser)

	return nil
}

func getJsonData(config *DevModeConfig) ([]byte, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("pwsh", "-c", config.Cmd)
	} else {
		cmd = exec.Command("sh", "-c", config.Cmd)
	}

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("error on run cmd", err)
	}

	var testFile *model.TestFile
	if config.Lang == golang {
		testFile, err = gotest_input.Convert(out)
		if err != nil {
			fmt.Println("error on convert gotest_input", err)
			return nil, err
		}
	} else if config.Lang == dotnet {
		var fileTrx string
		for _, line := range strings.Split(string(out), "\n") {
			if strings.HasSuffix(line, ".trx") {
				parts := strings.SplitN(line, ":", 2)
				if len(parts) == 2 {
					fileTrx = strings.TrimSpace(parts[1])
				}
			}
		}

		if fileTrx == "" {
			return nil, fmt.Errorf("Not found trx file path")
		}

		raw, err := os.ReadFile(fileTrx)
		if err != nil {
			fmt.Println("Error while read trx report:", err)
			return nil, err
		}

		testFile, err = dotnet_trx_input.Convert(raw)
		if err != nil {
			fmt.Println("error on convert dotnet_trx_input", err)
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("Unknown lang %s. Pls use one of: 'go', '.net'", config.Lang)
	}

	report := model.TestReport{
		Title:     config.Title,
		Coverage:  config.Coverage,
		Timestamp: time.Now().UTC(),
		Duration:  0,
	}

	report.Enrich(testFile)

	jsonPrinter := json_output.JsonPrinter{
		Report: &report,
	}

	jsonData, err := jsonPrinter.Print()
	if err != nil {
		fmt.Println("Error convert json:", err)
		return nil, err
	}

	return jsonData, nil
}
