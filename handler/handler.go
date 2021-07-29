package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Cars struct {
	Id           int    `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Name         string `json:"name"`
	Fuel         string `json:"fuel"`
}

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) Configure() *http.ServeMux {

	userMux := http.NewServeMux()
	userMux.HandleFunc("/docker", getAllCars)
	userMux.HandleFunc("/smth", getSmth)

	return userMux
}

func getAllCars(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "root:epam_pass@tcp(dbase:3306)/carsdb") //127.0.0.1
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM cars")

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer res.Close()
	cars := make([]Cars, 0)
	for res.Next() {

		var car Cars
		err := res.Scan(&car.Id, &car.Manufacturer, &car.Name, &car.Fuel)

		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}
	carsJson, err := json.Marshal(cars)

	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(carsJson))
	return
}

func getSmth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string("this is a smth request"))
	return
}
