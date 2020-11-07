package main

import (
	"log"
	"net/http"

	"github.com/mariaines00/golang-rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)

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

	log.Println("Server started at port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))

	// TODO:
	// graceful shutdown
}

func index(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/hello", http.StatusSeeOther)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%v %v %v", req.Method, req.Host, req.RequestURI)
		next.ServeHTTP(w, req)
	})
}
