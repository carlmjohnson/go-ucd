// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates unicodedata.go. It can be invoked by running
// go generate
package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

const unicodeDataURL = "http://unicode.org/Public/UCD/latest/ucd/UnicodeData.txt"

type UnicodeRecord struct {
	Codepoint, Name string
}

// getUnicodeRecords returns a slice rather than map so that the file is
// ordered consistently on repeated invocations.
func getUnicodeRecords() []UnicodeRecord {
	rsp, err := http.Get(unicodeDataURL)
	defer rsp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	records := make([]UnicodeRecord, 0)

	reader := csv.NewReader(rsp.Body)
	reader.Comma = ';'
	reader.FieldsPerRecord = 15

	for {
		record, err := reader.Read()

		if err == io.EOF {
			return records
		} else if err != nil {
			log.Fatal(err)
		}
		records = append(records, UnicodeRecord{record[0], record[1]})
	}
}

const templateString = `// go generate
// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

// This file was generated by robots at
// {{ .Timestamp }}
// using data from
// {{ .URL }}

package unicodedata

var names = map[rune]string{
{{ range .Records }}
	0x{{ .Codepoint }}: {{ printf "%q" .Name }},{{ end }}
}
`

var packageTemplate = template.Must(template.New("package").Parse(templateString))

var debug = flag.Bool("debug", false, "Enable to output to stdout instead of writing unicodedata.go")

func init() {
	flag.Parse()
}

func main() {
	records := getUnicodeRecords()
	data := struct {
		URL       string
		Timestamp time.Time
		Records   []UnicodeRecord
	}{
		URL:       unicodeDataURL,
		Timestamp: time.Now().UTC(),
		Records:   records,
	}

	var file io.WriteCloser

	if *debug {
		file = os.Stdout
	} else {
		var err error
		if file, err = os.Create("unicodedata.go"); err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	if err := packageTemplate.Execute(file, data); err != nil && file != os.Stdout {
		log.Fatal(err)
	}
}
