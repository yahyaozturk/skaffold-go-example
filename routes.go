package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Record is basic defination for API
type Record struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"dateOfBirth"`
}

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Happy Birthday API - Welcome")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "OK"}`))
}

func calculateDuration(date string) int {
	t, _ := time.Parse(layoutISO, date)
	fmt.Println(t) // 1999-12-31 00:00:00 +0000 UTC
	var duration int
	today := time.Now().YearDay()
	birtday := t.YearDay()

	duration = int(time.Now().AddDate(1, 0, 0).Sub(time.Now()).Hours())/24 - (today - birtday)

	if birtday >= today {
		duration = birtday - today
	}

	return duration

}

func saveORupdateRecord(w http.ResponseWriter, r *http.Request) {
	var newRecord Record
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "no valid body data"}`))
		return
	}
	if val, ok := pathParams["name"]; ok {
		err := json.Unmarshal(reqBody, &newRecord)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "invalid datatype for parameter"}`))
			return
		}

		newRecord.Name = val
		err = insertRecord(newRecord)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error inserting data"}`))
			return
		}
		w.WriteHeader(http.StatusNoContent)
		//w.Write(b)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func searchByName(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	if val, ok := pathParams["name"]; ok {
		if val == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "error marshalling data"}`))
			return
		}

		record, err := findRecord(val)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error": "error no record found"}`))
			return
		}
		duration := calculateDuration(record.DateOfBirth)
		log.Println("duration", duration)

		var greetingMessage string
		greetingMessage = "{ \"message\": \"Hello, " + record.Name + "! Your birthday is in " + strconv.Itoa(duration) + " days\" }"
		if duration == 0 {
			greetingMessage = "{ \"message\": \"Hello, " + record.Name + "! Happy birthday\" }"
		}
		log.Println("greetingMessage", greetingMessage)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(greetingMessage))

		return
	}
	w.WriteHeader(http.StatusNotFound)
}
