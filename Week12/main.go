package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func enableCros(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Method", "POST, GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func saveIPHandler(w http.ResponseWriter, r *http.Request) {
	enableCros(&w)
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed!", http.StatusMethodNotAllowed)
	}
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Error parsing the JSON request!", http.StatusBadRequest)
	}
	currentTime := time.Now().Format("2006-01-02 15:04:05`")
	data["time"] = currentTime

}
func main() {

	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)
	http.HandleFunc("/api/saveIP", saveIPHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
