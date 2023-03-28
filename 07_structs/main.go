package main

import "fmt"

//In Go lang we dont have concepts of inheritance and there's super, parent this all are the concepts of classes and in Go lang we don't have classes, we have structs
func main() {
	fmt.Println("welcome to the chapter of structs in Go")

	emp1 := User{"ramesh", "ramesh@cloudEQ.com", true, 18}
	fmt.Println(emp1)

	//want all details of emp1
	fmt.Printf("All details for emp1 %+v", emp1)

	//Only want name and email of emp1
	fmt.Printf("\n Name of emp1 is - %v and the email of emp1 is - %v", emp1.Name, emp1.Email)

}

//Obviously we are gonna call the structure user from outside, and in Go lang if we want to make something public we have to keep its first letter capital, just like U in User and all in its attributes
type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}
