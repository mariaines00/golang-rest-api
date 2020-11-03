package main

import (
	"fmt"
	"log"
	"net/http"

	"./controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/hello", controllers.Hello)

	r.HandleFunc("/robots", controllers.AllRobots).Methods("GET")
	r.HandleFunc("/robots", controllers.CreateRobot).Methods("POST")

	r.HandleFunc("/robots/{id}", controllers.OneRobot).Methods("GET")
	r.HandleFunc("/robots/{id}", controllers.UpdateRobot).Methods("PUT")
	r.HandleFunc("/robots/{id}", controllers.RemoveRobot).Methods("DELETE")

	fmt.Println("Server started at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

	//TODO:
	/*
		r.HandleFunc("/robots/{id}/buddies", controllers.).Methods("GET")
		r.HandleFunc("/robots/{id}/buddies", controllers.).Methods("PUT")
		r.HandleFunc("/robots/{id}/buddies", controllers.).Methods("DELETE")
	*/
	// r.Use(loggingMiddleware)
	// graceful shutdown
}

func index(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/hello", http.StatusSeeOther)
}
