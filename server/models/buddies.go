package models

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mariaines00/golang-rest-api/config"
)

// AllBuddies displays all the buddies for a single robot
func AllBuddies(req *http.Request) ([]Robot, error) {
	b := make([]Robot, 0)

	name := mux.Vars(req)["name"]
	if name == "" {
		return b, errors.New("400. Bad Request")
	}

	_, err := GetRobotByName(name)
	if err != nil {
		return b, err
	}

	rows, err := config.DB.Query(`SELECT name
							FROM robots
							INNER JOIN friendships
							ON robots.id = friendships.buddy1
							WHERE robots.name = $1;`, name)
	if err != nil {
		return b, err
	}

	for rows.Next() {
		bot := Robot{}
		err := rows.Scan(&bot.ID, &bot.Name, &bot.Model)
		if err != nil {
			return nil, err
		}
		b = append(b, bot)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return b, nil
}

// AddBuddy adds a single new buddy to the robot buddies list
func AddBuddy(req *http.Request) (Robot, error) {
	b := Robot{}
	return b, nil
}

// RemoveBuddy removes a single buddy from the buddies list
func RemoveBuddy(req *http.Request) error {
	return nil
}
