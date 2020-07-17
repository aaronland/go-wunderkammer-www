package templates

import (
	"context"
	"errors"
	"fmt"
	html_template "html/template"
	txt_template "text/template"
)

func LoadHTMLTemplates(ctx context.Context, path_templates string) (*html_template.Template, error) {

	t := html_template.New("orthis").Funcs(html_template.FuncMap{
		"TemplateURL": func(raw string) html_template.URL {
			return html_template.URL(raw)		
		},
	})

	if path_templates != "" {

		parsed, err := t.ParseGlob(path_templates)

		if err != nil {
			msg := fmt.Sprintf("Failed to parse templates (%s), %v", path_templates, err)
			return nil, errors.New(msg)
		}

		t = parsed

	} else {

		return nil, errors.New("Bundled templates are not implemented yet.")

		/*
			for _, name := range templates.AssetNames() {

				body, err := templates.Asset(name)

				if err != nil {
					return nil, err
				}

				t, err = t.Parse(string(body))

				if err != nil {
					return nil, err
				}
			}
		*/
	}

	return t, nil
}

func LoadXMLTemplates(ctx context.Context, path_templates string) (*txt_template.Template, error) {

	t := txt_template.New("orthis").Funcs(txt_template.FuncMap{})

	if path_templates != "" {

		parsed, err := t.ParseGlob(path_templates)

		if err != nil {
			msg := fmt.Sprintf("Failed to parse templates (%s), %v", path_templates, err)
			return nil, errors.New(msg)
		}

		t = parsed

	} else {

		return nil, errors.New("Bundled templates are not implemented yet.")
	}

	return t, nil
}
