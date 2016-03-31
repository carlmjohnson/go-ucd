package unicodedata

// Names is a map from runes to their names in the Unicode standard
var Names = names

// Rune is an alias for rune that adds a String() method for displaying its name
type Rune rune

func (r Rune) String() string {
	return Names[rune(r)]
}
