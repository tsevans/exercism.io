// Package reverse is an exercism exercise to reverse a string
package reverse

// String reverses the given string and returns it
func String(s string) string {

    b := []rune(s)
    for x := 0; x < len(b)/2; x++ {
        y := len(b)-x-1
        b[x], b[y] = b[y], b[x]
    }
    return string(b)
}
