package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Drone struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	FlightTime       int     `json:"flight_time"`
	SensorSize       string  `json:"sensor_size"`
	WeightMetric     float64 `json:"weight_metric"`
	WeightImperial   float64 `json:"weight_imperial"`
	TopSpeedMetric   float64 `json:"top_speed_metric"`
	TopSpeedImperial float64 `json:"top_speed_imperial"`
	Cost             float64 `json:"cost"`
}

func main() {

	// open the sqlite database
	db, err := sql.Open("sqlite3", "./AeroAnalyzer.db")
	if err != nil {
		panic(err)
	}
	// close it when done
	defer db.Close()

	// query the database
	rows, err := db.Query("SELECT Id,Name,FlightTime,SensorSize,WeightMetric,WeightImperial,TopSpeedMetric,TopSpeedImperial,Cost FROM Drone ORDER BY id")
	if err != nil {
		panic(err)
	}
	// close it when done
	defer rows.Close()

	// our collection of drones
	drones := []Drone{}

	// scanning in our data
	for rows.Next() {
		var drone Drone
		err := rows.Scan(&drone.ID, &drone.Name, &drone.FlightTime, &drone.SensorSize, &drone.WeightMetric, &drone.WeightImperial, &drone.TopSpeedMetric, &drone.TopSpeedImperial, &drone.Cost)
		if err != nil {
			panic(err)
		}
		drones = append(drones, drone)
	}

	// if something goes wonky error out
	if err = rows.Err(); err != nil {
		panic(err)
	}

	// fire up a router
	router := http.NewServeMux()

	if router == nil {
		fmt.Println("Router didn't start")
	}

	// handle an incoming GET request
	router.HandleFunc("/drone", func(w http.ResponseWriter, r *http.Request) {
		// CORS stuff

		// Set the Access-Control-Allow-Origin header to allow requests from any domain
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Set the Access-Control-Allow-Methods header to allow GET requests
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		// Set the Content-Type header to application/json
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(drones)
	})

	// something to put at the prompt to know it's started up
	fmt.Println("Serving")

	// listen and serve!
	serverr := http.ListenAndServe(":8080", router)
	if serverr != nil {
		fmt.Printf("Error Starting Server: %s", serverr)
	}
}
