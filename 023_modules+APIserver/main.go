//go get -u <lorem ipsum>

//The -u flag instructs get to update modules providing dependencies of packages named on the command line to use newer minor or patch releases when available.
//The -u=patch flag (not -u patch) also instructs get to update dependencies, but changes the default to select patch releases.
//When the -t and -u flags are used together, get will update test dependencies as well.

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to the chapter of go.mod file in Go language")
	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")

	http.ListenAndServe(":3000", router)
	greeter()

}

func greeter() {
	fmt.Println("Hello, greeter")
}

func serveHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Hello, welcome to the tutorial for Go language</h1>"))
}
