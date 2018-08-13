package tour

func Fibonacci() func() int {
    numbers := make([]int, 0)
    i := 0

    return func() int {
        if i == 0 {
            numbers = append(numbers, 0)
        } else if i == 1 {
            numbers = append(numbers, 1)
        } else {
            numbers = append(numbers, numbers[i-1]+numbers[i-2])
        }

        i += 1

        return numbers[i-1]
    }

}

func Arrays(dx, dy int) [][]uint8 {
    result := make([][]uint8, dy)

    for i := range(result) {
        result[i] = make([]uint8, dx)
        for j := range(result[i]) {
            result[i][j] = uint8((i+j)/2)
        }
    }

    return result
}
