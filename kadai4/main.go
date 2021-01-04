package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type Omikuji struct {
	Result string `json:"result"`
}

var omikujiStrings = []string{"大吉", "中吉", "吉", "小吉", "凶"}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var result string
	month := r.FormValue("month")
	day := r.FormValue("day")
	if month == "1" && (day == "1" || day == "2" || day == "3") {
		result = "大吉"
	} else {
		result = omikujiStrings[rand.Intn(len(omikujiStrings))]
	}
	omikuji := Omikuji{result}
	res, err := json.Marshal(omikuji)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func main() {
	rand.Seed(42)
	http.HandleFunc("/", omikujiHandler)
	http.ListenAndServe(":8080", nil)
}
