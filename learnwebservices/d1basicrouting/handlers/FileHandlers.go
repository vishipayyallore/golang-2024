// handlers/FileHandlers.go

package handlers

import (
	"net/http"
)

const (
	HomeHtmlPath = "./static/home.html"
)

func ServerHomeHtml(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, HomeHtmlPath)
}
