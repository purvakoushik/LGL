package main

import "fmt"

func main() {

	fmt.Println("More on Go slice")

	fruitSlice := []string{"apple", "banana", "guava", "mango", "grapes"}
	fmt.Println("fruitSlice slice contains", fruitSlice)

	// fruitSlice = append(fruitSlice[1:2])
	// fmt.Println(fruitSlice)

	//The first arguement you give in append func, the first value is inclusive and the second value which we pass is non-inclusive
	//i.e. means its start from including 0 which is apple and then go till index-2 which is non-inclusive so [apple banana] now adding that
	// index is index+1 which means 3, its inclusive so mango and till end which is grapes so [mango grapes]
	// So [apple banana] + [mango grapes] = [apple banana mango grapes]
	var index int = 2
	fruitSlice = append(fruitSlice[:index], fruitSlice[index+1:]...)
	fmt.Println(fruitSlice)

}
