package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"strings"

	"github.com/alazyer/gocodes/pkg"
	"github.com/alazyer/gocodes/tour"
)

func main() {
	// basic()
	checkHttp()
}

func checkHttp() {
	client := http.DefaultClient

	issureUrl := "http://94.191.119.70:31532/auth/realms/master"

	response, err := client.Get(fmt.Sprintf("%s/.well-known/openid-configuration", issureUrl))
	if err != nil {
		fmt.Println("err when request endpoints, err: ", err)
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(string(body))

	var endpoints struct {
		Issuer   string `json:"issuer"`
		AuthUrl  string `json:"authorization_endpoint"`
		TokenUrl string `json:"token_endpoint"`
	}
	err = json.Unmarshal(body, &endpoints)
	if err != nil {
		return
	}

	payloadStr := fmt.Sprintf("client_id=dex&client_secret=1a52ee75-9c63-47e6-afc4-513d6d44e59e&grant_type=client_credentials&response_type=token")
	response, err = client.Post(endpoints.TokenUrl, "application/x-www-form-urlencoded", strings.NewReader(payloadStr))
	if err != nil {
		return
	}
	fmt.Println(response.Status)
	body, _ = ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	if response.StatusCode != http.StatusOK {
		fmt.Println("insuccess client")
		return
	}

	payloadStr = fmt.Sprintf("client_id=dex&client_secret=1a52ee75-9c63-47e6-afc4-513d6d44e59e&username=admin&password=123456&grant_type=password&response_type=token")
	response, err = client.Post(endpoints.TokenUrl, "application/x-www-form-urlencoded", strings.NewReader(payloadStr))
	if err != nil {
		return
	}
	fmt.Println(response.Status)

	if response.StatusCode != http.StatusOK {
		fmt.Println("insuccess auth")
		return
	}
	fmt.Println("success auth")
}

func basic() {
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

	pkg.CommonStrings()

}
