package main

import "fmt"
import "math"

// "const" can appear anywhere "var" can
const s string = "constant"

func main() {
    fmt.Println(s)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    // a numeric constant has no type until it's given one (here, it's cast)
    fmt.Println(int64(d))

    // give   n a type by using it in a context that requires one
    fmt.Println(math.Sin(n))
}
