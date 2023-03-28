package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to the pizza shop")
	fmt.Println("Please give us the rating")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for the rating", input)

	numInput, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		// panic(err)
		fmt.Println(err)
	} else {
		fmt.Println("we have added +1 to your rating", numInput+1)
	}

}
