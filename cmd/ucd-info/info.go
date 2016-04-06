package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"unicode"

	"github.com/carlmjohnson/unicodechess/unicodedata"
	"github.com/carlmjohnson/unicodechess/unihan"
)

func scripts(r rune) (scripts []string) {
	for script, rt := range unicode.Scripts {
		if unicode.In(r, rt) {
			scripts = append(scripts, script)
		}
	}
	return scripts
}

var templateString = `Codepoint: {{printf "0x%04X" .Rune}}	Decimal: {{printf "%d" .Rune}}
{{- with .Name}}
Name: {{.}}
{{- end}}
Display: {{ .Display }}
{{- with .Definition }}
Definition: {{ . }}
{{- end}}
Scripts: {{range .Scripts}}{{.}} {{end}}

`

var packageTemplate = template.Must(template.New("package").Parse(templateString))

func print(r rune) {
	packageTemplate.Execute(os.Stdout, struct {
		Definition string
		Display    string
		Name       string
		Rune       rune
		Scripts    []string
	}{
		Definition: unihan.Definitions[r],
		Display:    string(r),
		Name:       unicodedata.Names[r],
		Rune:       r,
		Scripts:    scripts(r),
	})
}

func byCodepoint(input string, base int) {
	input = strings.TrimPrefix(input, "0x")
	cp, err := strconv.ParseInt(input, base, 32)
	if err != nil {
		fmt.Printf("Could not process codepoint, base-%d: %v\n", base, err)
		os.Exit(1)
	}
	print(rune(cp))
}

func main() {
	dec := flag.Bool("use-decimal", false, "Look up codepoint by decimal (default hex)")
	flag.Parse()

	base := 16
	if *dec {
		base = 10
	}

	if flag.NArg() > 0 {
		for _, arg := range flag.Args() {
			byCodepoint(arg, base)
		}
	} else {
		s := bufio.NewScanner(os.Stdin)

		for s.Scan() {
			byCodepoint(s.Text(), base)
		}
		// Ignoring errors
	}
}
