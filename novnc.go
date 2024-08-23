// Licensed under the Apache License 2.0.

// Package novnc is a embed module for noVNC.
package novnc

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
	"text/template"
)

//go:embed dist
var distFS embed.FS

var (
	assets = must(fs.Sub(distFS, "dist"))
	tpl    = template.Must(template.New("").ParseFS(distFS, "dist/*.html"))
)

// aliasFS alias FS
type aliasFS struct {
	rawFS     fs.FS
	nameAlias map[string]string
}

// Open noVNC dist Open
func (fs *aliasFS) Open(name string) (fs.File, error) {
	for k, v := range fs.nameAlias {
		if strings.HasPrefix(name, k) {
			name = strings.ReplaceAll(name, k, v)
		}
	}
	return fs.rawFS.Open(name)
}

// FS noVNC dist FS
func FS() fs.FS {
	f := &aliasFS{
		rawFS:     assets,
		nameAlias: map[string]string{"vendor": "my_vendor"},
	}
	return f
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
