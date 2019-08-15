//Package raindrops returns a string representation of factorials
package raindrops

import "strconv"

//Convert converts a factorial into a well defined string
func Convert(a int) string {
	s := ""

	for i := 1; i <= a; i++ {
		m := a % i
		if m == 0 {
			/*
			   If the number has 3 as a factor, output 'Pling'.
			   If the number has 5 as a factor, output 'Plang'.
			   If the number has 7 as a factor, output 'Plong'.
			*/
			if i == 3 {
				s = s + "Pling"
			}
			if i == 5 {
				s = s + "Plang"
			}
			if i == 7 {
				s = s + "Plong"
			}
		}
	}
	//	If the number does not have 3, 5, or 7 as a factor, just pass the number's digits straight through.
	if s == "" {
		s = strconv.Itoa(a)
	}
	return s

}
