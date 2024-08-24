package dev_mode

import (
	"fmt"
	"net/http"
	"os/exec"
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
	var testFile *model.TestFile

	out, err := exec.Command(config.Cmd).Output()
	if err != nil {
		fmt.Println("error on run cmd", err)
	}

	if config.Lang == golang {
		testFile, err = gotest_input.Convert(out)
		if err != nil {
			fmt.Println("error on convert gotest_input", err)
			return nil, err
		}
	} else if config.Lang == dotnet {
		testFile, err = dotnet_trx_input.Convert(out)
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
