package main

import (
	"bytes"
	"text/template"
)

type (
	// DefaultTemplate arguments for DEFAULT_TEMPLATE
	DefaultTemplate struct {
		Color           string
		Title           string
		MarkdownContent string
	}
)

const (
	// DEFAULT default template
	DEFAULT = `
{
    "msg_type": "interactive",
    "card": {
        "config": {
            "wide_screen_mode": true
        },
        "header": {
            "template": "{{.Color}}",
            "title": {
                "content": "{{.Title}}",
                "tag": "plain_text"
            }
        },
        "elements": [
            {
                "tag": "markdown",
                "content": {{.MarkdownContent}}
            }
        ]
    }
}
`
)

func (d DefaultTemplate) Content() (string, error) {
	var buf bytes.Buffer
	t := template.Must(template.New("default_template").Parse(DEFAULT))
	if err := t.Execute(&buf, d); err != nil {
		return "", err
	}
	return buf.String(), nil
}
