package main

import (
    "fmt"
    "math"
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
    return math.Floor(z)

}

func main() {

    var usrNum float64 

    fmt.Println("Please enter a number")

    fmt.Scanln(&usrNum)

    fmt.Println(Sqrt(usrNum))
}
