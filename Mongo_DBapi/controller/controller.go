package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/purvakoushik/mongo_DB_API/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://gopurva:N6WZF77YMbQlwwpW@mygodb.xkjm4ek.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// V. important,
var collection *mongo.Collection

// func to connect with mongo DB
// Init is a special method that runs, when initially when the code starts to execute.
func init() {
	clientOption := options.Client().ApplyURI(connectionString)

	//Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DataBase connected successfully")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
	fmt.Println("Collection instance is ready")

}

// MongoDB helpers - diff file - our controllers will use these helpers to do CRUD operations in MongoDB.

// Everything inside the MongoDB is not json but technically its bson and they look exactly the same but bson provides more things like this ID as _id, these things are provided by bson

// Insert one movie
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	checkerr(err)

	//if inserted successfully then print
	fmt.Println("The movie inserted successfully", inserted.InsertedID)
}

// Function to update the movie field by matching the ID given to the function with _id that is present in the DB.
// Update one movie record.
func updateOneMovie(movieId string) {

	// This function will convert the objectID in the form of string to the valid objectID that mongoDB accepts.
	id, _ := primitive.ObjectIDFromHex(movieId)

	//Finding/filtering the movie on the basis of the id that is passed to the function.
	filter := bson.M{"_id": id}

	//updating the field of that particular id, with the help of bson.M function, and setting the field to watched == true(diff between bson.D and bson.M)
	update := bson.M{"$set": bson.M{"watched": true}}

	//This function will probably return an error and in result the count of the odified/updated fields.
	result, err := collection.UpdateOne(context.Background(), filter, update)
	checkerr(err)

	//here the result stores the count of updates fields, that hw many fields are being updated.
	fmt.Println("The modified count is : ", result.ModifiedCount)
}

// Delete one movie record from DB
func deleteOneMovie(movieId string) {
	// This function will convert the objectID in the form of string to the valid objectID that mongoDB accepts.
	id, _ := primitive.ObjectIDFromHex(movieId)

	//Finding/filtering the movie on the basis of the id that is passed to the function.
	filter := bson.M{"_id": id}

	//delete that one movie record correspond to the id passed to the function.
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	checkerr(err)

	fmt.Println("The delete count is : ", deleteCount)
}

// function to delete many movies
func deleteAllMovie() int {

	//filter left empty without any key or value to select all the records present in DB.
	filter := bson.M{} //--> bson.M{} -- The structure of M type -- which is used when order of the elements doesn't matter.

	//--> bson.D{{}} -- The structure of D type -- which is used when order of the elements matter.

	//In DeleteResult, its not like only the count, its more like like key-value pair, so we have to dig more deep to find
	// deleteResult, err := collection.DeleteMany(context.Background(), bson.M{})  //--> This way is also correct.
	deleteResult, err := collection.DeleteMany(context.Background(), filter)
	checkerr(err)

	fmt.Println("The count of total deleted movie is : ", deleteResult.DeletedCount)
	fmt.Printf("\n %T", deleteResult.DeletedCount)
	return int(deleteResult.DeletedCount)
}

// Function to get all the movies from DB
// This function getting all movies from the DataBase and returning the array of the movies of primitive.M type where "Primitive" is a package available in Go for BSON types and ".M" represents the unordered type of the movies/elements present in the DB
func getAllMovie() []primitive.M {
	fmt.Println("hello all movies from mongDB helper")
	//The ordered way of elements bson.D to use without any parameters/ all values selected.
	filter := bson.D{}

	//Returning an error and cursor -->(cursor)-which is a gigantic object coming from mongoDB with whole lots of data, looping through it will give us the actual data we need.
	cursor, err := collection.Find(context.Background(), filter)
	checkerr(err)

	//Array of movies of primitive.M type for bson.M type. //unordered //Both are same but we are using the primitive package that is why we r utilizing it.
	var movies []primitive.M

	//Looping through the object cursor with the help of Next function which accepts a context.
	for cursor.Next(context.Background()) {

		//movie var of bson.M --> remember to be consistent from here onwards you cant use ordered D type above in array and unordered M type in the loop.
		var movie bson.M

		//If no error while decoding the object cursor for the data, paas the reference of var movie to store that data for the movie in the movie var.
		err := cursor.Decode(&movie)
		checkerr(err)                  //check error.
		movies = append(movies, movie) //appending the movie var of bson.M type to the array/slice movies of primitive.M type.
	}
	defer cursor.Close(context.Background()) //Closing the connecttion after getting the data.
	return movies                            //return the array
}

// Function to check for the error.
func checkerr(err error) {
	log.Fatal(err)
}

//*********************************************************Actual Controllers -- diff file*********************************************************

// function to serve at the root / path.
func ServeHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to the MnongoDB API for CRUD operations."))
}

// Controller function to get all the movies from the DB
func GetAllMovie(res http.ResponseWriter, req *http.Request) {
	println("hello all movies from controller")
	res.Header().Set("Content-Type", "application/json") //Setting the response header and its accepting the application/json

	movies := getAllMovie()             //getting all movies by calling the function, its returning an array of primitive.M type
	json.NewEncoder(res).Encode(movies) //encoding the response to json and sending it to the user hitting the endpoint.

}

// Controller function to get the user json request for creating  a movie.
func CreateOneMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")    //setting the header to accept application/json type of data
	res.Header().Set("Allow-Control-Allow-Methods", "POST") //setting the header to accept only the post method, coming from the user when user is hitting the endpoint.

	var movie model.Netflix                      //Creating a var movie of type Netflix struct that contains all the properties of Netflix struct.
	_ = json.NewDecoder(req.Body).Decode(&movie) //Decoding the user req json and passing the reference of movie to store that into the movie var.
	insertOneMovie(movie)                        //Calling the MongoDB helper function to insert the movie into the DB
	json.NewEncoder(res).Encode(movie)           //Encoding the movie var which is of Netflix struct to json and sending it to the user as a response.
}

// Controller function to update the watched field of Netflix struct or udating the movie.
func MarkAsWatched(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")    //setting the header to accept application/json type of data
	res.Header().Set("Allow-Control-Allow-Methods", "POST") //setting the header to accept only the post method, coming from the user when user is hitting the endpoint.

	params := mux.Vars(req) // Calling the mux.Vars function, which is accepting the req as a parameter, from which we can destructure the required properties/data.

	updateOneMovie(params["id"])              // Calling the MongoDB helper function to update the field in the DB, which is accepting ID of the movie as a parameter.
	json.NewEncoder(res).Encode(params["id"]) //Sending the ID itself as a response to the user back.
}

func DeleteOneMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")      //setting the header to accept application/json type of data
	res.Header().Set("Allow-Control-Allow-Methods", "DELETE") //setting the header to accept only the post method, coming from the user when user is hitting the endpoint.

	params := mux.Vars(req)

	deleteOneMovie(params["id"])
	json.NewEncoder(res).Encode(params["id"])
}

func DeleteAllMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")      //setting the header to accept application/json type of data
	res.Header().Set("Allow-Control-Allow-Methods", "DELETE") //setting the header to accept only the post method, coming from the user when user is hitting endpoint.

	result := deleteAllMovie()
	json.NewEncoder(res).Encode(result)

}
