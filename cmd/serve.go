package cmd

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"github.com/yurvon-screamo/utsukushii/utsukushii_ui"
)

func handleServe(cmd *cobra.Command, args []string) {
	distFS, err := fs.Sub(utsukushii_ui.UtsukushiiUi, "staticBuild")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/utsukushii.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.Handle("/", http.FileServer(http.FS(distFS)))

	log.Println("Starting HTTP server at", addr)

	if open {
		url := addr
		if addr[0] == ':' {
			url = "localhost" + url
		}
		url = "http://" + url
		browser.OpenURL(url)
	}

	log.Fatal(http.ListenAndServe(addr, nil))
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
