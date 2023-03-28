package main

import "fmt"

func main() {
	fmt.Println("welcome to the chapter of If-else")

	//The simple syntax of if-else
	var myNumber = 55
	if myNumber < 45 {
		fmt.Println("myNumber is less than 45")
	} else if myNumber > 45 {
		fmt.Println("myNumber is greater than 45")
	}

	//The simple syntax of if-else and else-if
	var anotherNumber = 10
	if anotherNumber > 10 {
		fmt.Println("anotherNumber is greater than 10")
	} else if anotherNumber < 10 {
		fmt.Println("anotherNumber is less than 10")
	} else {
		fmt.Println("anotherNumber is exactly 10")
	}

	//Declaring and initializing the variable at sametime and checking the value of it's in the if-else condition
	if urNum := 10; urNum > 10 {
		fmt.Println("urnum is greater than 10")
	} else {
		fmt.Println("urnum is exactly 10")
	}

}
