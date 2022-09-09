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

    // returns the estimated square root as a whole number
    return math.Floor(z)

}

func main() {

    // This variable will be the number the user wants to get the square root of
    var usrNum float64 

    // This line asks the user for input
    fmt.Println("Please enter a number")

    // Captures user input
    fmt.Scanln(&usrNum)

    // Outputs the solution
    fmt.Println(Sqrt(usrNum))
}
