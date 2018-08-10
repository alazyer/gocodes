package tour

import (
    "fmt"
    "math"
)

const Pi = 3.14

// Exported names
func Add(x, y int) int {
    return x + y
}

// Named return value
func CircleArea(r float64) (area, area1 float64) {
    area = math.Pi * r * r
    area1 = Pi * r * r
    return
}

func BasicTypes() {
    var isGolang bool = true
    var str string = "I love China"
    var maxInt8, minInt8 int8 = 1 << 7 -1, -1 << 7
    var maxUint8, minUint8 uint8 = 1 << 8 -1, 0
    var aByte byte = 'a'
    var maxRune rune = 1 << 31 -1

    fmt.Printf("isGolang is a %T with value %v\n", isGolang, isGolang)
    fmt.Printf("str is a %T with value %v\n", str, str)
    fmt.Printf("maxInt8 is a %T with value %v\n", maxInt8, maxInt8)
    fmt.Printf("minInt8 is a %T with value %v\n", minInt8, minInt8)
    fmt.Printf("maxUint8 is a %T with value %v\n", maxUint8, maxUint8)
    fmt.Printf("minUint8 is a %T with value %v\n", minUint8, minUint8)
    fmt.Printf("aByte is a %T with value %v\n", aByte, aByte)
    fmt.Printf("maxRune is a %T with value %v\n", maxRune, maxRune)
}
