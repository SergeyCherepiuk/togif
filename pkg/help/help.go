package help

import (
	"fmt"
	"io"
	"reflect"
	"slices"
	"strings"
	"text/template"

	"github.com/SergeyCherepiuk/togif/pkg/config"
	"github.com/SergeyCherepiuk/togif/pkg/internal"
)

var helpTemplate *template.Template

func init() {
	helpTemplate = template.Must(template.New("help").Parse(
		`A tool for converting videos into GIF images

Usage: togif [OPTIONS] [FILE]
* If input file is omitted stdin will be used

List of available options:
{{range .OptionsInfo}}    {{.}}
{{end}}`,
	))
}

func Display(out io.Writer) {
	data := struct{ OptionsInfo []string }{optionsInfo()}
	helpTemplate.Execute(out, data)
}

func optionsInfo() []string {
	var config config.Config
	rt := reflect.TypeOf(config)

	optionsInfo := make([][]string, 0)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		options := internal.Filter[string](
			[]string{
				fmt.Sprintf("-%s", field.Tag.Get("short")),
				fmt.Sprintf("--%s", field.Tag.Get("long")),
			},
			func(s string) bool { return strings.ReplaceAll(s, "-", "") != "" },
		)

		if !slices.Equal[[]string](options, []string{}) {
			optionsInfo = append(optionsInfo, make([]string, 0))
			lastIndex := len(optionsInfo) - 1
			optionsInfo[lastIndex] = append(optionsInfo[lastIndex], strings.Join(options, ", "))
			optionsInfo[lastIndex] = append(optionsInfo[lastIndex], field.Tag.Get("info"))
		}
	}

	return internal.Tabulate(optionsInfo, 3)
}
