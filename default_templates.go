package betterlog

import (
	"fmt"
	"strings"
	"text/template"
	"time"
)

func BasicLogTemplate(name string, includeCaller bool, includeTraceback bool) *template.Template {
	tmpl := &strings.Builder{}
	fmt.Fprintf(tmpl, "{{.Time.Format %#v}} ", time.DateTime)
	fmt.Fprint(tmpl, "[{{.Level.Name}}]")
	if includeCaller {
		fmt.Fprint(tmpl, " ({{(index .Traceback 0).Filename}}:{{(index .Traceback 0).LineNumber}})")
	}
	fmt.Fprint(tmpl, "{{if .Message}} {{.Message}}{{end}}")
	fmt.Fprint(tmpl, "{{if .Data}} {{.Data}}{{end}}")
	if includeTraceback {
		fmt.Fprint(tmpl, "{{if (lt 1 (len .Traceback))}}")
		fmt.Fprintln(tmpl)
		fmt.Fprint(tmpl, "\tTraceback:{{range .Traceback}}\n\t{{.Filename}}:{{.LineNumber}}{{end}}")
		fmt.Fprint(tmpl, "{{end}}")
	}
	return template.Must(template.New(name).Parse(tmpl.String()))
}
