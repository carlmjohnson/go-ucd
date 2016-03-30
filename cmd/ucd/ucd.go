package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/carlmjohnson/unicodechess/unicodedata"
)

var input io.Reader = os.Stdin

func init() {
	flag.Parse()
	if flag.NArg() == 0 {
		return
	}
	var buf bytes.Buffer
	for _, s := range flag.Args() {
		buf.WriteString(s)
	}
	input = &buf
}

func main() {
	s := bufio.NewScanner(input)
	s.Split(bufio.ScanRunes)

	for s.Scan() {
		s := s.Text()
		r := []rune(s)[0]
		name := unicodedata.UCD[r]
		fmt.Printf("0x%04X\t%s\t%s\n", r, s, name)
	}
}
