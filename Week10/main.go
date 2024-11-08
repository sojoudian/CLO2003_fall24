package main

import (
	"log"
	"net/http"
)
func main(){
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)
	log.Fatal(http.ListenAndServe(":80",nil))
}
