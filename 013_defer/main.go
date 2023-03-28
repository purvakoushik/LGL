package main

import "fmt"

func main() {
	//Defer key is usually used to make a statement execute at last in a function, simply reverse order, means in a function on whichever statement the defer keyword is mentioned it will execute at last in the function.
	//It works like Last In First Out LIFO manner, if there are too many multiple defer statements in a fuunction the last line which goes in with defer statement will execute first
	fmt.Println("Hello, Welcome to the chapter of Defer keyword in Go.")

	//According to the function execution the execution should be like hello world but due to 'defer' keyword print line hello will go first in stack and then print line world goes to stack so according to LIFO the O/P will world hello
	defer fmt.Print("Hello")
	defer fmt.Print("world")
	nLoop()
}

//Due to the defer used in print statement of loop every print statement will go to the stack and will execute in LIFO manner, so O/P will be like 4,3,2,1,0
func nLoop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
