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
    problemOne()
}

func problemOne() {
    fmt.Println(math.Mod(13, 4))
}
