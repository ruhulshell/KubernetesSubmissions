package main

import (
	"html/template"
	"io/fs"
	"path"
	"ruhultodoapp/ui"
)

type templateData struct {
	PageData
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := path.Base(page)

		ts, err := template.New(name).ParseFS(ui.Files, page)
		if err != nil {
			return nil, err
		}

		// ts, err = ts.ParseFS(ui.Files, "*.layout.gohtml")
		// if err != nil {
		// 	return nil, err
		// }

		cache[name] = ts
	}

	return cache, nil
}
