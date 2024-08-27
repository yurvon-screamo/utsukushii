package serve

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/yurvon-screamo/utsukushii/utsukushii_ui"
)

func RunUiHandle() {
	distFS, err := fs.Sub(utsukushii_ui.UtsukushiiUi, "staticBuild")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(distFS)))
}

type UtsukushiiContent struct {
	JsonContent []byte
}

func RunUtsukushiiContentHandle(content *UtsukushiiContent) {
	http.HandleFunc("/utsukushii.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write(content.JsonContent)
	})
}

func RunUtsukushiiFsContentHandle(path string) {
	http.HandleFunc("/utsukushii.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path)
	})
}
