// euler.go
// My efforts to learn the Go syntax by solving Project Euler problems.
// https://projecteuler.net/
package main

import (
    "os"
    "path/filepath"
    "math"
    "fmt"
)

func main() {
    if len(os.Args) == 1 {
        fmt.Printf("Output solution to problem <integer>. Usage: %s <integer>\n", filepath.Base(os.Args[0]))
        os.Exit(1)
    }
    problemNo := os.Args[1]
    if (problemNo == "1") {
        problemOne()
    }
}

func problemOne() {
    // Problem 1
    // Multiples of 3 and 5
    // If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
    // The sum of these multiples is 23.
    // Find the sum of all the multiples of 3 or 5 below 1000.
    sum := 0.0
    // Loop from 1 to 999; if i modulo 3 or i modulo 5 is 0, add i to the sum.
    for i := 1.0; i < 1000.0; i++ {
        if (math.Mod(i, 3.0) == 0 || math.Mod(i, 5.0) == 0) {
            sum = sum + i
        }
    }
    fmt.Println(int(sum))
}
