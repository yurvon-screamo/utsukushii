package serve

import (
	"log"
	"net/http"

	"github.com/pkg/browser"
)

func ListenAndServe(addr string, openBrowser bool) {
	url := addr
	if addr[0] == ':' {
		url = "localhost" + url
	}
	url = "http://" + url

	log.Println("Starting HTTP server at", url)

	if openBrowser {
		browser.OpenURL(url)
	}

	log.Fatal(http.ListenAndServe(addr, nil))
}
