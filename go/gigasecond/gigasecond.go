// Package gigasecond is an exercism exercise for calculaing the moment a person has lived 10^9 seconds
package gigasecond

import "time"

// AddGigasecond takes a date as input and returns the date occuring 10^9 seconds later.
func AddGigasecond(t time.Time) time.Time {
    return t.Add(time.Duration(1e9)*time.Second)
}
