package tour

import (
    "fmt"
    "math"
)

func Sqrt1(x float64) float64 {
    // always run 10 times
    var seed, positiveX float64 = 1.0, 0

    if x < 0 {
        positiveX = -x
    } else {
        positiveX = x
    }

    for i := 0; i < 10; i++ {
        seed -= (seed*seed - positiveX) / (2*seed)
        fmt.Println("Round:", i+1, " \tSeed:", seed)
    }

    return seed
}

func Sqrt2(x float64) float64 {
    // margin should be small than 0.0001
    var seed, positiveX float64 = 5.0, 0
    var i int = 1

    if x < 0 {
        positiveX = -x
    } else {
        positiveX = x
    }

    for math.Abs(seed * seed - positiveX) > 0.0001 {
        seed -= (seed*seed - positiveX) / (2*seed)
        fmt.Println("Round:", i, " \tSeed:", seed)
        i += 1
    }

    return seed
}

func Defers() {
    // each encountered defer func is push into stack in order
    // result: hello 2\n world 1\n ! 0
    i := 0
    defer fmt.Println("!", i)
    i += 1

    defer fmt.Println("world", i)
    i += 1

    fmt.Println("hello", i)
}
