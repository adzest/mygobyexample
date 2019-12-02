// Go supports
// <a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a>.
// Here's a classic factorial example.

package main

import ("fmt";"time"; "log")

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s", name, elapsed)
}

// This `fact` function calls itself until it reaches the
// base case of `fact(0)`.
func fact(n float64) float64 {
    defer timeTrack(time.Now(), "factorial")
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(170))
}
