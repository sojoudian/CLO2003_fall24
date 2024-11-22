package main

import (
	"log"
	"net/http"
)

func enableCros(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Method", "POST, GET")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func saveIPHandler(w http.ResponseWriter, r *http.Request) {

}
func main() {
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)
	http.HandleFunc("/api/saveIP", saveIPHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
