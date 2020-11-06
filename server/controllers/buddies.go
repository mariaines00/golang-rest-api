package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mariaines00/golang-rest-api/models"
)

// AllBuddies displays all the buddies for a single robot
func AllBuddies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	b, err := models.AllBuddies(req)
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	j, _ := json.Marshal(b)
	w.Write(j)
}

// AddBuddy adds a single new buddy to the robot buddies list
func AddBuddy(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

// RemoveBuddy removes a single buddy from the buddies list
func RemoveBuddy(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}
