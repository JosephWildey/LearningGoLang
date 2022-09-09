package main

import (
    "fmt"
)


// This function loops ten times and estimates the square root of a number.

func Sqrt(x float64) float64 {

    // initialize z at 1.0
    z := 1.0

    // loop ten times, using a common formula for estimating a square root
    for i := 0; i < 10; i++ {

        z -= (z * z - x) / (2 * z)

    }

    // returns the estimated square root
    return z

}

func main() {
    fmt.Println(Sqrt(16))
}
