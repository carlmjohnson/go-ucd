package main

import "strings"

// matchWord returns true if s contains all the runes in pat in order.
// E.g. "a1b2c3" matches "abc" but "cba" does not match "abc".
func matchWord(s, pat string) bool {
	for _, c := range pat {
		first := strings.IndexRune(s, c)
		if first == -1 {
			return false
		}
		s = s[first:]
	}
	return true
}

func matchmaker(pat string) func(string) bool {
	// Normalize and split
	pats := strings.Split(strings.ToUpper(pat), " ")

	return func(s string) bool {
		for _, pat := range pats {
			// Check for negative
			if strings.HasPrefix(pat, "-") {
				if strings.Contains(s, pat[1:]) {
					return false
				}
			} else if !matchWord(s, pat) {
				return false
			}
		}
		return true
	}
}

func main() {
	// s := bufio.NewScanner(os.Stdin)
	// for s.Scan() {
	// 	m := matchmaker(s.Text())

	// 	var matches []string

	// 	for _, v := range unicodedata.UCD {
	// 		if m(v) {
	// 			matches = append(matches, v)
	// 		}
	// 	}

	// 	sort.Strings(matches)
	// 	for _, match := range matches {
	// 		fmt.Println(match)
	// 	}
	// }

	// if err := s.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
