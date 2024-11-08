package main

import ("net/http")
func main(){
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fs)
	http.ListenAndServe(":80",nil)
}
