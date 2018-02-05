// Package sieve is an exercism exercise implementing the sieve of eratosthenes algorithm
package sieve

// Sieve finds all prime numbers up to a given number limit
func Sieve(n int) []int {

    nums := make([]int, n-1)
    for x := 0; x < n-1; x++ {
        nums[x] = x + 2
    }

    mask := make([]bool, n-1)
    for y := 0; y < len(nums); y++ {
        if !mask[y] {
            for i := y+1; i < len(nums); i++ {
                if nums[i] % nums[y] == 0 {
                    mask[i] = true
                }
            }
        }
    }

    var ans []int
    for z := 0; z < len(mask); z++ {
        if !mask[z] {
            ans = append(ans, nums[z])
        }
    }

    return ans
}
