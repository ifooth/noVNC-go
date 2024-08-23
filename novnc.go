// Licensed under the Apache License 2.0.

// Package novnc is a embed module for noVNC.
package novnc

import (
	"embed"
	"io/fs"
	"net/http"
	"text/template"
)

//go:embed dist
var distFS embed.FS

var (
	assets = must(fs.Sub(distFS, "dist"))
	tpl    = template.Must(template.New("").ParseFS(distFS, "dist/*.html"))
)

// FS noVNC dist FS
func FS() fs.FS {
	return assets
}

// Handler noVNC dist handler
func Handler() http.Handler {
	return http.FileServer(http.FS(assets))
}

// Template noVNC html template
func Template() *template.Template {
	return tpl
}

// must panic on error
func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
