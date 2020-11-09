package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mariaines00/golang-rest-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	server := &http.Server{
		Addr:         "0.0.0.0:3000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

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

	go func() {
		log.Println("Server started at port 3000")
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	server.Shutdown(ctx)

	log.Println("see you later aligator")
	os.Exit(0)
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
