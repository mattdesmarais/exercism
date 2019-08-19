//Package isogram determines if a word or phrase is a isogram
package isogram

import (
	"strings"
)

//Global Array of the english alpahbet
var a = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

//IsIsogram returns true or false depending on what the lower methods discovers
func IsIsogram(w string) bool {
	//Array stores keeps the counts for each letter in w
	var c = [26]int{}
	//Process the word
	for _, v := range w {
		//Find letter in global array "a"
		for i, e := range a {
			if (strings.ToUpper(string(v))) == e {
				c[i]++
				break
			}
		}
	}
	//eval c array for duplicates
	for _, e := range c {
		if e >= 2 {
			return false
		}
	}
	return true
}
