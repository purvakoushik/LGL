package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dfdjf154364"

func main() {
	fmt.Println("Welcome to the chapter of url handling in Go")

	//Parse function, parses the myUrl into chunks and we can do then manipulations on different parts of the URL
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//result.Query() makes the rawQuery in a better format of map (Key-value pair)
	qparams := result.Query()
	fmt.Printf("The type of it is : %T\n", qparams)
	fmt.Println(qparams["coursename"])

	//url.URL() function can make a url from parts individually created and modified
	partsOfUrl := url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		// Path:    "/learn",
		RawPath: "user=hitesh",
	}

	fmt.Println(partsOfUrl.String())
}
