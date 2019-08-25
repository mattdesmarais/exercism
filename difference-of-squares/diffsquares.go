//Package diffsquares implements helper functions get the difference of squares 
package diffsquares

import (
	"math"
)

func Difference (n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

//SqSum calculates the square of the sum of the first n natural numbers 
func SquareOfSum(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s = s + i 	
	}
	return int(math.Pow(float64(s),2))
}

//SumSq calculates the sum of the squares of the first n natural numbers 
func SumOfSquares(n int) int {
//1² + 2² + ... + 10² = 385.
	var s float64
	for i := 1; i <= n; i++ {
		s = s + math.Pow(float64(i),2) 	
	}
	return int(s)
}
