//This code will print the for loop statements in greeter function in the order of the funtion call, first the greeter function will be called with parameter "hello" and the second time it will be called with parameter "bye"

// package main

// import "fmt"

// func main() {

// 	greeter("hello")
// 	greeter("bye")
// }

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("This is  parallel thread for", s)
// 	}
// }

// This code will print the loop statements in greeter function using Go routines which is responsible for multi-threading.

package main

import (
	"fmt"
	"time"
)

func main() {
	//The main fnction is same as above, just a small difference the first function call now has a go statement before the function call, which indicates to run this function parallely.

	//If we run this code without waiting for the mult-threaded function code to return, it won't get us the result and the execution of main method will get completed and we will see the O/P only for the second function call.

	go greeter("hello") //Go is used here to make this function call run parallely
	greeter("bye")      //Normal function call to greeter function
}

func greeter(s string) {
	//For loop to execute the print statement multiple times.
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond) //Waiting for the parallel function call to finish the execution
		fmt.Println("This is  parallel thread for", s)
	}
}
