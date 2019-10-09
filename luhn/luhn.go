//Package luhn determines if a string (credit card, social insurance) is valid according to luhn's algorithm by returning true or false.
package luhn

import (
	"strings"
	"strconv"
)


//Valid return true if the string passes Luhn algorithm  
func Valid(i string) bool {
	t := 0 //total 
	v := prechecks(i)

	if v != "" {
		b := reverse(v)
		for pos, c := range b {
			s := 0
			//finds even positions 
			if pos % 2 == 1 {
				s = int(c - '0') * 2 
				if s > 9 { 
					s = s - 9 
				} 
				t = t + s 
			} else {
				t = t + int(c - '0') 
			}	
		}
		return ((t % 10) == 0) 

	} else {
		return false 
	}

	return false 
}

//reverse returs a string in reverse order 
func reverse(l string) string {
	e := ""
	for i := range l {
		v := l[len(l) - i - 1]
		e = e + string(v)
	}
	return e 
}

//prechecks validates the length and ensures there are no space, and that it can be converted to an int.
func prechecks(i string) string {

	//remove spaces 
	n := strings.Replace(i, " ", "", -1)

	//if n can be converted to an int, then this ensure it does't contain any non-digit characters 
	if _, err := strconv.Atoi(n); err == nil {
		if len(n) > 1 {
			return n
		}	
	}
	return ""
}
