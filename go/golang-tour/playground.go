package main

import (
    "fmt"
)

// set types
var x, y, z int

// Infer types
var c, python, java = true, false, "no!"

var (
  p = Name{"Gabe", "B-W"}  // has type Name
  q = &Name{"Gabe", "B-W"} // has type *Name
  r = Name{First: "Gabe"} // Last = "" is implied
  s = Name{}  // both fields are ""
)

type Name struct {
  First string
  Last  string
}

func main() {
  array := []int{1,2,3,4}
  name := Name{"Gabe", "B-W"}
  other := &name
  other.First = "Pat"

  for i := 0; i < len(array); i++ {
    fmt.Println(array[i])
  }
  fmt.Println(array[1:3]) // like 1...3

  fmt.Println(name)
  fmt.Println(p, q, r, s)
}
