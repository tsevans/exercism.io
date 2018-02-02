// Package raindrops is an exercism exercise to convert a number to a string
package raindrops

import "strconv"

// Convert a number to a string based on certain math properties
func Convert(n int) string {

    var conv string

    if n % 3 == 0 {
        conv += "Pling"
    }

    if n % 5 == 0 {
        conv += "Plang"
    }

    if n % 7 == 0 {
        conv += "Plong"
    }

    if conv == "" {
        conv = strconv.Itoa(n)
    }

    return conv
}
