package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the lesson of fetch API response")
	getData()
}

func getData() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts/1"

	//Getting response pointer
	response, err := http.Get(myUrl)
	Checkerr(err)

	//Reading the body of response using the readall method, and it returns in the bytes format
	content, err := ioutil.ReadAll(response.Body)
	Checkerr(err)

	defer response.Body.Close()

	fmt.Printf("The Type of variable response is %T", response)
	fmt.Println(content)

	//Converting the bytes format which is in variable content to string.
	fmt.Println(string(content))
	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)

}

func usingStringBuilder() {
	const myUrl = "https://jsonplaceholder.typicode.com/posts/1"

	response, err := http.Get(myUrl)
	Checkerr(err)

	//A Builder is used to efficiently build a string using Write methods
	var buildString strings.Builder
	content, err := ioutil.ReadAll(response.Body)

	//This function write returns two values first one is the length of the content and second one is the error.
	//only this extra line has to be written for using string builder but its more convinient way in some cases.
	byteCount, err := buildString.Write(content)
	Checkerr(err)
	fmt.Println(byteCount)

	//This will build the content variable, which is in the form of bytes and write method helped to build a string.
	fmt.Println(buildString.String())

	//Closing the connection
	defer response.Body.Close()
}

func Checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
