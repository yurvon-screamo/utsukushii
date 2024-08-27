package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yurvon-screamo/utsukushii/serve"
)

func handleServe(cmd *cobra.Command, args []string) {

	serve.RunUiHandle()
	serve.RunUtsukushiiFsContentHandle(content)

	serve.ListenAndServe(addr, open)
}

var addr string
var content string
var open bool

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Serve report",
	Long:    `Run server and serve json report in beautify ui`,
	Example: "\nutsukushii serve --addr :8080 --content utsukushii.json",
	Run:     handleServe,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&addr, "addr", "a", ":8080", "Address for listen and serve")
	serveCmd.Flags().StringVarP(&content, "content", "c", "utsukushii.json", "Utsukushii content file for serve")
	serveCmd.Flags().BoolVar(&open, "open-browser", true, "Open default browser with report after start")
}
