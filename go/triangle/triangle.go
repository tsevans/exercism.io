// Package triangle is an exercism exercise to determine the class of a given triangle
package triangle

import "math"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
    NaT = "NaT" // not a triangle
    Equ = "Equ" // equilateral
    Iso = "Iso" // isosceles
    Sca = "Sca" // scalene
)

// KindFromSides determines the kind of a triangle based on the length of its sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind

    if a <= 0 || math.IsNaN(a) || math.IsInf(a, 1) ||
       b <= 0 || math.IsNaN(b) || math.IsInf(b, 1) ||
       c <= 0 || math.IsNaN(c) || math.IsInf(c, 1) {
        k = NaT
    } else if a == b && b == c {
        k = Equ
    } else if a == b || a == c || b == c {
        k = Iso
    }

    if a + b < c || a + c < b || b + c < a {
        k = NaT
    }

    if k == "" {
        k = Sca
    }

	return k
}
