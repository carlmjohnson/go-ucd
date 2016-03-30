package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/carlmjohnson/unicodechess/unicodedata"
)

func main() {
	hcp := flag.String("codepoint", "", "Display data for hex codepoint")
	dcp := flag.String("dec-codepoint", "", "Display data for decimal codepoint")
	flag.Parse()

	if *hcp != "" {
		byCodepoint(*hcp, 16)
		return
	}

	if *dcp != "" {
		byCodepoint(*dcp, 10)
		return
	}

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

func byCodepoint(input string, base int) {
	cp, err := strconv.ParseInt(input, base, 32)
	if err != nil {
		fmt.Printf("Could not process codepoint, base-%d: %v\n", base, err)
		os.Exit(1)
	}
	print(rune(cp))
}

func print(r rune) {
	name := unicodedata.UCD[r]
	fmt.Printf("0x%04X\t%s\t%s\n", r, string(r), name)
}
