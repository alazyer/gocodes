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

    fmt.Println("Sqrt of 10 with 10 time loop is: ", tour.Sqrt1(10))
    fmt.Println()
    fmt.Println("Sqrt of 10 with margin smaller than 0.0001 is: ", tour.Sqrt2(10))

    fmt.Println()
    tour.Defers()

    fmt.Println()
    fmt.Printf("Two dimension array of dx: %d, dy: %d is: %v\n", 2, 5, tour.Arrays(2, 5))

    fmt.Println()
    fmt.Println("First 10 numbers in Fibnacci array")
    fibonacci := tour.Fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(fibonacci())
    }

}
