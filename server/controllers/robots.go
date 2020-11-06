package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mariaines00/golang-rest-api/config"
	"github.com/mariaines00/golang-rest-api/models"
)

// Hello will just list all the existing endpoint and something else
func Hello(w http.ResponseWriter, req *http.Request) {
	config.TPL.ExecuteTemplate(w, "hello.gohtml", time.Now())
}

// AllRobots displays all robots
func AllRobots(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r, err := models.AllRobots()
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	j, _ := json.Marshal(r)
	w.Write(j)
}

// OneRobot shows info about a single robot
func OneRobot(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r, err := models.OneRobot(req)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, req)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	j, _ := json.Marshal(r)
	w.Write(j)
}

// CreateRobot creates a new robot
func CreateRobot(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r, err := models.CreateRobot(req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	j, _ := json.Marshal(r)
	w.Write(j)
}

// UpdateRobot updates information on a robot
func UpdateRobot(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r, err := models.UpdateRobot(req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}

	j, _ := json.Marshal(r)
	w.Write(j)
}

// RemoveRobot removes a robot
func RemoveRobot(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := models.RemoveRobot(req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//TODO: error handler
