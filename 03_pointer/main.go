package main

import "fmt"

func main() {

	fmt.Println("welcome to the class of pointers")

	//The pointer variable is usualy made to store the address of the value that is stored in the memory
	//we can create a pointer variable like this : --->"var ptr *int"<--- : This pointer is of int type, however "var ptr = &myNumber" can figure out on its own that ptr is a pointer variable cause of lexer

	//whenever we r passing the value like this &myNumber, the address of myNumber goes to the ptr but not value.
	//Can access the value stored in memory by using *ptr.
	myNumber := 255
	var ptr = &myNumber

	fmt.Println("The value stored in ptr is - ", *ptr)
	fmt.Println("The address ptr points to is - ", ptr)

	//Using walrus operator to explicitly declare variable.
	helloPointer := &myNumber
	fmt.Println("The address stored in helloPointer is", helloPointer)

}
