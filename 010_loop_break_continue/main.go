package main

import "fmt"

func main() {
	fmt.Println("Welcome to the chapter of loops, break and continue statements")

	fruitList := []string{"apple", "guava", "watermelon", "grapes", "kiwi", "dragonFruit", "orange"}
	// fmt.Println(fruitList)

	//This is one of the way, you can use for loops in Go, and for using loops in Go we have only one keyword "For" to use every variation of loops
	// for i := 0; i < len(fruitList); i++ {
	// 	fmt.Println(fruitList[i])
	// }

	//Using range keywords in loops
	// for i := range fruitList {
	// 	fmt.Println(fruitList[i])
	// }

	//Key, value in for loops
	// for key, value := range fruitList {
	// 	fmt.Printf("\n on the key %v, the value is %v", key, value)
	// }

	//If we are not concerned about the key and only value in for we can use _ in the place of key, comma Ok syntax
	for _, value := range fruitList {
		fmt.Printf("\n on the key, the value is %v", value)
	}

	//using loop like while
	// interValue := 1
	// for interValue < 10 {
	// 	fmt.Println("\n", interValue)
	// 	interValue++
	// }

	//break statement
	// breakValue := 1
	// for breakValue < 10 {
	// 	if breakValue == 5 {
	// 		break
	// 	}

	// 	fmt.Println("\n", breakValue)
	// 	breakValue++
	// }

	//Continue statement
	// contValue := 1
	// for contValue < 10 {
	// 	if contValue == 5 {
	// 		contValue++
	// 		continue
	// 	}

	// 	fmt.Println("\n", contValue)
	// 	contValue++
	// }

	//Goto statement
	// 	gotoValue := 1
	// 	for gotoValue < 10 {
	// 		if gotoValue == 5 {
	// 			goto tempStatement
	// 		}

	// 		fmt.Println("\n", gotoValue)
	// 		gotoValue++
	// 	}

	// tempStatement:
	// 	fmt.Println("welcome to loops again")
}
