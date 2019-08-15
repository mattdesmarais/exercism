//Package scrabble translate a scrabble word to a total numerical value
package scrabble

import (
	"strings"
)

//Score finds the total value for a given word
func Score(word string) int {
	t := 0
	for _, v := range word {
		t = t + FindValue(strings.ToUpper(string(v)))
	}
	return t
}

//FindValue finds the numerical value for a given string
func FindValue(letter string) int {
	switch letter {
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		return 1
	case "D", "G":
		return 2
	case "B", "C", "M", "P":
		return 3
	case "F", "H", "V", "W", "Y":
		return 4
	case "K":
		return 5
	case "J", "X":
		return 8
	case "Q", "Z":
		return 10
	}
	return 0 //Not in the scrabble Score
}
