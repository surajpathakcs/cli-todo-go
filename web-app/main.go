package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter , r *http.Request){
	fmt.Fprintln(w,"Welcome to our Go Web Api!")
}

func helloHandler(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	response := map[string] string{"message":"Hello from the Go!"}
	json.NewEncoder(w).Encode(response)
}

func main(){
	http.HandleFunc("/",handleHome)
	http.HandleFunc("/hello",helloHandler)

	http.ListenAndServe(":8000",nil)
}