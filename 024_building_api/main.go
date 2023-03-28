//In the bigger applications the whole backend is developed in chunks and different files and folders, so wherever we are seeing the "diff file" in this code, it means that function or code must be in diff file or folder

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model structure of course for DB - diff file
// author passed as pointer, because we have just declare the author not initialized it.
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

// Model structure of author - diff file
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Go slice acts as fake DB - diff file
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to the API building in Go language.")

	//Seeding the database with some initial data
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "React JS",
		CoursePrice: 399,
		Author:      &Author{Fullname: "purva koushik", Website: "YouTube"}})

	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "HTML",
		CoursePrice: 599,
		Author:      &Author{Fullname: "mark lewis", Website: "Meta"}})

	//New Router for handling the routes
	router := mux.NewRouter()

	//Handlers/controllers
	router.HandleFunc("/", serveHome).Methods("GET")                     //Checked
	router.HandleFunc("/courses", GetAllCourses).Methods("GET")          //Checked
	router.HandleFunc("/course/{id}", Getonecourse).Methods("GET")       //Checked
	router.HandleFunc("/course", Createonecourse).Methods("POST")        //Checked
	router.HandleFunc("/course/{id}", Updateonecourse).Methods("PUT")    //
	router.HandleFunc("/course/{id}", Deleteonecourse).Methods("DELETE") //Checked

	//Listening to the post 3000
	http.ListenAndServe(":3000", router)
}

// This function will serve the response as a get request
func serveHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>welcome to the API server build in Go language"))
}

// controller, request handlers - diff file
func GetAllCourses(res http.ResponseWriter, req *http.Request) {
	fmt.Println("You can get all your courses")
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(courses)
}

func Getonecourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Pass ID and get course corresponding to that course id")
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(res).Encode(course)
			return
		}
	}
	json.NewEncoder(res).Encode("No course found with the given ID")
	return
}

func Createonecourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("This func is to create a course")

	res.Header().Set("Content-Type", "application/json")

	//If body of req is empty, means in the req body no data is sended by the user then the reponse to be given, if data is nil means for these types "pointer types, map types, slice types, function types, channel types, interface types, and so on" nil can represent zero.
	if req.Body == nil {
		json.NewEncoder(res).Encode("No data found ERROR 404, please send some data")
	}

	//if data structure is empty like this "{}" - scenario for that

	//New slice is created of type "Course"
	var course Course
	//The coming req is decoded from json and inserted into that slice called "course" we created for that only we passes the reference of course to reflect the changes on real course slice and not on copy of that..
	_ = json.NewDecoder(req.Body).Decode(&course)

	//Calling the method IsEmpty that relates to struct course to check if the Coursename or courseId is present.
	if course.IsEmpty() {
		// json.NewEncoder(res).Encode("No CourseID and Coursename found ERROR 404, please send some data")
		json.NewEncoder(res).Encode("No Coursename found ERROR 404, please send some data")
	}

	//Generating random unique number for courseId
	//and appending the created 'course' object to course slice.

	//rand.seed initialize the time for the first time, it will seed no the basis of current time, and strconv.Itoa can convert any int to string, rand.Intn will take max value as 100.
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	//Appending the course object to courses slice.
	courses = append(courses, course)

	//sending course as response to the user
	json.NewEncoder(res).Encode(course)

}

func Updateonecourse(res http.ResponseWriter, req *http.Request) {
	//Can be improve by adding the functionality like, if we want to update only the particular fields of the data and the remain fields in the data will be go in the DB as it is, they were.

	fmt.Println("This func is to update a course by ID")
	//Setting the header as Content-Type to application/json to accept the json.
	res.Header().Set("Content-Type", "application/json")

	//To grab the courseID that user gonna pass, to update that particular course corresponding to that ID.
	params := mux.Vars(req)

	//Looping through all the course in the courses slice, by setting the range.
	for index, course := range courses {
		//if any course's courseID matches with the params ID that user passed then ->
		if course.CourseId == params["id"] {
			//-> then that course will get first removed from the courses slice
			courses = append(courses[:index], courses[index+1:]...)

			//-> then A new Course type slice is made
			var course Course

			//-> then decoding the req.body to json and updating that in newly made course slice.
			_ = json.NewDecoder(req.Body).Decode(&course)

			//-> then Changing the course ID with the same ID that is passed to controller for updation.
			course.CourseId = params["id"]

			//-> then appending that particular newly made course slcie to the courses
			courses = append(courses, course)

			//-> then response to the user, with params ID and updated course
			json.NewEncoder(res).Encode("The course is updated for the given course Id")

			//-> Returning the function
			return
		}
	}

}

func Deleteonecourse(res http.ResponseWriter, req *http.Request) {

	fmt.Println("Function to delete the course matching to the ID passed by the user")
	res.Header().Set("Content-Type", "application/json")

	//storing the req.body variables in params to fetch id from that to delete that course corresponding to that ID.
	params := mux.Vars(req)

	//looping through all the courses in courses slice -> if course ID matches simply remove that course object from the slice -> send the json response of successfull deletion -> break the loop.
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(res).Encode("The course is deleted successfully for the given courseID")
			break
		}
	}
}
