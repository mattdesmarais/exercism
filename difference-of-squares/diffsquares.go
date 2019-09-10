//Package diffsquares implements helper functions get the difference of squares
package diffsquares

//Difference returns the difference of squars
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

//SquareOfSum calculates the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	s := (n * (n + 1)) / 2
	return int(float64(s * s))
}

//SumOfSquares calculates the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	s := (n * (n + 1) * ((2 * n) + 1)) / 6
	return int(float64(s))
}
