package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello cloudEQ")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the CloudEQ: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf("Thank you for giving rating of %s",input)
}
