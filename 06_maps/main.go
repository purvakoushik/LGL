package main

import "fmt"

func main() {
	fmt.Println("Welcome to the chapter of maps in Go lang")
	//In some languages thay are called hash tables and simple pretty key-value pairs in Go lang.

	languages := make(map[string]string)
	languages["JS"] = "javascript"
	languages["RB"] = "Ruby"
	languages["Go"] = "Go_language"

	fmt.Println("JS shorts for", languages["JS"])
	fmt.Println(languages)

	//if we want to delete a key and value from the map
	delete(languages, "RB")

	//Going through the map using loop
	for key, value := range languages {
		fmt.Printf("For key %v value is %v, \n", key, value)
	}

	//in this loop we don't care about the value only the key
	for key, _ := range languages {
		fmt.Printf("For key %v, \n", key)
	}

	//This is returning two values first one is the found in which the value javascript for respetive key JS will get stored and in bool the if the value there or not kinda bool value gonna store either true or false, means key exists or not.
	found, bool := languages["JS"]
	fmt.Println(found, bool)
	// var anotherMap map[string]string
	// anotherMap["car"] = "volvo"
	// anotherMap["truck"] = "eicher"
	// anotherMap["cycle"] = "atlas"

	// fmt.Println(anotherMap)
}
