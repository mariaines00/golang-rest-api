package models

import (
	"errors"
	"net/http"
	"strconv"

	"../config"
	"github.com/gorilla/mux"
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
	bot := Robot{}

	id := mux.Vars(req)["id"]

	if id == "" {
		return bot, errors.New("400. Bad Request")
	}

	row := config.DB.QueryRow("SELECT * FROM robots WHERE id = $1", id)

	err := row.Scan(&bot.ID, &bot.Name, &bot.Model, []string{})
	if err != nil {
		return bot, err
	}

	return bot, nil
}

// CreateRobot creates a single robot
func CreateRobot(req *http.Request) (Robot, error) {
	// get form values
	bot := Robot{}
	bot.Name = req.FormValue("name")
	m := req.FormValue("model")

	// validate form values
	if bot.Name == "" || m == "" {
		return bot, errors.New("400. Bad request. All fields must be complete")
	}

	// convert form values
	f64, err := strconv.ParseFloat(m, 32)
	if err != nil {
		return bot, errors.New("406. Not Acceptable. Model must be a number")
	}
	bot.Model = float32(f64)

	// insert values
	_, err = config.DB.Exec("INSERT INTO robots (name, model) VALUES ($1, $2)", bot.Name, bot.Model)
	if err != nil {
		return bot, errors.New("500. Internal Server Error" + err.Error())
	}
	return bot, nil
}

// UpdateRobot updates data about a single robot given the ID
func UpdateRobot(req *http.Request) (Robot, error) {
	// get form values
	bot := Robot{}
	bot.Name = req.FormValue("name")
	m := req.FormValue("model")

	id := mux.Vars(req)["id"]
	if id == "" {
		return bot, errors.New("400. Bad Request")
	}

	// validate form values
	if bot.Name == "" || m == "" {
		return bot, errors.New("400. Bad request. All fields must be complete")
	}

	// convert form values
	f64, err := strconv.ParseFloat(m, 32)
	if err != nil {
		return bot, errors.New("406. Not Acceptable. Model must be a number")
	}
	bot.Model = float32(f64)

	// insert values
	_, err = config.DB.Exec("UPDATE robots SET name = $1, model=$2, WHERE id=$1;", id, bot.Name, bot.Model)
	if err != nil {
		return bot, err
	}

	return bot, nil
}

// RemoveRobot removes a single robot given the ID
func RemoveRobot(req *http.Request) error {
	id := mux.Vars(req)["id"]
	if id == "" {
		return errors.New("400. Bad Request")
	}

	_, err := config.DB.Exec("DELETE FROM robots WHERE id=$1;", id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil

}

// ParseRobotRequestValues will look at the form and return an Robot struct
func ParseRobotRequestValues(req *http.Request) (Robot, error) {
	// get form values
	bot := Robot{}
	bot.Name = req.FormValue("name")
	m := req.FormValue("model")

	// validate form values
	if bot.Name == "" || m == "" {
		return bot, errors.New("400. Bad request. All fields must be complete")
	}

	// convert form values
	f64, err := strconv.ParseFloat(m, 32)
	if err != nil {
		return bot, errors.New("406. Not Acceptable. Model must be a number")
	}
	bot.Model = float32(f64)

	return bot, nil
}

// ParseRobotID returns the id string from the request
func ParseRobotID(req *http.Request) (string, error) {
	id := mux.Vars(req)["id"]
	if id == "" {
		return "", errors.New("400. Bad Request")
	}
	return id, nil
}
