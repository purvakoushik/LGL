package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to the chapter of PostForm Req to the server")
	postFormReqApi()

}

func postFormReqApi() {
	const myUrl = "http://localhost:8000/postform"

	postData := url.Values{}

	postData.Add("firstname", "Purva")
	postData.Add("lastname", "Koushik")
	postData.Add("email", "purva@lco.com")

	response, err := http.PostForm(myUrl, postData)
	checkerr(err)

	dataBytes, err := ioutil.ReadAll(response.Body)
	fmt.Println("The response from the server for the PostForm data is", string(dataBytes))
	defer response.Body.Close()
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
