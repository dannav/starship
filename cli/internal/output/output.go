// Package output handles writing text/templates to stdout
package output

import (
	"os"
	"text/template"
)

const (
	// SearchType represents an output for the search command
	SearchType = iota
	// IndexType represents an output for the index command
	IndexType
	// HelpType represents an output for the help command
	HelpType
)

// typeToTemplate retrieves template text from a template type
var typeToTemplate = map[int]string{
	SearchType: search,
}

// Template represents the template that will be used to write to stdout
type Template struct {
	Type int
	Data interface{}
}

// NewTemplate returns a new template builder
func NewTemplate(t int, d interface{}) Template {
	return Template{
		Type: t,
		Data: d,
	}
}

// Write attempts to write the string to stdout using the template type
func (t *Template) Write() error {
	tmpl, err := template.New(string(t.Type)).Parse(typeToTemplate[t.Type])
	if err != nil {
		return err
	}

	err = tmpl.Execute(os.Stdout, t.Data)
	if err != nil {
		return err
	}

	return nil
}
