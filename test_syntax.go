package main

import (
  "fmt"
)

func main() {
  x := 1    // Shorthand variable assignment.
  y := 4    // Shorthand variable assignment

  fmt.Println(x, "+", y, "=", x+y)

  var a1 [4]int             // An array of 4 ints, initialized to all 0.
  // var a1 [x]int             // An array of 4 ints, initialized to all 0.
  // var a1 [x][][][]int             // An array of 4 ints, initialized to all 0.
  a2 := make([]int, x)      // An array of `x` ints
  a3 := make([]int, x, y)      // An array of `x` ints
  a4 := []int{1,2,3}         // An array of 4 ints, initialized to all 0.

  fmt.Println("a1 has size of: ", len(a1), "and capacitiy: ", cap(a1))
  fmt.Println("a2 has size of: ", len(a2), "and capacitiy: ", cap(a2))
  fmt.Println("a3 has size of: ", len(a3), "and capacitiy: ", cap(a3))
  fmt.Println("a4 has size of: ", len(a4), "and capacitiy: ", cap(a4))

}
