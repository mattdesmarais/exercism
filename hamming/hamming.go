/*
Package hamming implements helpers for the DNA strand problem
*/

package hamming

import "errors"

// Distance compares a and b to get the Hamming distance between these two DNA strands.
func Distance(a, b string) (int, error) {
	//Empty Strand
	if a == "" && b == "" {
		return 0, nil
	}

	if (len(a) > 0 && len(b) > 0) && (len(a) == len(b)) {
		d := 0 //holds the current distant count
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				d++
			}
		}
		return d, nil
	}
	return 0, errors.New("error")
}
