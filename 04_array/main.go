package main

import "fmt"

func main() {
	fmt.Println("welcome to the demonstration of arrays")

	var fruitList [5]string
	fruitList[0] = "apple"
	fruitList[1] = "orange"
	fruitList[2] = "kiwi"
	fruitList[4] = "guava"

	//In the O/P you will see that there's some more space in between 3rd and last element because the extra space indicates that there is also a fifth element that is not added into the array.
	fmt.Println(fruitList)

	//At last you can see the extra space that is there because extra space indicates tha last element.
	var vegList = [3]string{"tomato", "brinjal"}
	fmt.Println(vegList)

	//Different type to declare var.
	carList := [3]string{"honda", "suzuki"}
	fmt.Println(carList)

}
