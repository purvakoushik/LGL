package main

import "fmt"

//The main function will always run by default without any caller statement whenever we run the program.
func main() {
	fmt.Println("welcome to the chapter of Functions")

	//calling the greeter function, caller statement
	greeter()

	//*****storing the result of adder function in variable result
	// result := adder(10, 20)
	// fmt.Println("The adder function returns : ", result)

	//*****Printing the return result of proAdder function, which can take any no. of inputs utilizing Go slice.
	//*****Calling the function in print statement itself
	// fmt.Println("The result of function proAdder is : ", proAdder(1, 2, 3, 4))

}

//decalaration and initialization of greeter function
func greeter() {
	fmt.Println("hello world, How are you this is a response from greeter")
}

//*****adder function, which will add two int inputs and will return the sum of it in int.
// func adder(valOne int, valTwo int) int {
// 	return valOne + valTwo
// }

//*****proAdder function, which can take any number of inputs and can add them return the result in int form.
// func proAdder(values ...int) int {
// 	sum := 0
//*****key will save the index while , value will save the data on that index
// 	for key, value := range values {
// 		sum += value
// 		fmt.Println("The value of the key is : ", key, " and the value is : ", value)
// 	}
// 	return sum
// }
