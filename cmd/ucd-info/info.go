package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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

func inRange(r, start, end rune) bool {
	return (r >= start) && (r <= end)
}

func normalizeInput(input string) string {
	i, j := 0, 0

	// Trim prefix
	if len(input) >= 2 {
		if prefix := input[:2]; prefix == "0x" || prefix == "0X" {
			i, j = 2, 2
		}
	}

	// Trim suffix
	var r rune
	for _, r = range input[i:] {
		if inRange(r, '0', '9') || inRange(r, 'a', 'f') || inRange(r, 'A', 'F') {
			j++
			continue
		}
		break
	}
	return input[i:j]
}

func getRuneFromCodepoint(input string, base int) (r rune, err error) {
	cp, err := strconv.ParseInt(input, base, 32)
	if err != nil {
		return 0, err
	}
	return rune(cp), nil
}

func die(err error, input string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not process as codepoint:", input)
		os.Exit(1)
	}
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
			r, err := getRuneFromCodepoint(normalizeInput(arg), base)
			die(err, arg)
			print(r)
		}
	} else {
		s := bufio.NewScanner(os.Stdin)

		for s.Scan() {
			input := s.Text()
			r, err := getRuneFromCodepoint(normalizeInput(input), base)
			die(err, input)
			print(r)
		}
		// Ignoring errors
	}
}
