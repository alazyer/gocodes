package main

import (
	"fmt"
)

type Own struct {
	Labels map[string]string `json:"labels"`
}

func main() {
	fmt.Println("Hello, playground")
	a := make([]interface{}, 0)
	a = append(a, 123)
	a = append(a, "123")
	fmt.Println(a)

}
