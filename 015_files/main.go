package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to the chapter of file manipulation in Go lang")

	myFile, err := os.Create("./mytextfile")
	// if err != nil {
	// 	panic(err)
	// }

	Checkerr(err)

	len, err := io.WriteString(myFile, "hello my name is Go and i am able to manipulate files with my functions")
	Checkerr(err)
	fmt.Println("The length of file is : ", len)

	readFile("./mytextfile")
	defer myFile.Close()

}

func readFile(filename string) {
	databytes, err := ioutil.ReadFile(filename)
	Checkerr(err)

	fmt.Println("The data in the file is : ", string(databytes))
}

func Checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
