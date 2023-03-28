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
	fmt.Println("welcome to the chapter of JSON data creation in Go lang")
	encodeJson()
}

func encodeJson() {

	myCourses := []course{
		{
			Name:      "React JS",
			Course_id: 1,
			Platform:  "Youtube",
			Price:     10,
			Password:  "hgis525",
			Tags:      []string{"web-dev", "website"},
		},
		{
			Name:      "Python",
			Course_id: 2,
			Platform:  "FreeCodeCamp",
			Price:     20,
			Password:  "nccn528",
			Tags:      []string{"AI", "ML"},
		},
		{
			Name:      "HTML+CSS",
			Course_id: 3,
			Platform:  "Youtube",
			Price:     20,
			Password:  "jskdjk526",
			Tags:      nil,
		},
	}

	// finaJson, err := json.Marshal(myCourses)
	// checkerr(err)
	// fmt.Println(string(finaJson))

	indentFinalJson, err := json.MarshalIndent(myCourses, "", "\t")
	checkerr(err)
	fmt.Println(string(indentFinalJson))

	// fmt.Println(finaJson)

}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
