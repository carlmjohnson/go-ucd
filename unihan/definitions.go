package unihan

// Definitions is a map from runes to their Definitions in the Unicode Unihan standard
var Definitions = definitions

// Rune is an alias for rune that adds a String() method for displaying its name
type Rune rune

func (r Rune) String() string {
	return Definitions[rune(r)]
}
