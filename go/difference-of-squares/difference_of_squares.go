// Package diffsquares finds the difference between the square of sum and sum of squares of numbers
package diffsquares

import "math"

// SumOfSquares finds the difference between sums/squares up to n numbers
func SumOfSquares(n int) int {

    var s int

    for x := 1; x <= n; x++ {
        s += int(math.Pow(float64(x), float64(2)))
    }

    return s
}

func SquareOfSums(n int) int {

    var s int

    for x := 1; x <= n; x++ {
        s += x
    }

    return int(math.Pow(float64(s), float64(2)))
}

func Difference(n int) int {

    return SquareOfSums(n) - SumOfSquares(n)
}
