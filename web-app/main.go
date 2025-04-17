package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter , r *http.Request){
		fmt.Fprintln(w,"Hello There")
	})

	http.ListenAndServe(":8000",nil)
}