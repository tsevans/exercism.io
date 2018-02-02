// Package hamming is an exercism exercise to find the hamming distance between two DNA strands
package hamming

import "errors"

// Dinstance finds the Hamming distance between a and b strands of DNA
func Distance(a, b string) (int, error) {

    if len(a) != len(b) {
        return -1, errors.New("Different length strands")
    }

    count := 0
    for x := 0; x < len(a); x++ {
        if a[x] != b[x] {
            count++
        }
    }

    return count, nil
}
