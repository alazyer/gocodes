package main

import "C"

//export Sum
func Sum(a, b int) int {
    return a + b
}

func main() {}

// $ go build -buildmode=c-shared -o sum.so share_with_python.go

// $ python
// >>> import ctypes
// >>> lib = ctypes.CDLL('./sum.so')
// >>> lib.Sum(1, 2)  # return 3
