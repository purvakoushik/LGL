package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name      string `json:"coursename"`
	Course_id int
	Platform  string `json:"website"`
	Price     int
	Password  string   `json:"-"`              //This "-" dash simply means don't reflect this field whosoever uses my api.
	Tags      []string `json:"tags,omitempty"` //omitempty simply says if value is null simply don't reflect it.
}

func main() {
	fmt.Println("Welcome to the New chapter of consuming Json data in Go language.")

	decode_Json_2_Put_Values_In_Struct()
	// decode_Json_2_Put_Values_In_Key_Value()
	// myMap()
}

func decode_Json_2_Put_Values_In_Struct() {
	myOnlineData := []byte(`
	{
		"coursename": "Python",
		"Course_id": 2,
		"website": "FreeCodeCamp",
		"password": "jdhud1256",
		"Price": 20,
		"tags": ["AI","ML"]
	}
	`)

	var myCourse course

	isJson := json.Valid(myOnlineData)

	if isJson {
		fmt.Println("The data is a valid json")
		json.Unmarshal(myOnlineData, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("The data is not a valid json")
	}

	fmt.Println(myCourse.Name)
	fmt.Println(myCourse.Course_id)
	fmt.Println(myCourse.Platform)
	fmt.Println(myCourse.Price)
	fmt.Println(myCourse.Password)
	fmt.Println(myCourse.Tags)
}

// In some cases you just want to add json data to the key value pair for looping through it or access it later form the map.
func decode_Json_2_Put_Values_In_Key_Value() {
	myOnlineData := []byte(`
	{
		"coursename": "Python",
		"Course_id": 2,
		"website": "FreeCodeCamp",
		"Price": 20,
		"password": "jdhud1256",
		"tags": ["AI","ML"]
	}
	`)

	var myMapforJson map[string]interface{}
	json.Unmarshal(myOnlineData, &myMapforJson)

	for key, value := range myMapforJson {
		fmt.Printf("The key is: %v, the value is: %v \n", key, value)
	}
}

func myMap() {
	myMap2 := make(map[string]string)
	myMap2["name"] = "purva"
	myMap2["lastname"] = "koushik"

	for key, value := range myMap2 {
		fmt.Printf("The key %v represents %v value correspondingly\n", key, value)
	}
}
