package models

import (
	"net/http"

	"../config"
)

// Robot struct
type Robot struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Model   float32  `json:"model"`
	Buddies []string `json:"buddies"`
}

// AllRobots returns all the robots
func AllRobots() ([]Robot, error) {
	rows, err := config.DB.Query("SELECT * FROM robots")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	b := make([]Robot, 0)
	for rows.Next() {
		bot := Robot{}
		err := rows.Scan(&bot.ID, &bot.Name, &bot.Model, []string{})
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

// OneRobot returns a single robot given the ID
func OneRobot(req *http.Request) (Robot, error) {
	//TODO:
	b := Robot{}

	return b, nil
}

// CreateRobot creates a single robot
func CreateRobot(req *http.Request) (Robot, error) {
	//TODO:
	b := Robot{}

	return b, nil
}

// UpdateRobot updates data about a single robot given the ID
func UpdateRobot(req *http.Request) (Robot, error) {
	//TODO:
	b := Robot{}

	return b, nil
}

// RemoveRobot removes a single robot given the ID
func RemoveRobot(req *http.Request) (Robot, error) {
	//TODO:
	b := Robot{}

	return b, nil
}
