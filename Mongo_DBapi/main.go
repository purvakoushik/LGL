package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/purvakoushik/mongo_DB_API/router"
)

func main() {
	fmt.Println("welcome to the chapter of API mongo DB CRUD ")

	router := router.Routers()
	fmt.Println("The server is starting..........")

	fmt.Println("The server is listening at port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))

}
