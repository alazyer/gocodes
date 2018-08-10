package main

import(
    "fmt"
    "math"
    "math/rand"

    "github.com/alazyer/gocodes/tour"
)


func main() {
    fmt.Printf("Hello, world.\n")

    x, y := rand.Intn(100), rand.Intn(100)
    sum := tour.Add(x, y)

    resultStr := fmt.Sprintf("Sum of %v and %v is %v", x, y, sum)

    fmt.Println(resultStr)
    fmt.Printf("Sum of %v and %v is %v\n", x, y, sum)

    var r float64 = 2.0
    area, area1 := tour.CircleArea(r)
    fmt.Printf("Area of circle with Pi: %g radius: %g is %g\n", math.Pi, r, area)
    fmt.Printf("Area of circle with Pi: %g radius: %g is %g\n", tour.Pi, r, area1)
    tour.BasicTypes()
}
