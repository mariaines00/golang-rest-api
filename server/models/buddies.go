package models

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mariaines00/golang-rest-api/config"
)

type Buddies []string

// AllBuddies displays all the buddies for a single robot
func AllBuddies(req *http.Request) (Buddies, error) {
	b := make(Buddies, 0)

	name := mux.Vars(req)["name"]
	if name == "" {
		return b, errors.New("400. Bad Request")
	}

	r, err := GetRobotByName(name)
	if err != nil {
		return b, err
	}

	rows, err := config.DB.Query(`SELECT name FROM robots JOIN friendships
							ON robots.id = friendships.buddy2 WHERE friendships.buddy1 = $1;`, r.ID)
	if err != nil {
		return b, err
	}

	for rows.Next() {
		n := ""
		err := rows.Scan(&n)
		if err != nil {
			return nil, err
		}
		b = append(b, n)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return b, nil
}

// AddBuddy adds a single new buddy to the robot buddies list
func AddBuddy(req *http.Request) error {
	name := mux.Vars(req)["name"]
	if name == "" {
		return errors.New("400. Bad Request")
	}

	b1, err := GetRobotByName(name)
	if err != nil {
		return err
	}

	buddy := req.FormValue("name")
	if buddy == "" {
		return errors.New("400. Bad Request")
	}

	// mate needs to exist ahah
	b2, err := GetRobotByName(buddy)
	if err != nil {
		return err
	}

	statement := fmt.Sprintf("INSERT INTO friendships (buddy1, buddy2) VALUES (%v, %v); INSERT INTO friendships (buddy2, buddy1) VALUES (%v, %v);", b1.ID, b2.ID, b2.ID, b1.ID)
	_, err = config.DB.Exec(statement)
	if err != nil {
		return errors.New("500. Internal Server Error " + err.Error())
	}

	return nil
}

// RemoveBuddy removes a single buddy from the buddies list
func RemoveBuddy(req *http.Request) error {
	name := mux.Vars(req)["name"]
	if name == "" {
		return errors.New("400. Bad Request")
	}

	_, err := GetRobotByName(name)
	if err != nil {
		return err
	}

	buddy := req.FormValue("name")
	if buddy == "" {
		return errors.New("400. Bad Request")
	}

	// mate needs to exist ahah
	b2, err := GetRobotByName(buddy)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("DELETE FROM friendships WHERE buddy2=$1;", b2.ID)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}

	return nil
}
