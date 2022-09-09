package main

import ( 
    "fmt"
)


/* This reverses the order of input strings.
   This will return a value for x, y, and z, which are all strings here.
*/
func swap(x, y, z string) (string, string, string) {
    return z, y, x
}

/*  This splits numbers and returns the product. It utilizes "naked" returns.
    It will return a value for x, y, and z.
*/
func split(prod int) (x, y, z int) {
    x = prod * 4/9
    y = prod - x
    z = prod - y
    return
}

func main() {
    a, b, c := swap("hello", "world", "fedora")
    fmt.Println(a, b, c)
    fmt.Println(split(20))
}
