package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the chapter of POST req to the api in Go lang")

	//Function caller for the function PostReqApi()
	PostReqApi()
}

func PostReqApi() {

	//The local server URL
	const myUrl = "http://localhost:8000/post"

	//Req body for getting the same response as the data we are sending to the server.
	reqBody := strings.NewReader(`
	{
		    "name": "John Doe",
            "email": "nnheo@example.com",
            "phone": "+1-555-1212",
            "address": "123 Main Street",
            "city": "San Francisco",
            "state": "CA",
            "country": "US"
	}`)

	//Returns a reference of the reader
	fmt.Println(reqBody)

	//Send the request to the server with url, Content-Type and the body data
	content, err := http.Post(myUrl, "application/json", reqBody)

	//ioutil.ReadAll() will read the reference of the response fetched from the server and will store in dataBytes in the form of Bytes.
	dataBytes, err := ioutil.ReadAll(content.Body)
	Checkerr(err)
	fmt.Println(dataBytes)

	//Converting the response to string
	fmt.Println("The response from the server is : \n", string(dataBytes))
}

func Checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
