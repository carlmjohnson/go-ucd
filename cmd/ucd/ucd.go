package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cooperhewitt/go-ucd/unicodedata"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanRunes)

	for s.Scan() {
		s := s.Text()
		r := []rune(s)[0]
		name := unicodedata.UCD[r]
		fmt.Printf("0x%04X\t%s\t%s\n", r, s, name)
	}
}
