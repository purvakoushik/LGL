package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Go Slice")

	//under the hood slices are just like an array with some abstractions and more features

	//In Go slice we don't give the indexes, that how much we want to use, because slices change dynamically according to its size.
	var fruitList = []string{"apple", "mango", "banana"}
	fmt.Println("The fruit list is :", fruitList)

	//In slice if we want to add something, we have to use append method
	fruitList = append(fruitList, "peach", "guava", "grapes")
	fmt.Println("The fruit list is :", fruitList)

	//In slice if we want to cut the slice from between we can do that too
	//The index numbers are not inclusive so we get=== mango, banana
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	myNumbers := make([]int, 4)
	myNumbers[0] = 245
	myNumbers[1] = 144
	myNumbers[2] = 564
	myNumbers[3] = 585
	//if we will do this, then it will throw error of index out of range
	// myNumbers[4] = 654

	//The adding of value works oin this way in Go slices.
	//myNumbers = append(myNumbers, 555,654,985)

	fmt.Println(myNumbers)

	//This methods only works on slice, not on arrays
	sort.Ints(myNumbers)
	fmt.Println(myNumbers)
	fmt.Println(sort.IntsAreSorted(myNumbers))

}
