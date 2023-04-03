// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	//You can check the execution time of this main function it is taking approx. 7 seconds to execute the main function, now we use the Go routines and see if there is any differ in this main func and the main func with go routines.

// 	var start = time.Now()

// 	var endpoint = []string{
// 		"https://google.com",
// 		"https://lco.dev",
// 		"https://github.com",
// 		"https://fb.com",
// 		"https://go.dev",
// 	}

// 	for _, web := range endpoint {
// 		getStatusCode(web)
// 	}

// 	fmt.Println("\n", time.Since(start))
// }

// func getStatusCode(endpoint string) {
// 	res, err := http.Get(endpoint)
// 	checkerr(err)

// 	fmt.Printf("\n%d is the response code for %s", res.StatusCode, endpoint)

// }

// func checkerr(err error) {
// 	if err != nil {
// 		fmt.Println("oops the error")
// 	}
// }

// ************************************************************************************************************************************************
// The program to get the status codes of various endpoints using go routines and utilizing the multi threading.
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Waitgroup type to instruct the main func, when to wait for the parallel func, when to done that the parallel func is done with its execution and to when to add a parallel function that this func is out of main method for some executon wait till the response is back.
var wg sync.WaitGroup

func main() {

	//You can clearly see the difference time, the main func with go routines is taking only 1 second approx to execute, compare to previous func which was taking approx. 7 seconds to execute.

	//The current time saving into the var start
	var start = time.Now()

	//String slice of multiple endpoints.
	var endpoint = []string{
		"https://google.com",
		"https://lco.dev",
		"https://github.com",
		"https://fb.com",
		"https://go.dev",
	}

	//Loop to call the getStatusCode func for every endpoint present in the slice "endpoint".
	for _, web := range endpoint {
		go getStatusCode(web) //To call the function parallely with go routine
		//To inform the main that one func is out to execute parallely, everytime loop is running one func call is outside thats why 1 is passed as parameter in Add() func.
		wg.Add(1)
	}

	//To instruct main to wait for the threaded function to finish its execution.
	wg.Wait()

	//Printing out the tht total time exeuton for the main function.
	fmt.Println("\n", time.Since(start))
}

// Function accepting a parameter endpoint as string.
func getStatusCode(endpoint string) {
	defer wg.Done() //defer to execute at last in the function that this threaded func has done its execution, go back to main.

	res, err := http.Get(endpoint) //Calling the api with get method, saving response in res var.
	checkerr(err)                  //Checkerr func call

	fmt.Printf("\n%d is the response code for %s", res.StatusCode, endpoint)

}

// func for checking the error.
func checkerr(err error) {
	if err != nil {
		fmt.Println("oops the error")
	}
}
