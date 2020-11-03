package controllers

import (
	"fmt"
	"net/http"
	"time"

	"../config"
)

// Hello will just list all the existing endpoint and something else
func Hello(w http.ResponseWriter, req *http.Request) {
	config.TPL.ExecuteTemplate(w, "hello.gohtml", time.Now())
}

// AllRobots displays all robots
func AllRobots(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>AllRobots</h1>")
}

// CreateRobot creates a new robot
func CreateRobot(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>CreateRobot</h1>")
}

// OneRobot shows info about a single robot
func OneRobot(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>OneRobot</h1>")
}

// UpdateRobot updates information on a robot
func UpdateRobot(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>UpdateRobot</h1>")
}

// RemoveRobot removes a robot
func RemoveRobot(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>RemoveRobot</h1>")
}
