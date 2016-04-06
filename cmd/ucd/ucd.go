package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/carlmjohnson/unicodechess/unicodedata"
	"github.com/carlmjohnson/unicodechess/unihan"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		byCharacter(os.Stdin)
		return
	}

	var buf bytes.Buffer
	for _, s := range flag.Args() {
		buf.WriteString(s)
	}
	byCharacter(&buf)
}

func byCharacter(input io.Reader) {
	s := bufio.NewScanner(input)
	s.Split(bufio.ScanRunes)

	for s.Scan() {
		s := s.Text()
		r := []rune(s)[0]
		print(r)
	}
}

func print(r rune) {
	name := unicodedata.Rune(r).String()

	if definition, ok := unihan.Definitions[r]; ok {
		name = definition
	}

	s := string(r)

	if name == "<control>" {
		replacementChar := rune(0xfffd)
		s = string(replacementChar)
	}

	fmt.Printf("0x%04X\t%s\t%s\n", r, s, name)

}
