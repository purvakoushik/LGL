package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Its always a good option to declare the url globally
const url = "https://cloudeq-employee.w3spaces.com/"

func main() {
	fmt.Println("Welcome to the chapter of handling web request using Go lang")

	//The http.get() func returns two values, one is err and second one is the response in reference format.
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	//To chech the type of the response, we are getting
	fmt.Printf("The response is type is %T", response)

	//The response has so many fields and the body contains the data and it is in the bytes format, so we are storing it in databytes variable.
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	//Converting the response to string, to see the data
	fmt.Println("The response from the site is : \n", string(databytes))

	//Closing the connection
	defer response.Body.Close()
}
