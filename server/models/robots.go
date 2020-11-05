package models

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"../config"
	"github.com/gorilla/mux"
)

// Robot struct
type Robot struct {
	ID      string   `json:"-"`
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

// OneRobot returns a single robot given the Name
func OneRobot(req *http.Request) (Robot, error) {
	bot := Robot{}

	name := mux.Vars(req)["name"]
	if name == "" {
		return bot, errors.New("400. Bad Request")
	}

	row := config.DB.QueryRow("SELECT * FROM robots WHERE name = $1", name)
	err := row.Scan(&bot.ID, &bot.Name, &bot.Model)
	if err != nil {
		return bot, err
	}

	return bot, nil
}

// CreateRobot creates a single robot
func CreateRobot(req *http.Request) (Robot, error) {
	bot := Robot{}

	err := json.NewDecoder(req.Body).Decode(&bot)
	if err != nil {
		return bot, errors.New("400. Bad request. All fields must be complete")
	}

	// validate form values
	if bot.Name == "" || bot.Model == 0 {
		return bot, errors.New("400. Bad request. All fields must be complete")
	}

	// insert values
	_, err = config.DB.Exec("INSERT INTO robots (name, model) VALUES ($1, $2)", bot.Name, bot.Model)
	if err != nil {
		return bot, errors.New("500. Internal Server Error " + err.Error())
	}

	return bot, nil
}

// UpdateRobot updates data about a single robot given the Name
func UpdateRobot(req *http.Request) (Robot, error) {
	bot := Robot{}

	oldName := mux.Vars(req)["name"]
	if oldName == "" {
		return bot, errors.New("400. Bad Request")
	}

	err := json.NewDecoder(req.Body).Decode(&bot)
	if err != nil {
		return bot, errors.New("400. Bad request. All fields must be complete")
	}

	// validate form values
	if bot.Name == "" || bot.Model == 0 {
		return bot, errors.New("400. Bad request. All fields must be complete")
	}

	// insert values
	_, err = config.DB.Exec("UPDATE robots SET name = $2, model=$3 WHERE name=$1;", oldName, bot.Name, bot.Model)
	if err != nil {
		return bot, err
	}

	return bot, nil
}

// RemoveRobot removes a single robot given the Name
func RemoveRobot(req *http.Request) error {
	name := mux.Vars(req)["name"]
	if name == "" {
		return errors.New("400. Bad Request")
	}

	_, err := config.DB.Exec("DELETE FROM robots WHERE name=$1;", name)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil

}

// parseRobotRequestValues will look at the form and return an Robot struct
func parseRobotRequestValues(req *http.Request) (Robot, error) {
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

// ParseRobotName returns the name string from the request
func ParseRobotName(req *http.Request) (string, error) {
	name := mux.Vars(req)["name"]
	if name == "" {
		return "", errors.New("400. Bad Request")
	}
	return name, nil
}

// GetRobotByName is an helper function
func GetRobotByName(name string) (Robot, error) {
	bot := Robot{}
	row := config.DB.QueryRow("SELECT * FROM robots WHERE name = $1", name)

	err := row.Scan(&bot.ID, &bot.Name, &bot.Model)
	if err != nil {
		return bot, err
	}

	return bot, nil
}
