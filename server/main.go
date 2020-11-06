package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mariaines00/golang-rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/hello", controllers.Hello)

	r.HandleFunc("/robots", controllers.AllRobots).Methods("GET")
	r.HandleFunc("/robots", controllers.CreateRobot).Methods("POST")

	r.HandleFunc("/robots/{name}", controllers.OneRobot).Methods("GET")
	r.HandleFunc("/robots/{name}", controllers.UpdateRobot).Methods("PUT")
	r.HandleFunc("/robots/{name}", controllers.RemoveRobot).Methods("DELETE")

	r.HandleFunc("/robots/{name}/buddies", controllers.AllBuddies).Methods("GET")
	r.HandleFunc("/robots/{name}/buddies", controllers.AddBuddy).Methods("PUT")
	r.HandleFunc("/robots/{name}/buddies", controllers.RemoveBuddy).Methods("DELETE")

	fmt.Println("Server started at port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

	// TODO:
	// r.Use(loggingMiddleware)
	// graceful shutdown
}

func index(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/hello", http.StatusSeeOther)
}
