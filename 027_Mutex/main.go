//Mutex stands for Mutual exclusion lock, which is basically responsible to lock the memory resource when one or more functions are in execution simultaneously in multi-threading.

// basically in this program, we are going to create a mutex and will use two functions from mutex which are provided by default i.e. -lock and -unclock

// In this scenario, the mutex is not looking very helpful but what is happening here is this that, what if every functional thread that goes into execution simultaneously like here we are only firing up one go routine there will be case where we have to fire 4 or 5 go routine and in that case every go routine is trying to write on same memory simultaneously that is not good, in that case mutex is very helpful to lock that memory resource for that particular go routine.

// Read write Mutex - its basically does that it will allow anyone to read from that memory and if anyone comes to write on that it will lock that memory space and if anyone comes in between to read it will kick it and doesn't allow to read untill the write operation is done.
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Waitgroup type to instruct the main func, when to wait for the parallel func, when to done that the parallel func is done with its execution and to when to add a parallel function that this func is out of main method for some executon wait till the response is back.
var wg sync.WaitGroup

var signals = []string{"test"} //string slice to store the signals.

var mut sync.Mutex //Mutex provide a lock over the memory, it locks the particular memory untill one go routine is working and it won't allow anybody or any other func execution to use that memory.

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

	fmt.Print("\n", signals) //Printing the signals that we got.

	//Printing out the tht total time exeuton for the main function.
	fmt.Println("\n", time.Since(start))
}

// Function accepting a parameter endpoint as string.
func getStatusCode(endpoint string) {
	defer wg.Done() //defer to execute at last in the function that this threaded func has done its execution, go back to main.

	res, err := http.Get(endpoint) //Calling the api with get method, saving response in res var.
	checkerr(err)                  //Checkerr func call

	mut.Lock() //Locking the memory space for writing on that.
	signals = append(signals, endpoint)
	mut.Unlock() //unlocking the same memory space after the execution.

	fmt.Printf("\n%d is the response code for %s", res.StatusCode, endpoint)

}

// func for checking the error.
func checkerr(err error) {
	if err != nil {
		fmt.Println("oops the error")
	}
}
