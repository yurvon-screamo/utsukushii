package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yurvon-screamo/utsukushii/dev_mode"
	"gopkg.in/yaml.v3"
)

type devModeYamlConfig struct {
	Cmd               string `yaml:"cmd"`
	Lang              string `yaml:"lang"`
	Title             string `yaml:"title"`
	Coverage          int16  `yaml:"coverage"`
	ServeAddress      string `yaml:"addr"`
	IgnoreOpenBrowser bool   `yaml:"ignore_open"`
}

func handleDevMode(cmd *cobra.Command, args []string) {
	rawConfig, err := os.ReadFile(devModeConfig)
	if err != nil {
		log.Fatalf("error read config: %v", err)
	}

	config := devModeYamlConfig{}
	err = yaml.Unmarshal([]byte(rawConfig), &config)
	if err != nil {
		log.Fatalf("error unmarshal: %v", err)
	}

	if config.ServeAddress == "" {
		config.ServeAddress = ":8080"
	}
	if config.Title == "" {
		config.Title = "Utsukushii dev"
	}

	devModeConfig := &dev_mode.DevModeConfig{
		Cmd:          config.Cmd,
		Lang:         config.Lang,
		Title:        config.Title,
		Coverage:     config.Coverage,
		ServeAddress: config.ServeAddress,
		OpenBrowser:  !config.IgnoreOpenBrowser,
	}

	err = dev_mode.RunDevMode(devModeConfig)
	if err != nil {
		log.Fatalf("run dev mode: %v", err)
	}
}

var devModeConfig string

var devModeCmd = &cobra.Command{
	Use:   "dev",
	Short: "Serve and open UI in dev mode",
	Long:  `Run your tests from utsukushii report and get feedback immediately`,
	Run:   handleDevMode,
}

func init() {
	rootCmd.AddCommand(devModeCmd)

	devModeCmd.Flags().StringVarP(&devModeConfig, "config", "c", "utsukushii-dev.yaml", "Path to dev mode config")
}
