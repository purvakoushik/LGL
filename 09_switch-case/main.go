package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Roll-dice game using switch")

	//To initialize the first value with some random number, we usually do it with Time.Now() function and for more precise we use Time.Now().UnixNano()
	rand.Seed(time.Now().UnixNano())

	//rand.Intn(6) func is used to set the upper value till 6 to choose randomly, every time a random number from 1 to 6 will get into the diceNumber
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Your case is 1, open the acc.")
	case 2:
		fmt.Println("Your case is 2")
	//In case 3 that "Fallthrough" is used to run all other cases because 
	case 3:
		fmt.Println("Your case is 3")
		fallthrough
	case 4:
		fmt.Println("Your case is 4")
	case 5:
		fmt.Println("your case is 5")
	case 6:
		fmt.Println("Your case is 6, play again you got an extra chance")
	default:
		fmt.Println("what was that")
	}

}
