package main

import "fmt"

//In Go lang we dont have concepts of inheritance and there's super, parent this all are the concepts of classes and in Go lang we don't have classes, we have structs
//In go lang functions that are used to operate on structs are termed as methods, methods are used to manipulate the struct data and we can perform operations on the datan of structs using methods.
func main() {
	fmt.Println("welcome to the chapter of Methods in Go")

	emp1 := User{"ramesh", "ramesh@cloudEQ.com", true, 18}
	// fmt.Println(emp1)

	//want all details of emp1
	// fmt.Printf("All details for emp1 %+v", emp1)

	//Only want name and email of emp1
	// fmt.Printf("\n Name of emp1 is - %v and the email of emp1 is - %v", emp1.Name, emp1.Email)

	// emp1.getVerified()
	// emp1.changeMail()
	//To test that we only modified in the copy of struct emp1, we can use print statement here
	fmt.Println("print email id : ", emp1.Email)
	// helloNum := 10
	// pointerNum := &helloNum
	// fmt.Println(pointerNum)
	emp1.ChangeRealEmail()
	//To test that we modified struct emp1's email using function emp1.ChangeRealEmail we can use print statement here
	fmt.Println("print email id : ", emp1.Email)

}

//Obviously we are gonna call the structure user from outside, and in Go lang if we want to make something public we have to keep its first letter capital, just like U in User and all in its attributes
type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}

//The fucntion to see if the emp1 is verified or not, if he/she is verified the function will return 'True' else function will return 'False'
func (emp1 User) getVerified() {
	fmt.Println("\n is user1 is verified ? : ", emp1.Verified)
}

//Whenever we pass the struct to any function(method) it sends a copy of struct to the function not the actual struct, so the manipulation can done only on the copy not the actual struct so there won't be any change in emp1 struct but in the copy of it.
func (emp1 User) ChangeEmail() {
	emp1.Email = "helloemp1@mail.com"
	fmt.Println("The email is updated in the copy of emp1 struct", emp1.Email)
}

//To really do manipulations in emp1's struct, you have to pass the address of the struct but not the copy of it using ampresand "*" infront of User struct we can make emp1 a pointer variable of type struct to pass the reference.
func (emp1 *User) ChangeRealEmail() {
	emp1.Email = "helloemp1@mail.com"
	fmt.Println("The email is updated in the copy of emp1 struct", emp1.Email)
}
